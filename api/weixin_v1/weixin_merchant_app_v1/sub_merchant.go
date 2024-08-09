package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

type GetAuditStateByBusinessCodeReq struct {
	g.Meta `path:"/:spMchId/getAuditStateByBusinessCode" method:"post" summary:"根据业务申请编号查询申请状态" tags:"WeiXin商户进件"`

	BusinessCode string `json:"business_code" dc:"商户进件申请业务编号"`
}

type GetAuditStateByApplymentIdReq struct {
	g.Meta      `path:"/:spMchId/getAuditStateByApplymentId" method:"post" summary:"根据申请单号查询申请状态" tags:"WeiXin商户进件"`
	ApplymentId string `json:"applymentId" dc:"商户进件申请单号"`
}

type GetSettlementReq struct {
	g.Meta `path:"/:subMchId/getSettlement" method:"post" summary:"查询结算账号" tags:"WeiXin商户进件"`
}

type UpdateSettlementReq struct {
	g.Meta `path:"/:subMchId/updateSettlement" method:"post" summary:"修改结算账号" tags:"WeiXin商户进件"`

	weixin_model.UpdateSettlementReq
}

type GetSettlementAuditStateReq struct {
	g.Meta        `path:"/:subMchId/getSettlementAuditState" method:"post" summary:"查询修改结算账户审核状态" tags:"WeiXin商户进件"`
	ApplicationNo string `json:"application_no" dc:""`
}
