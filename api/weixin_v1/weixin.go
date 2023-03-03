package weixin_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type WeiXinServicesReq struct { // 第三方应用的相关消息
	g.Meta `path:"/gateway.services" method:"post" summary:"消息接收" tags:"微信"`
}

type WeiXinCallbackReq struct { // 商家相关消息
	g.Meta `path:"/gateway.callback" method:"get"  summary:"网关回调" tags:"微信"`
}

type WeiXinAuthUserInfoReq struct {
	g.Meta `path:"/gateway.auth" method:"get" summary:"获取用户授权" tags:"微信"`
}

type CheckSignatureReq struct {
	g.Meta    `path:"/gateway.services" method:"get" summary:"校验" tags:"微信"`
	Signature string     `json:"signature" dc:"微信加密签名"`
	Timestamp gtime.Time `json:"timestamp" dc:"时间戳"`
	Nonce     string     `json:"nonce" dc:"随机数"`
	Echostr   string     `json:"echostr" dc:"随机字符串"`
}

type StringRes string
