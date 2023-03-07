package gateway

import "github.com/kysion/weixin-library/weixin_service"

func init() {
	weixin_service.RegisterGateway(NewGateway())
	weixin_service.RegisterTicket(NewTicket())
	weixin_service.RegisterThirdAppConfig(NewThirdAppConfig())
	weixin_service.RegisterMerchantAppConfig(NewMerchantAppConfig())
	weixin_service.RegisterConsumer(NewConsumerConfig())

}
