package weixin_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/weixin-library/weixin_model"
)

// appId 是用来区分不同的服务商   同一套项目服务不同的服务商

type ServicesReq struct { // 第三方应用的相关消息
	g.Meta `path:"/:appId/gateway.services" method:"post" summary:"消息接收" tags:"微信"`

	// 推送的Ticket等加密数据
	weixin_model.EventEncryptMsgReq
	weixin_model.MessageEncryptReq
}

type CallbackReq struct { // 回调消息
	g.Meta `path:"/:appId/gateway.callback" method:"get"  summary:"网关回调" tags:"微信"`
	weixin_model.AuthorizationCodeRes
	// 推送的Ticket等加密数据
	weixin_model.EventEncryptMsgReq
	weixin_model.MessageEncryptReq
}

type StringRes string

// CheckSignatureReq 这一个在设置服务器配置设置Token需要验证
type CheckSignatureReq struct {
	g.Meta    `path:"/:appId/gateway.services" method:"get" summary:"校验" tags:"微信"`
	Signature string     `json:"signature" dc:"微信加密签名"`
	Timestamp gtime.Time `json:"timestamp" dc:"时间戳"`
	Nonce     string     `json:"nonce" dc:"随机数"`
	Echostr   string     `json:"echostr" dc:"随机字符串"`
}
