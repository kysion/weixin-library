package weixin_controller

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/kysion/weixin-library/api/weixin_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

var WeiXinThirdAppConfig = cWeiXinThirdAppConfig{}

type cWeiXinThirdAppConfig struct{}

// UpdateState 修改状态
func (s *cWeiXinThirdAppConfig) UpdateState(ctx context.Context, req *v1.UpdateStateReq) (api_v1.BoolRes, error) {
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
func (s *cWeiXinThirdAppConfig) CreateThirdAppConfig(ctx context.Context, req *v1.CreateThirdAppConfigReq) (*v1.ThirdAppConfigRes, error) {
	//return funs.CheckPermission(ctx,
	//	func() (*v1.ThirdAppConfigRes, error) {
	//		ret, err := weixin_service.ThirdAppConfig().CreateThirdAppConfig(ctx, &req.WeixinThirdAppConfig)
	//		return (*v1.ThirdAppConfigRes)(ret), err
	//	},
	//	// 记得添加权限
	//	// weixin_permission.ThirdAppConfig.PermissionType.Update,
	//)

	ret, err := weixin_service.ThirdAppConfig().CreateThirdAppConfig(ctx, &req.WeixinThirdAppConfig)
	return (*v1.ThirdAppConfigRes)(ret), err
}

// GetThirdAppConfigByAppId 根据AppId查找第三方应用配置信息
func (s *cWeiXinThirdAppConfig) GetThirdAppConfigByAppId(ctx context.Context, req *v1.GetThirdAppConfigByIdReq) (*v1.ThirdAppConfigRes, error) {
	return funs.CheckPermission(ctx,
		func() (*v1.ThirdAppConfigRes, error) {
			ret, err := weixin_service.ThirdAppConfig().GetThirdAppConfigById(ctx, req.Id)
			return (*v1.ThirdAppConfigRes)(ret), err
		},
		// 记得添加权限
		// weixin_permission.ThirdAppConfig.PermissionType.Update,
	)

}

// UpdateAppConfig 修改服务商基础信息
func (s *cWeiXinThirdAppConfig) UpdateAppConfig(ctx context.Context, req *v1.UpdateThirdAppConfigReq) (api_v1.BoolRes, error) {
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

// AppAuthReq 应用授权
func (s *cWeiXinThirdAppConfig) AppAuthReq(ctx context.Context, _ *v1.AppAuthReq) (v1.StringRes, error) {
	// 通过appId将具体第三方应用配置信息从数据库获取出来

	appId := g.RequestFromCtx(ctx).Get("appId").String()
	app, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	// 4.获取与授权码
	proAuthCodeReq := weixin_model.ProAuthCodeReq{
		ComponentAppid: appId,
		// ComponentAccessToken: token,  // 不能写json结构体里面，一半数据写在上面url上，一半数据写在json结构体
	}
	encode, _ := gjson.Encode(proAuthCodeReq)
	proAuthCodeUrl := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + app.AppAuthToken
	fmt.Println(string(encode))

	proAuthCode := g.Client().PostContent(ctx, proAuthCodeUrl, encode)
	proAuthCodeRes := weixin_model.ProAuthCodeRes{}
	gjson.DecodeTo(proAuthCode, &proAuthCodeRes)
	/*
		{
			"pre_auth_code": "preauthcode@@@pxvu7JW0hDQqNf38HcEXF6ejB4pnzVnA_GXlqqb1XcSmS3GjEhy-TfJOIqjAODk3MmmTZpNHi7Brgc_ugz0RCg",
			"expires_in": 1800
		}
	*/

	// 5.引导用户进入授权页面
	redirect_url := gurl.Encode("https://weixin.jditco.com/weixin/$APPID$/gateway.callback")
	authUrl := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" +
		//authUrl := "https://mp.weixin.qq.com/safe/bindcomponent?" +
		"component_appid=" + appId +
		"&pre_auth_code=" + proAuthCodeRes.PreAuthCode +
		"&redirect_url=" + redirect_url
	fmt.Println("授权全链接：\n", authUrl)

	g.RequestFromCtx(ctx).Response.Header().Set("referer", "https://douyin.jditco.com/weixin/gateway.services")

	g.RequestFromCtx(ctx).Response.RedirectTo(authUrl)
	//r.Response.Header().Set("Content-Type", "text/html; charset=UTF-8")
	//r.Response.WriteTplContent(`<html lang="zh"><head><meta charset="utf-8"></head><body>测试页面：<a href="{{.url}}">{{.label}}</a></body></html>`, g.Map{
	//	"url":   authUrl,
	//	"label": "授权",
	//})

	return "", nil
}

// UpdateThirdAppConfigHttps 修改服务商应用Https配置
func (s *cWeiXinThirdAppConfig) UpdateThirdAppConfigHttps(ctx context.Context, req *v1.UpdateThirdAppConfigHttpsReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.ThirdAppConfig().UpdateAppConfigHttps(ctx, &req.UpdateThirdAppConfigHttpsReq)
	return ret == true, err
}
