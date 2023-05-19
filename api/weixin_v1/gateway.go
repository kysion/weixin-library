package weixin_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/weixin-library/weixin_model"
)

// appId 是用来区分不同的服务商   同一套项目服务不同的服务商

type ServicesReq struct { // 第三方应用的相关消息
	g.Meta `path:"/:appId/gateway.services" method:"post" summary:"消息接收" tags:"WeiXin"`

	// 推送的Ticket等加密数据
	weixin_model.EventEncryptMsgReq
	weixin_model.MessageEncryptReq
}

// https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

type CallbackReq struct { // 回调消息
	g.Meta `path:"/:APPID/:appId/gateway.callback" method:"get"  summary:"网关回调Get" tags:"WeiXin"`
	weixin_model.AuthorizationCodeRes
	// 推送的Ticket等加密数据
	weixin_model.EventEncryptMsgReq
	weixin_model.MessageEncryptReq
}

type CallbackPostReq struct { // 回调消息
	g.Meta `path:"/:APPID/:appId/gateway.callback" method:"post"  summary:"网关回调Post" tags:"WeiXin"`
	weixin_model.AuthorizationCodeRes
	// 推送的Ticket等加密数据
	weixin_model.EventEncryptMsgReq
	weixin_model.MessageEncryptReq
}

type StringRes string

// CheckSignatureReq 这一个在设置服务器配置设置Token需要验证
type CheckSignatureReq struct {
	g.Meta    `path:"/:appId/gateway.services" method:"get" summary:"校验" tags:"WeiXin"`
	Signature string     `json:"signature" dc:"微信加密签名"`
	Timestamp gtime.Time `json:"timestamp" dc:"时间戳"`
	Nonce     string     `json:"nonce" dc:"随机数"`
	Echostr   string     `json:"echostr" dc:"随机字符串"`
}

type StartPushTicketReq struct {
	g.Meta `path:"/:appId/api_start_push_ticket" method:"get" summary:"让微信重新推送票据" tags:"WeiXin"`
}

// https://www.kuaimk.com/weixin/wxfuvfrnkh1lkm7/gateway.notify

type NotifyServicesReq struct {
	g.Meta `path:"/:appId/gateway.notify" method:"post" summary:"支付异步通知" tags:"WeiXin"`
}
