package merchant

import (
	"context"
	v1 "github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
)

var SubAccount = cSubAccount{}

type cSubAccount struct{}

// GetSubAccountMaxRatio 查询最大分账比例
func (c *cSubAccount) GetSubAccountMaxRatio(ctx context.Context, req *v1.GetSubAccountMaxRatioReq) (*weixin_model.QueryMerchantRatioRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().GetSubAccountMaxRatio(ctx, appId)

	return ret, err
}

// QuerySubAccountOrder 查询分账结果
func (c *cSubAccount) QuerySubAccountOrder(ctx context.Context, req *v1.QuerySubAccountOrderReq) (*weixin_model.OrdersEntityRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().QuerySubAccountOrder(ctx, appId, &req.QueryOrderRequest)

	return (*weixin_model.OrdersEntityRes)(ret), err
}

// UnfreezeOrder 解冻剩余资金API
func (c *cSubAccount) UnfreezeOrder(ctx context.Context, req *v1.UnfreezeOrderReq) (*weixin_model.OrdersEntityRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().UnfreezeOrder(ctx, appId, &req.UnfreezeOrderRequest)

	return (*weixin_model.OrdersEntityRes)(ret), err
}

// SubAccountRequest 请求分账
func (c *cSubAccount) SubAccountRequest(ctx context.Context, req *v1.SubAccountRequestReq) (*weixin_model.OrdersEntityRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().SubAccountRequest(ctx, appId, &req.SubAccountReq)

	return (*weixin_model.OrdersEntityRes)(ret), err
}

// QueryOrderAmount 查询剩余待分金额API
func (c *cSubAccount) QueryOrderAmount(ctx context.Context, req *v1.QueryOrderAmountReq) (*weixin_model.QueryOrderAmountRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().QueryOrderAmount(ctx, appId, &req.QueryOrderAmountRequest)

	return (*weixin_model.QueryOrderAmountRes)(ret), err
}

// AddReceiver 添加分账接收方（相当于绑定分账关系）
func (c *cSubAccount) AddReceiver(ctx context.Context, req *v1.AddReceiverReq) (*weixin_model.AddReceiverRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().AddReceiver(ctx, appId, &req.AddReceiverRequest)

	return (*weixin_model.AddReceiverRes)(ret), err
}

// AddProfitSharingReceivers 添加多个分账关系
func (c *cSubAccount) AddProfitSharingReceivers(ctx context.Context, req *v1.AddProfitSharingReceiversReq) (*weixin_model.AddReceiverRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().AddProfitSharingReceivers(ctx, appId, req.AddReceiverRequestList)

	return (*weixin_model.AddReceiverRes)(ret), err
}

// DeleteReceiver 删除分账接收方（相当于分账关系解绑）
func (c *cSubAccount) DeleteReceiver(ctx context.Context, req *v1.DeleteReceiverReq) (*weixin_model.DeleteReceiverRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := weixin_service.SubAccount().DeleteReceiver(ctx, appId, &req.DeleteReceiverRequest)

	return (*weixin_model.DeleteReceiverRes)(ret), err
}
