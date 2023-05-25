package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

var SubMerchant = cSubMerchant{}

type cSubMerchant struct{}

// GetAuditStateByBusinessCode 根据业务申请编号查询申请状态
func (c *cSubMerchant) GetAuditStateByBusinessCode(ctx context.Context, req *v1.GetAuditStateByBusinessCodeReq) (*weixin_model.SubMerchantAuditStateRes, error) {
	spMchId := g.RequestFromCtx(ctx).Get("spMchId").String()

	ret, err := weixin_service.SubMerchant().GetAuditStateByBusinessCode(ctx, spMchId, req.BusinessCode)

	return ret, err
}

// GetAuditStateByApplymentId 根据申请单号查询申请状态
func (c *cSubMerchant) GetAuditStateByApplymentId(ctx context.Context, req *v1.GetAuditStateByApplymentIdReq) (*weixin_model.SubMerchantAuditStateRes, error) {
	spMchId := g.RequestFromCtx(ctx).Get("spMchId").String()

	ret, err := weixin_service.SubMerchant().GetAuditStateByApplymentId(ctx, spMchId, req.ApplymentId)

	return ret, err
}

// GetSettlement 查询结算账号
func (c *cSubMerchant) GetSettlement(ctx context.Context, req *v1.GetSettlementReq) (*weixin_model.SettlementRes, error) {
	subMchId := g.RequestFromCtx(ctx).Get("subMchId").String()

	ret, err := weixin_service.SubMerchant().GetSettlement(ctx, subMchId)

	return ret, err
}

// UpdateSettlement 修改结算账号,成功会返回application_no，作为查询申请状态的唯一标识
func (c *cSubMerchant) UpdateSettlement(ctx context.Context, req *v1.UpdateSettlementReq) (api_v1.StringRes, error) {
	subMchId := g.RequestFromCtx(ctx).Get("subMchId").String()

	ret, err := weixin_service.SubMerchant().UpdateSettlement(ctx, subMchId, &req.UpdateSettlementReq)

	return (api_v1.StringRes)(ret), err
}

// GetSettlementAuditState 查询结算账户修改审核状态
func (c *cSubMerchant) GetSettlementAuditState(ctx context.Context, req *v1.GetSettlementAuditStateReq) (*weixin_model.SettlementRes, error) {
	subMchId := g.RequestFromCtx(ctx).Get("subMchId").String()

	ret, err := weixin_service.SubMerchant().GetSettlementAuditState(ctx, subMchId, req.ApplicationNo)

	return ret, err
}
