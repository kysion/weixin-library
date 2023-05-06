package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetUserInfoReq 获取微信会员信息，相当于静默登录
type GetUserInfoReq struct {
	g.Meta `path:"/:appId/userInfo" method:"get" summary:"用户授权" tags:"WeiXin商户服务"`
}

type AppAuthReq struct {
	g.Meta `path:"/:appId/gateway.auth" method:"get" summary:"商户应用授权" tags:"WeiXin商户服务"`
}
