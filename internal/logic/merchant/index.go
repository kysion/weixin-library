package merchant

import (
	"github.com/kysion/weixin-library/internal/logic/gateway"
	"github.com/kysion/weixin-library/weixin_service"
)

/*
	微信商家应用服务
*/

func init() {
	// 【微信应用授权】 API服务逻辑
	weixin_service.RegisterAppAuth(NewAppAuth())

	// 【微信异步通知】 API服务逻辑
	weixin_service.RegisterMerchantNotify(gateway.NewMerchantNotify())

	// 【微信支付】 API服务逻辑
	weixin_service.RegisterWeiXinPay(NewWeiXinPay())

	// 【微信小程序开发管理】 API服务逻辑
	weixin_service.RegisterAppVersion(NewAppVersion())

	// 【微信分账】 API服务逻辑
	weixin_service.RegisterSubAccount(NewSubAccount())

	// 【微信支付特约商户】 API服务逻辑
	weixin_service.RegisterSubMerchant(NewSubMerchant())

	// 【微信基础消息能力/用户事件userEvent】 API服务逻辑
	weixin_service.RegisterUserEvent(NewUserEvent())

	// 【微信用户授权】 API服务逻辑
	weixin_service.RegisterUserAuth(NewUserAuth())

	// 【微信小程序消息/订阅消息】 API服务逻辑
	weixin_service.RegisterSubscribeMessage(NewSubscribeMessage())

	// 【微信小程序码&小程序链接】 API服务逻辑
	weixin_service.RegisterTinyAppUrl(NewTinyAppUrl())

}
