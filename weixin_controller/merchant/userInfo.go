package merchant

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/utility"
	"github.com/kysion/weixin-library/weixin_service"
	"net/url"
)

// UserInfo 网关
var UserInfo = cUserInfo{}

type cUserInfo struct{}

// 构建授权链接
func buildAuthURL(redirectURI, appID string) (string, error) {
	redirectURIEncoded := url.QueryEscape(redirectURI)

	authURL := "https://open.weixin.qq.com/connect/oauth2/authorize?" +
		"appid=" + appID +
		"&redirect_uri=" + redirectURIEncoded +
		"&response_type=code" +
		"&scope=snsapi_userinfo" +
		"&state=STATE" +
		"#wechat_redirect"

	return authURL, nil
}

// GetUserInfo 获取微信用户信息
func (c *cUserInfo) GetUserInfo(ctx context.Context, _ *weixin_merchant_app_v1.GetUserInfoReq) (api_v1.StringRes, error) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	app, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)

	//app, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	if err != nil {
		return "", err
	}

	authURL, err := buildAuthURL(app.AppCallbackUrl, appId)

	// https:www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback
	// https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

	fmt.Println(authURL)

	g.RequestFromCtx(ctx).Response.RedirectTo(authURL)

	return "success", nil
}

// 我们的系统是第三方代调用，那么操作用户授权登陆，AppId是第三方应用的还是商家应用的

// UserAuth 用户授权
//func (c *cUserInfo) UserAuth(ctx context.Context, _ *weixin_merchant_app_v1.GetUserInfoReq) (api_v1.StringRes, error) {
//	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
//	appIdLen := len(pathAppId)
//	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00
//
//	appId := "wx" + utility.Base32ToHex(subAppId)
//
//	app, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
//
//	//app, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)
//
//	if err != nil {
//		return "", err
//	}
//
//	authURL, err := buildAuthURL(app.AppCallbackUrl, appId)
//
//	// https:www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback
//	// https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback
//
//	fmt.Println(authURL)
//
//	g.RequestFromCtx(ctx).Response.RedirectTo(authURL)
//
//	return "success", nil
//}
