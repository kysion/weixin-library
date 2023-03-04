package gateway

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_service"

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
		if ok {
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
		}

		fmt.Println("验签失败")
		g.RequestFromCtx(ctx).Response.Write("success")
		return
	}

	//
	//// 找出服务商
	//service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, gconv.Int64(data.AppId))
	//// 更新Token

	g.RequestFromCtx(ctx).Response.Write("success")
}

// Callback 接收回调  C端消息
func (s *sGateway) Callback(ctx context.Context, info *weixin_model.AuthorizationCodeRes) {
	// 处理授权
	fmt.Println("callback....")

	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()

	config, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, pathAppId)

	if err == nil && config != nil {
		// 1.验签

		// 2.解密

		return
	}

	// 授权码 过期时间 authorization_code +

	// 获取

	fmt.Println("授权码：\n", info)
}
