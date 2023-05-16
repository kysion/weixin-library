package merchant

import (
	"context"
	"github.com/kysion/weixin-library/weixin_service"
)

/*
	异步通知地址 (接收支付结果通知接口)
*/

type sMerchantNotify struct {
	//// 异步通知Hook
	//NotifyHook base_hook.BaseHook[hook.NotifyKey, hook.NotifyHookFunc]
	//
	//// 交易Hook
	//TradeHook base_hook.BaseHook[hook.TradeHookKey, hook.TradeHookFunc]
	//
	//// 分账Hook (暂时没用到)
	//SubAccountHook base_hook.BaseHook[hook.SubAccountHookKey, hook.SubAccountHookFunc]
}

func init() {
	weixin_service.RegisterMerchantNotify(NewMerchantNotify())
}

func NewMerchantNotify() *sMerchantNotify {
	return &sMerchantNotify{}
}

// NotifyServices 异步通知地址  用于接收支付宝推送给商户的支付/退款成功的消息。
func (s *sMerchantNotify) NotifyServices(ctx context.Context) (string, error) {

	return "success", nil
}
