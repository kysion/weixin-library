package gateway

import (
	"github.com/kysion/weixin-library/weixin_service"
)

func init() {
	// 【微信网关】 API服务逻辑
	weixin_service.RegisterGateway(NewGateway())

	// 【第三方平台接口调用凭据】 API服务逻辑
	weixin_service.RegisterTicket(NewTicket())

}
