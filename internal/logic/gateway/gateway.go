package gateway

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
	"log"
	"sort"
	"strings"

	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_hook"
	"github.com/kysion/weixin-library/weixin_utility"
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
	sys_service.SysLogs().InfoSimple(ctx, nil, "-------------微信的消息通知：serviceInfo....", "WeiXin-CallBack")

	appId := weixin_utility.GetAppIdFormContext(ctx)

	// 第三方代管理
	thirdConfig, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	if err == nil && thirdConfig != nil {
		// 1.验签
		ok := weixin_utility.VerifyByteDanceServer(thirdConfig.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
		if !ok {
			fmt.Println("Services 验签失败")
			g.RequestFromCtx(ctx).Response.Write("success")
			return "验签失败", nil
		}

		// 2.解密
		data := weixin.DecryptEvent(ctx, *eventInfo, *msgInfo)
		fmt.Println("Services 解密后的内容：", data)
		if data != nil && data.AppId != appId { // 说明跨服务商应用操作了
			return "不可跨服务商应用操作了", nil
		}

		/*
			应用授权通知类型InfoType ：
				authorized 授权成功
				updateauthorized 更新授权
				unauthorized 取消授权
		*/
		s.ServiceNotifyTypeHook.Iterator(func(key weixin_enum.ServiceNotifyType, value weixin_hook.ServiceNotifyHookFunc) {
			if data.InfoType == key.Code() {
				g.Try(ctx, func(ctx context.Context) {
					info := g.Map{
						"MsgType": data.InfoType,
						"info":    data,
						"appId":   appId,
					}
					sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin应用通知广播： ------- "+data.InfoType, "sGateway")
					value(ctx, info)
				})
			}
		})

		{
			// 公众号直连
			merchantConfig, _ := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
			if thirdConfig.Id == 0 && merchantConfig != nil {
				// 1.验签
				ok := weixin_utility.VerifyByteDanceServer(merchantConfig.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
				if !ok {
					fmt.Println("Services 验签失败")
					g.RequestFromCtx(ctx).Response.Write("success")
					return "验签失败", nil
				}

				// 2.解密
				//data := weixin.DecryptEvent(ctx, *eventInfo, *msgInfo)
				data := weixin.DecryptMessage(ctx, *eventInfo, *msgInfo)
				fmt.Println("Services 解密后的内容：", data)
				//if data != nil && data.AppId != appId { // 说明跨服务商应用操作了
				//	return "不可跨服务商应用操作了", nil
				//}

				s.ServiceNotifyTypeHook.Iterator(func(key weixin_enum.ServiceNotifyType, value weixin_hook.ServiceNotifyHookFunc) {
					if data.MsgType == key.Code() {
						_ = g.Try(ctx, func(ctx context.Context) {
							info := g.Map{
								"MsgType": data.MsgType,
								"info":    data,
								"appId":   appId,
							}
							_ = sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin应用通知广播： ------- "+data.MsgType, "sGateway")
							value(ctx, info)
						})
					}
				})
			}
		}
	}

	// 找出服务商 Hook
	// 更新Token Hook

	g.RequestFromCtx(ctx).Response.Write("success")

	return "success", nil
}

// Callback 接收回调  C端消息 例如授权通知等。。。 事件回调
func (s *sGateway) Callback(ctx context.Context, info *weixin_model.AuthorizationCodeRes, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) (string, error) {
	// 处理授权
	sys_service.SysLogs().InfoSimple(ctx, nil, "-------------微信的回调消息：callback....", "WeiXin-CallBack")

	fmt.Println("授权码：\n", info)

	appId := weixin_utility.GetAppIdFormContext(ctx)

	thirdConfig, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	if err == nil && thirdConfig != nil {
		// 1.验签
		ok := weixin_utility.VerifyByteDanceServer(thirdConfig.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
		if !ok {
			fmt.Println("Callback 验签失败")
			g.RequestFromCtx(ctx).Response.Write("success")
			err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n验签失败： ", "sGateway")
			return "success", nil
		}

		//// 2.解密
		////eventInfo.AppId = weixin_utility.WeiXinAppIdDecode(eventInfo.AppId)
		//eventInfo.AppId = appId
		//data := weixin.DecryptEvent(ctx, *eventInfo, *msgInfo)
		////data := weixin.DecryptMessage(ctx, *eventInfo, *msgInfo)
		//fmt.Println("解密后的内容：", data)
		//if data == nil {
		//	g.RequestFromCtx(ctx).Response.Write("success")
		//	err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n解密失败： ", "sGateway")
		//	return "success", nil
		//}
		//
		//if data != nil && data.AppId != appId { // 说明跨服务商应用操作了
		//	g.RequestFromCtx(ctx).Response.Write("success")
		//	err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n不可跨服务商操作： "+data.InfoType, "sGateway")
		//
		//	return "success", nil
		//}
		//s.CallbackMsgHook.Iterator(func(key weixin_enum.CallbackMsgType, value weixin_hook.ServiceMsgHookFunc) {
		//	if data.InfoType == key.Code() {
		//		g.Try(ctx, func(ctx context.Context) {
		//			info := g.Map{
		//				"MsgType":   data.InfoType,
		//				"info":      data,
		//				"thirdInfo": thirdConfig,
		//				// "code"
		//				// "openid"
		//			}
		//			sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin回调消息广播： ------- "+data.InfoType, "sGateway")
		//			value(ctx, info)
		//		})
		//	}
		//})

		// 2.解密 --- 自动化测试过程
		// 微信自动化测试账号组
		if eventInfo.AppId == "wx570bc396a51b8ff8" ||
			eventInfo.AppId == "wx9252c5e0bb1836fc" ||
			eventInfo.AppId == "wx8e1097c5bc82cde9" ||
			eventInfo.AppId == "wx14550af28c71a144" ||
			eventInfo.AppId == "wxa35b9c23cfe664eb" ||
			eventInfo.AppId == "wxd101a85aa106f53e" ||
			eventInfo.AppId == "wxc39235c15087f6f3" ||
			eventInfo.AppId == "wx7720d01d4b2a4500" {
			// 需要替换成我们的appId
			eventInfo.AppId = appId
		}

		//data := weixin.DecryptEvent(ctx, *eventInfo, *msgInfo)
		data := weixin.DecryptMessage(ctx, *eventInfo, *msgInfo)
		fmt.Println("Callback 解密后的内容：", data)
		if data != nil && data.MsgType == "text" { // TODO 处理用户消息
			if data.Content == "TESTCOMPONENT_MSG_TYPE_TEXT" {
				newdata := weixin_model.MessageBodyDecrypt{}
				newdata.Content = "TESTCOMPONENT_MSG_TYPE_TEXT_callback"
				newdata.ToUserName = data.FromUserName
				newdata.FromUserName = data.ToUserName
				newdata.CreateTime = gtime.Now().String()
				newdata.MsgType = "text"
				g.RequestFromCtx(ctx).Response.WriteXml(newdata)
				return "successs", nil
			}
		}

		if data == nil {
			g.RequestFromCtx(ctx).Response.Write("success")
			err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n解密失败： ", "sGateway")
			return "success", nil
		}

		if data != nil && data.AppID != "" && data.AppID != appId { // 说明跨服务商应用操作了
			g.RequestFromCtx(ctx).Response.Write("success")
			err = sys_service.SysLogs().ErrorSimple(ctx, nil, "\n不可跨服务商操作： "+data.MsgType, "sGateway")
			return "success", nil
		}

		s.CallbackMsgHook.Iterator(func(key weixin_enum.CallbackMsgType, value weixin_hook.ServiceMsgHookFunc) {
			if data.MsgType == key.Code() {
				g.Try(ctx, func(ctx context.Context) {
					info := g.Map{
						"MsgType":   data.MsgType,
						"info":      data,
						"thirdInfo": thirdConfig,
						// "code"
						// "openid"
					}
					sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin回调消息广播： ------- "+data.MsgType, "sGateway")
					value(ctx, info)
				})
			}
		})

	}

	{

		// 直连
		merchartConfig, _ := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
		if thirdConfig.Id == 0 && merchartConfig != nil {
			// 1.验签
			ok := weixin_utility.VerifyByteDanceServer(merchartConfig.MsgVerfiyToken, msgInfo.TimeStamp, msgInfo.Nonce, msgInfo.Encrypt, msgInfo.MsgSignature)
			if !ok {
				fmt.Println("Callback 验签失败")
				g.RequestFromCtx(ctx).Response.Write("success")
				return "验签失败", nil
			}

			// 2.解密
			data := weixin.DecryptMessage(ctx, *eventInfo, *msgInfo)
			fmt.Println("Callback 解密后的内容：", data)

			s.CallbackMsgHook.Iterator(func(key weixin_enum.CallbackMsgType, value weixin_hook.ServiceMsgHookFunc) {
				if data.MsgType == key.Code() {
					g.Try(ctx, func(ctx context.Context) {
						info := g.Map{
							"MsgType":   data.MsgType,
							"info":      data,
							"thirdInfo": thirdConfig,
							// "code"
							// "openid"
						}
						sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin回调消息广播： ------- "+data.MsgType, "sGateway")
						value(ctx, info)
					})
				}
			})

		}

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
	fmt.Println("gateway.services ----- GET")
	// 与填写的服务器配置中的Token一致
	//const Token = "comjditcokuaimk" // TODO 需要写成动态的
	const Token = "commianlajie" // TODO 需要写成动态的
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

	sign := weixin_utility.Sha1(b.String())

	ok := weixin_utility.CheckSignature(sign, timestamp, nonce, Token)

	if !ok {
		log.Println("微信公众号接入校验失败!")
		return ""
	}

	log.Println("微信公众号接入校验成功!")

	g.RequestFromCtx(ctx).Response.Write(echostr)
	return echostr
}
