package weixin_service

import "github.com/kysion/weixin-library/weixin_service"

func init() {

	// 【微信第三方应用SQL表】 服务逻辑
	weixin_service.RegisterThirdAppConfig(NewThirdAppConfig())

	// 【微信商家应用SQL表】 服务逻辑
	weixin_service.RegisterMerchantAppConfig(NewMerchantAppConfig())

	// 【微信消费者SQL表】 服务逻辑
	weixin_service.RegisterConsumer(NewConsumerConfig())

	// 【微信支付商户号SQL表】 服务逻辑
	weixin_service.RegisterPayMerchant(NewPayMerchant())

	// 【微信支付商户号SQL表】 服务逻辑
	weixin_service.RegisterPaySubMerchant(NewPaySubMerchant())

	// 【微信 订阅消息模板SQL表】 服务逻辑
	weixin_service.RegisterSubscribeMessageTemplate(NewSubscribeMessageTemplate())

}
