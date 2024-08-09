package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type UpdateStateReq struct {
	g.Meta `path:"/updateState" method:"post" summary:"修改状态" tags:"WeiXin商家应用"`
	Id     int64 `json:"id" dc:"商家应用Id"`
	State  int   `json:"state" dc:"状态"`
}

type CreateMerchantAppConfigReq struct {
	g.Meta `path:"/createMerchantAppConfig" method:"post" summary:"创建商家应用" tags:"WeiXin商家应用"`
	weixin_model.WeixinMerchantAppConfig
}

type GetMerchantAppConfigByIdReq struct {
	g.Meta `path:"/getMerchantAppConfigById" method:"post" summary:"根据id获取商家应用" tags:"WeiXin商家应用"`
	Id     int64 `json:"id" dc:"商家应用Id"`
}

type UpdateMerchantAppConfigReq struct {
	g.Meta `path:"/updateMerchantAppConfig" method:"post" summary:"修改商家应用基础信息" tags:"WeiXin商家应用"`
	weixin_model.UpdateMerchantAppConfigReq
}

type UpdateMerchantAppConfigHttpsReq struct {
	g.Meta `path:"/updateMerchantAppConfigHttps" method:"post" summary:"修改Https证书认证" tags:"WeiXin商家应用"`
	weixin_model.UpdateMerchantAppConfigHttpsReq
}

type GetPolicyReq struct {
	g.Meta `path:"/:appId/getPolicy" method:"get" summary:"获取协议" tags:"WeiXin商家应用"`
}

type MerchantAppConfigRes weixin_model.WeixinMerchantAppConfig
