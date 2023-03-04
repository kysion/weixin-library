package weixin_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type ServicesReq struct { // 第三方应用的相关消息
	g.Meta `path:"/:appId/gateway.services" method:"post" summary:"消息接收" tags:"微信"`

	// 推送的Ticket等加密数据
	weixin_model.EventEncryptMsgReq
	weixin_model.MessageEncryptReq
}

type CallbackReq struct { // 回调消息
	g.Meta `path:"/:appId/gateway.callback" method:"get"  summary:"网关回调" tags:"微信"`
	weixin_model.AuthorizationCodeRes
}

type StringRes string

// appId 是用来区分不同的服务商   同一套项目服务不同的服务商
