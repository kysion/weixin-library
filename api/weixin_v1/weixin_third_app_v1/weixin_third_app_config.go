package weixin_third_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type UpdateStateReq struct {
	g.Meta `path:"/updateState" method:"post" summary:"修改状态" tags:"WeiXin服务商应用"`
	Id     int64 `json:"id" dc:"服务商应用Id"`
	State  int   `json:"state" dc:"状态"`
}

type CreateThirdAppConfigReq struct {
	g.Meta `path:"/createThirdAppConfig" method:"post" summary:"创建服务商应用" tags:"WeiXin服务商应用"`
	weixin_model.WeixinThirdAppConfig
}

type GetThirdAppConfigByIdReq struct {
	g.Meta `path:"/getThirdAppConfigById" method:"post" summary:"根据id获取服务商应用" tags:"WeiXin服务商应用"`
	Id     int64 `json:"id" dc:"服务商应用Id"`
}

type UpdateThirdAppConfigReq struct {
	g.Meta `path:"/updateThirdAppConfig" method:"post" summary:"修改服务商应用基础信息" tags:"WeiXin服务商应用"`
	weixin_model.UpdateThirdAppConfigReq
}

type UpdateThirdAppConfigHttpsReq struct {
	g.Meta `path:"/updateThirdAppConfigHttps" method:"post" summary:"修改Https证书认证" tags:"WeiXin服务商应用"`
	weixin_model.UpdateThirdAppConfigHttpsReq
}

type ThirdAppConfigRes weixin_model.WeixinThirdAppConfig
