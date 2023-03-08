package weixin_third_app

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type AppAuthReq struct {
	// https://weixin.jditco.com/weixin/wx534d1a08aa84c529/gateway.auth
	g.Meta `path:"/:appId/gateway.auth" method:"get" summary:"第三方应用授权" tags:"服务商应用"`
}

// 这个文件属于我们调用

// gateway属于微信平台调用或推送

type UpdateStateReq struct {
	g.Meta `path:"/updateState" method:"post" summary:"修改状态" tags:"服务商应用"`
	Id     int64 `json:"id" dc:"服务商应用Id"`
	State  int   `json:"state" dc:"状态"`
}

type CreateThirdAppConfigReq struct {
	g.Meta `path:"/createThirdAppConfig" method:"post" summary:"创建服务商应用" tags:"服务商应用"`
	weixin_model.WeixinThirdAppConfig
}

type GetThirdAppConfigByIdReq struct {
	g.Meta `path:"/getThirdAppConfigById" method:"post" summary:"根据id获取服务商应用" tags:"服务商应用"`
	Id     int64 `json:"id" dc:"服务商应用Id"`
}

type UpdateThirdAppConfigReq struct {
	g.Meta `path:"/updateThirdAppConfig" method:"post" summary:"修改服务商应用基础信息" tags:"服务商应用"`
	weixin_model.UpdateThirdAppConfigReq
}

type UpdateThirdAppConfigHttpsReq struct {
	g.Meta `path:"/updateThirdAppConfigHttps" method:"post" summary:"修改Https证书认证" tags:"服务商应用"`
	weixin_model.UpdateThirdAppConfigHttpsReq
}

type ThirdAppConfigRes weixin_model.WeixinThirdAppConfig

type WeiXinAuthUserInfoReq struct {
	g.Meta `path:"/:appId/userInfo" method:"get" summary:"获取用户授权" tags:"微信"`
}

type StartPushTicketReq struct {
	g.Meta `path:"/:appId/api_start_push_ticket" method:"get" summary:"让微信重新推送票据" tags:"微信"`
}
