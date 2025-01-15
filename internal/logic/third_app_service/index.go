package third_app_service

import "github.com/kysion/weixin-library/weixin_service"

/*
	微信第三方平台应用服务 （微信服务商）
*/

func init() {
	// 【微信第三方平台应用服务】 API服务逻辑
	weixin_service.RegisterThirdService(NewThirdService())

	// 。。。。

}
