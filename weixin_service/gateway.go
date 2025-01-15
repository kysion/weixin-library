// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package weixin_service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_model/weixin_hook"
)

type (
	ITicket interface {
		// Ticket 票据具体服务
		Ticket(ctx context.Context, info g.Map) bool
		// GetTicket 获取票据
		GetTicket(ctx context.Context, appId string) (string, error)
	}
	IGateway interface {
		GetCallbackMsgHook() *base_hook.BaseHook[weixin_enum.CallbackMsgType, weixin_hook.ServiceMsgHookFunc]
		GetServiceNotifyTypeHook() *base_hook.BaseHook[weixin_enum.ServiceNotifyType, weixin_hook.ServiceNotifyHookFunc]
		// Services 接收消息通知
		Services(ctx context.Context, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) (string, error)
		// Callback 接收回调  C端消息 例如授权通知等。。。 事件回调
		Callback(ctx context.Context, info *weixin_model.AuthorizationCodeRes, eventInfo *weixin_model.EventEncryptMsgReq, msgInfo *weixin_model.MessageEncryptReq) (string, error)
		// WXCheckSignature 微信接入校验 设置Token需要验证
		WXCheckSignature(ctx context.Context, signature, timestamp, nonce, echostr string) string
	}
	IMerchantNotify interface {
		InstallNotifyHook(hookKey weixin_hook.NotifyKey, hookFunc weixin_hook.NotifyHookFunc)
		InstallTradeHook(hookKey weixin_hook.TradeHookKey, hookFunc weixin_hook.TradeHookFunc)
		// NotifyServices 异步通知地址  用于接收支付宝推送给商户的支付/退款成功的消息。
		NotifyServices(ctx context.Context) (string, error)
	}
)

var (
	localGateway        IGateway
	localMerchantNotify IMerchantNotify
	localTicket         ITicket
)

func Gateway() IGateway {
	if localGateway == nil {
		panic("implement not found for interface IGateway, forgot register?")
	}
	return localGateway
}

func RegisterGateway(i IGateway) {
	localGateway = i
}

func MerchantNotify() IMerchantNotify {
	if localMerchantNotify == nil {
		panic("implement not found for interface IMerchantNotify, forgot register?")
	}
	return localMerchantNotify
}

func RegisterMerchantNotify(i IMerchantNotify) {
	localMerchantNotify = i
}

func Ticket() ITicket {
	if localTicket == nil {
		panic("implement not found for interface ITicket, forgot register?")
	}
	return localTicket
}

func RegisterTicket(i ITicket) {
	localTicket = i
}
