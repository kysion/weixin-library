package weixin_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_service"
)

var WeiXinMerchantAppConfig = cWeiXinMerchantAppConfig{}

type cWeiXinMerchantAppConfig struct{}

// UpdateState 修改状态
func (s *cWeiXinMerchantAppConfig) UpdateState(ctx context.Context, req *weixin_merchant_app_v1.UpdateStateReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := weixin_service.MerchantAppConfig().UpdateState(ctx, req.Id, req.State)
			return ret == true, err
		},
		// 记得添加权限
		// weixin_permission.MerchantAppConfig.PermissionType.Update,
	)
}

// CreateMerchantAppConfig  创建第三方应用配置信息
func (s *cWeiXinMerchantAppConfig) CreateMerchantAppConfig(ctx context.Context, req *weixin_merchant_app_v1.CreateMerchantAppConfigReq) (*weixin_merchant_app_v1.MerchantAppConfigRes, error) {
	//return funs.CheckPermission(ctx,
	//	func() (*v1.MerchantAppConfigRes, error) {
	//		ret, err := weixin_service.MerchantAppConfig().CreateMerchantAppConfig(ctx, &req.WeixinMerchantAppConfig)
	//		return (*v1.MerchantAppConfigRes)(ret), err
	//	},
	//	// 记得添加权限
	//	// weixin_permission.MerchantAppConfig.PermissionType.Update,
	//)

	ret, err := weixin_service.MerchantAppConfig().CreateMerchantAppConfig(ctx, &req.WeixinMerchantAppConfig)
	return (*weixin_merchant_app_v1.MerchantAppConfigRes)(ret), err
}

// GetMerchantAppConfigByAppId 根据AppId查找第三方应用配置信息
func (s *cWeiXinMerchantAppConfig) GetMerchantAppConfigByAppId(ctx context.Context, req *weixin_merchant_app_v1.GetMerchantAppConfigByIdReq) (*weixin_merchant_app_v1.MerchantAppConfigRes, error) {
	return funs.CheckPermission(ctx,
		func() (*weixin_merchant_app_v1.MerchantAppConfigRes, error) {
			ret, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigById(ctx, req.Id)
			return (*weixin_merchant_app_v1.MerchantAppConfigRes)(ret), err
		},
		// 记得添加权限
		// weixin_permission.MerchantAppConfig.PermissionType.Update,
	)

}

// UpdateAppConfig 修改服务商基础信息
func (s *cWeiXinMerchantAppConfig) UpdateAppConfig(ctx context.Context, req *weixin_merchant_app_v1.UpdateMerchantAppConfigReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := weixin_service.MerchantAppConfig().UpdateAppConfig(ctx, &req.UpdateMerchantAppConfigReq)
			return ret == true, err
		},
		// 记得添加权限
		// weixin_permission.MerchantAppConfig.PermissionType.Update,
	)
	//
	//ret, err := weixin_service.MerchantAppConfig().UpdateAppConfig(ctx, &req.UpdateMerchantAppConfigReq)
	//
	//return ret == true, err
}

// UpdateMerchantAppConfigHttps 修改服务商应用Https配置
func (s *cWeiXinMerchantAppConfig) UpdateMerchantAppConfigHttps(ctx context.Context, req *weixin_merchant_app_v1.UpdateMerchantAppConfigHttpsReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.MerchantAppConfig().UpdateAppConfigHttps(ctx, &req.UpdateMerchantAppConfigHttpsReq)
	return ret == true, err
}
