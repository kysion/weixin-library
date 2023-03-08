package gateway

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_service"
	"log"
	"sort"
	"strings"

	"github.com/kysion/weixin-library/utility"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_model/weixin_hook"
)

type sGateway struct {
	base_hook.BaseHook[weixin_enum.InfoType, weixin_hook.ServiceMsgHookFunc]
}

func NewGateway() *sGateway {
	// 初始化文件内容
	return &sGateway{}
}

func (s *sGateway) InstallHook(infoType weixin_enum.InfoType, hookFunc weixin_hook.ServiceMsgHookFunc) {
	s.BaseHook.InstallHook(infoType, hookFunc)
}

// Services 接收消息通知
func (s *sGateway) Services(ctx context.Context, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()

	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, pathAppId)

	if err == nil && config != nil {
		// 1.验签
		ok := utility.VerifyByteDanceServer(config.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
		if !ok {
			fmt.Println("验签失败")
			g.RequestFromCtx(ctx).Response.Write("success")
			return
		}

		// 2.解密
		data := weixin.Decrypt(ctx, *eventInfo, *msgInfo)
		fmt.Println("解密后的内容：", data)

		if data.AppId != pathAppId { // 说明跨服务商应用操作了
			return
		}

		s.Iterator(func(key weixin_enum.InfoType, value weixin_hook.ServiceMsgHookFunc) {
			if data.InfoType == key.Code() {
				g.Try(ctx, func(ctx context.Context) {
					info := g.Map{
						"MsgType": data.InfoType,
						"info":    data,
					}
					value(ctx, info)
				})
			}
		})
		g.RequestFromCtx(ctx).Response.Write("success")
		return
	}

	//
	//// 找出服务商
	//service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, gconv.Int64(data.AppId))
	//// 更新Token

	g.RequestFromCtx(ctx).Response.Write("success")
}

// Callback 接收回调  C端消息 例如授权通知等。。。
func (s *sGateway) Callback(ctx context.Context, info *weixin_model.AuthorizationCodeRes, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) {
	// 处理授权
	fmt.Println("callback....")
	fmt.Println("授权码：\n", info)

	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()

	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, pathAppId)

	if err == nil && config != nil {
		// 1.验签
		ok := utility.VerifyByteDanceServer(config.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
		if !ok {
			fmt.Println("验签失败")
			g.RequestFromCtx(ctx).Response.Write("success")
			return
		}

		// 2.解密
		data := weixin.Decrypt(ctx, *eventInfo, *msgInfo)
		fmt.Println("解密后的内容：", data)

		if data.AppId != pathAppId { // 说明跨服务商应用操作了
			return
		}

		/*
			授权通知类型InfoType ：
				authorized 授权成功
				updateauthorized 更新授权
				unauthorized 取消授权
		*/
		s.Iterator(func(key weixin_enum.InfoType, value weixin_hook.ServiceMsgHookFunc) {
			if data.InfoType == key.Code() {
				g.Try(ctx, func(ctx context.Context) {
					info := g.Map{
						"MsgType":   data.InfoType,
						"info":      data,
						"thirdInfo": config,
					}
					value(ctx, info)
				})
			}
		})

		g.RequestFromCtx(ctx).Response.Write("success")

		return
	}

	// 授权码 过期时间 authorization_code + 时间

	// 获取商家接口调用凭据 authorizer_access_token  + authorizer_refresh_token + 时间

	// authorizer_access_token就又能调用各种接口了

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
