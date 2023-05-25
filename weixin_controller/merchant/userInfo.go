package merchant

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
	"net/url"
)

// UserInfo 网关
var UserInfo = cUserInfo{}

type cUserInfo struct{}

// 构建授权链接
func buildUserInfoURL(redirectURI, appID string) (string, error) {
	redirectURIEncoded := url.QueryEscape(redirectURI)
	//若该参数被设置为 'snsapi_base'，则只能获取到用户的 openid 和 unionid 等基本信息；若设置为 'snsapi_userinfo' 则可以获取到用户的昵称、头像和性别等完整资料信息。

	authURL := "" + redirectURIEncoded

	return authURL, nil
}

// GetUserInfo 获取微信用户信息
func (c *cUserInfo) GetUserInfo(ctx context.Context, _ *weixin_merchant_app_v1.GetUserInfoReq) (api_v1.StringRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	//merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)

	//thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)

	redirect_url := gurl.Encode("https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.userAuthRes")
	authURL, err := buildUserInfoURL(redirect_url, appId)
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

// GetTinyAppUserInfo 小程序通过encryptedData 获取用户信息
func (c *cUserInfo) GetTinyAppUserInfo(ctx context.Context, req *weixin_merchant_app_v1.GetTinyAppUserInfoReq) (*weixin_model.UserInfoRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	// 拿到当前登录用户的token,其实就是JwtToken，登录了就有JwtToken，然后拿到sysUserId，从而拿到session_key
	user := sys_service.SysSession().Get(ctx).JwtClaimsUser

	wConsumer, err := weixin_service.Consumer().GetConsumerBySysUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	// weixin_service.Consumer().GetConsumerByOpenId(ctx, req.OpenId)

	res, err := weixin_service.UserAuth().GetTinyAppUserInfo(ctx, wConsumer.SessionKey, req.EncryptedData, req.IV, appId, wConsumer.OpenId)

	return res, err
}
