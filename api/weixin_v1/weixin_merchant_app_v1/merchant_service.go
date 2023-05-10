package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AppAuthReq struct {
	g.Meta `path:"/:appId/gateway.auth" method:"get" summary:"商户应用授权" tags:"WeiXin商户服务"`
}

type AuthResReq struct { // AppAuth 应用授权响应接收地址 （后续会是前端地址）
	g.Meta `path:"/:appId/gateway.authRes" method:"get" summary:"授权响应" tags:"WeiXin"`

	AuthCode  string `json:"auth_code"`
	ExpiresIn string `json:"expires_in"`
}

type UserAuthReq struct {
	g.Meta `path:"/:appId/userAuth" method:"get" summary:"用户授权" tags:"WeiXin商户服务"`
}

// GetUserInfoReq 获取微信会员信息，相当于静默登录
type GetUserInfoReq struct {
	g.Meta `path:"/:appId/userInfo" method:"get" summary:"获取用户信息" tags:"WeiXin商户服务"`
}
