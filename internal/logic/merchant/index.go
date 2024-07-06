package merchant

import "github.com/kysion/weixin-library/weixin_service"

func init() {
	weixin_service.RegisterAppAuth(NewAppAuth())

	weixin_service.RegisterMerchantNotify(NewMerchantNotify())

	weixin_service.RegisterWeiXinPay(NewWeiXinPay())

	weixin_service.RegisterAppVersion(NewAppVersion())

	weixin_service.RegisterSubAccount(NewSubAccount())

	weixin_service.RegisterSubMerchant(NewSubMerchant())

	weixin_service.RegisterUserEvent(NewUserEvent())

	weixin_service.RegisterUserAuth(NewUserAuth())

	weixin_service.RegisterSubscribeMessage(NewSubscribeMessage())

}
