package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AppAuthReq struct {
	g.Meta `path:"/:appId/gateway.auth" method:"get" summary:"商户应用授权" tags:"WeiXin商户服务"`
}

type AuthResReq struct { // AppAuth 应用授权响应接收地址 （后续会是前端地址）
	g.Meta `path:"/:appId/gateway.authRes" method:"get" summary:"应用授权响应" tags:"WeiXin商户服务"`

	AuthCode  string `json:"auth_code"`
	ExpiresIn string `json:"expires_in"`
}

type UserAuthReq struct {
	g.Meta `path:"/:appId/userAuth" method:"get" summary:"用户授权" tags:"WeiXin商户服务"`
}

type UserAuthResReq struct { // UserAuth 用户授权响应接收地址 （后续会是前端地址）
	g.Meta `path:"/:appId/gateway.userAuthRes" method:"get" summary:"用户授权响应" tags:"WeiXin商户服务"`

	Code      string `json:"code"`
	ExpiresIn string `json:"expires_in"`
}

// GetUserInfoReq 获取微信会员信息，相当于静默登录
type GetUserInfoReq struct {
	g.Meta `path:"/:appId/userInfo" method:"get" summary:"获取用户信息" tags:"WeiXin商户服务"`
}

type UserLoginReq struct {
	g.Meta `path:"/:appId/userLogin" method:"get" summary:"用户登录" tags:"WeiXin商户服务"`

	Code      string `json:"code"`
	ExpiresIn string `json:"expires_in"`
}
