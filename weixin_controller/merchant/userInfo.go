package merchant

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/utility"
	"net/url"
)

// UserInfo 网关
var UserInfo = cUserInfo{}

type cUserInfo struct{}

// 构建授权链接
func buildAuthURL(redirectURI, appID string) (string, error) {
	redirectURIEncoded := url.QueryEscape(redirectURI)
	//若该参数被设置为 'snsapi_base'，则只能获取到用户的 openid 和 unionid 等基本信息；若设置为 'snsapi_userinfo' 则可以获取到用户的昵称、头像和性别等完整资料信息。

	authURL := "https://open.weixin.qq.com/connect/oauth2/authorize?" +
		"appid=" + appID +
		"&redirect_uri=" + redirectURIEncoded +
		"&response_type=code" +
		//"&scope=snsapi_userinfo" +
		"&scope=snsapi_base" +
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

	//merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)

	//thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)

	redirect_url := gurl.Encode("https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.userAuthRes")
	authURL, err := buildAuthURL(redirect_url, appId)
	if err != nil {
		return "", err
	}

	//authURL, err := buildAuthURL(merchantApp.AppCallbackUrl, appId)

	// https:www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback
	// https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

	fmt.Println(authURL)

	g.RequestFromCtx(ctx).Response.RedirectTo(authURL)

	return "success", nil
}

// 我们的系统是第三方代调用，那么操作用户授权登陆，AppId是第三方应用的还是商家应用的
