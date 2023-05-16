package weixin_pay_merchant_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type GetPayMerchantByIdReq struct {
	g.Meta `path:"/getPayMerchantById" method:"post" summary:"根据id查询|信息" tags:"WeiXin支付商户号"`
	Id     int64 `json:"id" dc:"商户号记录Id"`
}

type GetPayMerchantByMchidReq struct {
	g.Meta `path:"/getPayMerchantByMchid" method:"post" summary:"根据Mchid查询|信息" tags:"WeiXin支付商户号"`
	MchId  int `json:"mchId" dc:"特约商户商户号"`
}

type GetPayMerchantBySysUserIdReq struct {
	g.Meta    `path:"/getPayMerchantBySysUserId" method:"post" summary:"根据用户id查询|信息" tags:"WeiXin支付商户号"`
	SysUserId int64 `json:"id" dc:"用户Id"`
}

type CreatePayMerchantReq struct {
	g.Meta `path:"/createPayMerchant" method:"post" summary:"创建商户号配置信息" tags:"WeiXin支付商户号"`
	weixin_model.PayMerchant
}

type UpdatePayMerchantReq struct {
	g.Meta `path:"/updatePayMerchant" method:"post" summary:"更新商户号配置｜信息" tags:"WeiXin支付商户号"`
	Id     int64 `json:"id" dc:"商户号记录Id"`
	weixin_model.UpdatePayMerchant
}

type SetCertAndKeyReq struct {
	g.Meta `path:"/setCertAndKey" method:"post" summary:"设置商户号证书及密钥文件" tags:"WeiXin支付商户号"`
	Id     int64 `json:"id" dc:"商户号记录Id"`
	weixin_model.SetCertAndKey
}

type SetAuthPathReq struct {
	g.Meta `path:"/setAuthPath" method:"post" summary:"设置商户号授权目录" tags:"WeiXin支付商户号"`
	weixin_model.SetAuthPath
}

type SetPayMerchantUnionIdReq struct {
	g.Meta `path:"/setPayMerchantUnionId" method:"post" summary:"设置商户号关联的AppId" tags:"WeiXin支付商户号"`
	weixin_model.SetPayMerchantUnionId
}

type SetBankcardAccountReq struct {
	g.Meta `path:"/setBankcardAccount" method:"post" summary:"设置商户号银行卡号" tags:"WeiXin支付商户号"`
	weixin_model.SetBankcardAccount
}
