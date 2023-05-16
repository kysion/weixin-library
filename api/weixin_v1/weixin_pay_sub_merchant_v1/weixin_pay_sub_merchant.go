package weixin_pay_sub_merchant_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type GetPaySubMerchantByIdReq struct {
	g.Meta `path:"/getPaySubMerchantById" method:"post" summary:"根据id查询|信息" tags:"WeiXin支付特约商户"`
	Id     int64 `json:"id" dc:"特约商户Id"`
}

type GetPaySubMerchantByMchidReq struct {
	g.Meta `path:"/getPaySubMerchantByMchid" method:"post" summary:"根据Mchid查询|信息" tags:"WeiXin支付特约商户"`
	MchId  int `json:"mchId" dc:"商户号"`
}

type GetPaySubMerchantBySysUserIdReq struct {
	g.Meta    `path:"/getPaySubMerchantBySysUserId" method:"post" summary:"根据用户id查询|信息" tags:"WeiXin支付特约商户"`
	SysUserId int64 `json:"id" dc:"用户Id"`
}

type CreatePaySubMerchantReq struct {
	g.Meta `path:"/createPaySubMerchant" method:"post" summary:"创建特约商户" tags:"WeiXin支付特约商户"`
	weixin_model.WeixinPaySubMerchant
}

type UpdatePaySubMerchantReq struct {
	g.Meta `path:"/updatePaySubMerchant" method:"post" summary:"更新特约商户|信息" tags:"WeiXin支付特约商户"`
	Id     int64 `json:"id" dc:"特约商户Id"`
	weixin_model.UpdatePaySubMerchant
}

type SetSubMerchantAuthPathReq struct {
	g.Meta `path:"/setSubMerchantAuthPath" method:"post" summary:"设置特约商户授权目录" tags:"WeiXin支付特约商户"`
	weixin_model.SetSubMerchantAuthPath
}
