package weixin_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_third_app_v1"
	"github.com/kysion/weixin-library/weixin_service"
)

var WeiXinThirdAppConfig = cWeiXinThirdAppConfig{}

type cWeiXinThirdAppConfig struct{}

// UpdateState 修改状态
func (s *cWeiXinThirdAppConfig) UpdateState(ctx context.Context, req *weixin_third_app_v1.UpdateStateReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := weixin_service.ThirdAppConfig().UpdateState(ctx, req.Id, req.State)
			return ret == true, err
		},
		// 记得添加权限
		// weixin_permission.ThirdAppConfig.PermissionType.Update,
	)
}

// CreateThirdAppConfig  创建第三方应用配置信息
func (s *cWeiXinThirdAppConfig) CreateThirdAppConfig(ctx context.Context, req *weixin_third_app_v1.CreateThirdAppConfigReq) (*weixin_third_app_v1.ThirdAppConfigRes, error) {
	//return funs.CheckPermission(ctx,
	//	func() (*v1.ThirdAppConfigRes, error) {
	//		ret, err := weixin_service.ThirdAppConfig().CreateThirdAppConfig(ctx, &req.WeixinThirdAppConfig)
	//		return (*v1.ThirdAppConfigRes)(ret), err
	//	},
	//	// 记得添加权限
	//	// weixin_permission.ThirdAppConfig.PermissionType.Update,
	//)

	ret, err := weixin_service.ThirdAppConfig().CreateThirdAppConfig(ctx, &req.WeixinThirdAppConfig)
	return (*weixin_third_app_v1.ThirdAppConfigRes)(ret), err
}

// GetThirdAppConfigByAppId 根据AppId查找第三方应用配置信息
func (s *cWeiXinThirdAppConfig) GetThirdAppConfigByAppId(ctx context.Context, req *weixin_third_app_v1.GetThirdAppConfigByIdReq) (*weixin_third_app_v1.ThirdAppConfigRes, error) {
	return funs.CheckPermission(ctx,
		func() (*weixin_third_app_v1.ThirdAppConfigRes, error) {
			ret, err := weixin_service.ThirdAppConfig().GetThirdAppConfigById(ctx, req.Id)
			return (*weixin_third_app_v1.ThirdAppConfigRes)(ret), err
		},
		// 记得添加权限
		// weixin_permission.ThirdAppConfig.PermissionType.Update,
	)

}

// UpdateAppConfig 修改服务商基础信息
func (s *cWeiXinThirdAppConfig) UpdateAppConfig(ctx context.Context, req *weixin_third_app_v1.UpdateThirdAppConfigReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := weixin_service.ThirdAppConfig().UpdateAppConfig(ctx, &req.UpdateThirdAppConfigReq)
			return ret == true, err
		},
		// 记得添加权限
		// weixin_permission.ThirdAppConfig.PermissionType.Update,
	)
	//
	//ret, err := weixin_service.ThirdAppConfig().UpdateAppConfig(ctx, &req.UpdateThirdAppConfigReq)
	//
	//return ret == true, err
}

// UpdateThirdAppConfigHttps 修改服务商应用Https配置
func (s *cWeiXinThirdAppConfig) UpdateThirdAppConfigHttps(ctx context.Context, req *weixin_third_app_v1.UpdateThirdAppConfigHttpsReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.ThirdAppConfig().UpdateAppConfigHttps(ctx, &req.UpdateThirdAppConfigHttpsReq)
	return ret == true, err
}
