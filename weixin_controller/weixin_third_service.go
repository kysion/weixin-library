package weixin_controller

import (
	"context"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_third_app_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
)

var ThirdService = cThirdService{}

type cThirdService struct{}

// GetAuthorizerList 拉取已授权的账号列表
func (c *cThirdService) GetAuthorizerList(ctx context.Context, req *weixin_third_app_v1.GetAuthorizerListReq) (*weixin_model.GetAuthorizerListRes, error) {
	ret, err := weixin_service.ThirdService().GetAuthorizerList(ctx, &req.GetAuthorizerList)

	return ret, err
}

// GetOpenAccount 获取绑定的开放平台账号
func (c *cThirdService) GetOpenAccount(ctx context.Context, req *weixin_third_app_v1.GetOpenAccountReq) (*weixin_model.GetOpenAccountRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)

	ret, err := weixin_service.ThirdService().GetOpenAccount(ctx, appId, merchantApp.AppAuthToken)

	return ret, err
}
