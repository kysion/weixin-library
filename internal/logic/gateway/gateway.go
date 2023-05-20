package gateway

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
	"log"
	"sort"
	"strings"

	"github.com/kysion/weixin-library/utility"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_hook"
)

type sGateway struct {
	// 回调消息Hook
	CallbackMsgHook base_hook.BaseHook[weixin_enum.CallbackMsgType, weixin_hook.ServiceMsgHookFunc]

	// 应用通知Hook
	ServiceNotifyTypeHook base_hook.BaseHook[weixin_enum.ServiceNotifyType, weixin_hook.ServiceNotifyHookFunc]
}

func NewGateway() *sGateway {
	// 初始化文件内容
	return &sGateway{}
}

func (s *sGateway) GetCallbackMsgHook() *base_hook.BaseHook[weixin_enum.CallbackMsgType, weixin_hook.ServiceMsgHookFunc] {
	return &s.CallbackMsgHook
}

func (s *sGateway) GetServiceNotifyTypeHook() *base_hook.BaseHook[weixin_enum.ServiceNotifyType, weixin_hook.ServiceNotifyHookFunc] {
	return &s.ServiceNotifyTypeHook
}

// Services 接收消息通知
func (s *sGateway) Services(ctx context.Context, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) (string, error) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	if err == nil && config != nil {
		// 1.验签
		ok := utility.VerifyByteDanceServer(config.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
		if !ok {
			fmt.Println("验签失败")
			g.RequestFromCtx(ctx).Response.Write("success")
			return "验签失败", nil
		}

		// 2.解密
		data := weixin.Decrypt(ctx, *eventInfo, *msgInfo)
		fmt.Println("解密后的内容：", data)
		if data != nil && data.AppId != appId { // 说明跨服务商应用操作了
			return "不可跨服务商应用操作了", nil
		}
		s.ServiceNotifyTypeHook.Iterator(func(key weixin_enum.ServiceNotifyType, value weixin_hook.ServiceNotifyHookFunc) {
			if data.InfoType == key.Code() {
				g.Try(ctx, func(ctx context.Context) {
					info := g.Map{
						"MsgType": data.InfoType,
						"info":    data,
					}
					sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin应用通知广播： ------- "+data.InfoType, "sGateway")
					value(ctx, info)
				})
			}
		})

	}

	// 找出服务商 Hook
	// 更新Token Hook

	g.RequestFromCtx(ctx).Response.Write("success")

	return "success", nil
}

// Callback 接收回调  C端消息 例如授权通知等。。。
func (s *sGateway) Callback(ctx context.Context, info *weixin_model.AuthorizationCodeRes, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) (string, error) {
	// 处理授权
	fmt.Println("callback....")
	fmt.Println("授权码：\n", info)

	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	if err == nil && config != nil {
		// 1.验签
		ok := utility.VerifyByteDanceServer(config.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
		if !ok {
			fmt.Println("验签失败")
			g.RequestFromCtx(ctx).Response.Write("success")
			err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n验签失败： ", "sGateway")
			return "success", nil
		}

		// 2.解密
		data := weixin.Decrypt(ctx, *eventInfo, *msgInfo)
		fmt.Println("解密后的内容：", data)
		if data == nil {
			g.RequestFromCtx(ctx).Response.Write("success")
			err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n解密失败： ", "sGateway")
			return "success", nil
		}

		if data != nil && data.AppId != appId { // 说明跨服务商应用操作了
			g.RequestFromCtx(ctx).Response.Write("success")
			err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n不可跨服务商操作： "+data.InfoType, "sGateway")

			return "success", nil
		}

		/*
			应用授权通知类型InfoType ：
				authorized 授权成功
				updateauthorized 更新授权
				unauthorized 取消授权
		*/
		s.CallbackMsgHook.Iterator(func(key weixin_enum.CallbackMsgType, value weixin_hook.ServiceMsgHookFunc) {
			if data.InfoType == key.Code() {
				g.Try(ctx, func(ctx context.Context) {
					info := g.Map{
						"MsgType":   data.InfoType,
						"info":      data,
						"thirdInfo": config,
						// "code"
						// "openid"
					}
					sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin回调消息广播： ------- "+data.InfoType, "sGateway")
					value(ctx, info)
				})
			}
		})

	}

	// 授权码 过期时间 authorization_code + 时间

	// 获取商家接口调用凭据 authorizer_access_token  + authorizer_refresh_token + 时间

	// authorizer_access_token就又能调用各种接口了

	// 存储authorizer_access_token至数据库
	g.RequestFromCtx(ctx).Response.Write("success")

	return "success", nil
}

// WXCheckSignature 微信接入校验 设置Token需要验证
func (s *sGateway) WXCheckSignature(ctx context.Context, signature, timestamp, nonce, echostr string) string {
	// 与填写的服务器配置中的Token一致
	const Token = "comjditcokuaimk"
	fmt.Println(signature + "、" + timestamp + "、" + nonce + "、" + echostr)
	arr := []string{timestamp, nonce, Token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(Token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	sign := utility.Sha1(b.String())

	ok := utility.CheckSignature(sign, timestamp, nonce, Token)

	if !ok {
		log.Println("微信公众号接入校验失败!")
		return ""
	}

	log.Println("微信公众号接入校验成功!")

	g.RequestFromCtx(ctx).Response.Write(echostr)
	return echostr
}
