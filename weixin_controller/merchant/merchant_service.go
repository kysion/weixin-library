package merchant

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	v1 "github.com/kysion/weixin-library/api/weixin_v1"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/utility"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"net/url"
)

// AlipayAuthUserInfo 用户登录信息

// 获取用户openId

// 获取用户unionID

// 用户网页授权

// 获取用户基本信息

var MerchantService = cMerchantService{}

type cMerchantService struct{}

// 构建用户网页授权链接
func buildUserAuthURL(redirectURI, appID string) (string, error) {
	redirectURIEncoded := url.QueryEscape(redirectURI)
	//若该参数被设置为 'snsapi_base'，则只能获取到用户的 openid 和 unionid 等基本信息；若设置为 'snsapi_userinfo' 则可以获取到用户的昵称、头像和性别等完整资料信息。

	authUrl := "https://open.weixin.qq.com/connect/oauth2/authorize?" +
		"appid=" + appID +
		"&redirect_uri=" + redirectURIEncoded +
		"&response_type=code" +
		"&scope=snsapi_userinfo" + // 手动授权 （需要获取useInfo的情况）
		//"&scope=snsapi_base" + //  静默授权 （用户直接login的情况）
		"&state=STATE" +
		"#wechat_redirect"

	return authUrl, nil
}

// 构建应用授权链接
func buildAppAuthURL(redirectURI, appID, preAuthCode string) (string, error) {
	redirectURIEncoded := url.QueryEscape(redirectURI)
	//若该参数被设置为 'snsapi_base'，则只能获取到用户的 openid 和 unionid 等基本信息；若设置为 'snsapi_userinfo' 则可以获取到用户的昵称、头像和性别等完整资料信息。

	authUrl := "https://open.weixin.qq.com/wxaopen/safe/bindcomponent?" +
		"action=bindcomponent" +
		"&no_scan=1" +
		"&component_appid=" + appID +
		"&pre_auth_code=" + preAuthCode +
		"&redirect_uri=" + redirectURIEncoded +
		"&auth_type=3" +
		//"&biz_appid=xxxx" + // 目标appId，可不填
		"#wechat_redirect"

	return authUrl, nil
}

// UserAuth 用户授权 （代商家管理小程序）  --- 公众号网页授权方式
func (c *cMerchantService) UserAuth(ctx context.Context, _ *weixin_merchant_app_v1.UserAuthReq) (api_v1.StringRes, error) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return "", err
	}

	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	if err != nil {
		return "", err
	}

	// 问题：小程序不能正常用户授权，提示此公众号并没有这些scope的权限，错误码:10005，小程序需要使用wx.login方式
	//      公众号可以，因为这就是公众号网页授权方式

	// 1.用户授权，拿到登陆凭据code

	// 2.通过code拿到获得openId和accessKey (基本)

	// 3.获取用户信息userInfo (详细)

	// 1.获取预授权码
	proAuthCodeReq := weixin_model.ProAuthCodeReq{
		ComponentAppid: thirdApp.AppId,
		// ComponentAccessToken: token,  // 不能写json结构体里面，一半数据写在上面url上，一半数据写在json结构体
	}
	encode, _ := gjson.Encode(proAuthCodeReq)
	proAuthCodeUrl := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + thirdApp.AppAuthToken

	proAuthCode := g.Client().PostContent(ctx, proAuthCodeUrl, encode)
	proAuthCodeRes := weixin_model.ProAuthCodeRes{}
	gjson.DecodeTo(proAuthCode, &proAuthCodeRes)

	redirect_url := "https://www.kuaimk.com/weixin/" + appId + "/gateway.userAuthRes"

	//login_url := gurl.Encode("https://www.kuaimk.com/weixin/" + appId + "/userLogin")

	// 2.引导用户进入用户授权页面
	authUrl, _ := buildUserAuthURL(redirect_url, appId)
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n用户授权全链接： "+authUrl, "cUserAuth")

	g.RequestFromCtx(ctx).Response.Header().Set("referer", "https://www.kauimk.com/weixin/wx56j8q12l89h99/gateway.services") // https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

	g.RequestFromCtx(ctx).Response.Header().Set("Content-Type", "text/html; charset=UTF-8")

	g.RequestFromCtx(ctx).Response.WriteTplContent(`<script>window.location.href = "{{.url}}"</script>`, g.Map{
		"url": authUrl,
	})

	//g.RequestFromCtx(ctx).Response.WriteTplContent(`<html lang="zh"><head><meta charset="utf-8"></head><body>测试页面：<a href="{{.url}}">{{.label}}</a></body></html>`, g.Map{
	//	"url":   authUrl,
	//	"label": "用户授权",
	//})

	return "success", nil
}

// UserAuthRes 用户授权接收地址
func (c *cMerchantService) UserAuthRes(ctx context.Context, req *weixin_merchant_app_v1.UserAuthResReq) (v1.StringRes, error) {
	appId := g.RequestFromCtx(ctx).Get("appId").String()

	// 1.拿到登陆凭据code
	fmt.Println("appId：" + appId)
	fmt.Println("登陆凭据code：", req.Code) // 登陆凭据code： 011Koy200CX7UP12uD100PgJFk3Koy2P

	// 2.处理授权回调请求，获得openId和accessKey和userInfo
	weixin_service.UserAuth().UserAuthCallback(ctx, g.Map{ // 用户授权 （网页授权方式）
		"code":   req.Code,
		"app_id": appId,
		// "sys_user_id": 0,
		// "merchant_id": 0,
	})

	return "success", nil
}

// UserLogin 用户授权登录  （小程序登录wx.login方式）
func (c *cMerchantService) UserLogin(ctx context.Context, req *weixin_merchant_app_v1.UserLoginReq) (v1.StringRes, error) {
	appId := g.RequestFromCtx(ctx).Get("appId").String()

	// 1.拿到登陆凭据code
	fmt.Println("appId：" + appId)
	fmt.Println("登陆凭据code：", req.Code) // 登陆凭据code： 011Koy200CX7UP12uD100PgJFk3Koy2P

	// 2.获得用户openId和session_key
	weixin_service.UserAuth().UserLogin(ctx, g.Map{
		"code":   req.Code,
		"app_id": appId,
	})

	return "success", nil
}

// RefreshToken 刷新Token
func (c *cMerchantService) RefreshToken(ctx context.Context, _ *weixin_merchant_app_v1.RefreshTokenReq) (api_v1.BoolRes, error) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return false, err
	}

	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)

	if err != nil {
		return false, err
	}

	ret, err := weixin_service.AppAuth().RefreshToken(ctx, appId, thirdApp.AppId)

	return ret == true, err
}

// AppAuthReq 应用授权
func (c *cMerchantService) AppAuthReq(ctx context.Context, _ *weixin_merchant_app_v1.AppAuthReq) (v1.StringRes, error) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	app, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	// 4.获取预授权码
	proAuthCodeReq := weixin_model.ProAuthCodeReq{
		ComponentAppid: appId,
		// ComponentAccessToken: token,  // 不能写json结构体里面，一半数据写在上面url上，一半数据写在json结构体
	}
	encode, _ := gjson.Encode(proAuthCodeReq)
	proAuthCodeUrl := "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=" + app.AppAuthToken

	proAuthCode := g.Client().PostContent(ctx, proAuthCodeUrl, encode)
	proAuthCodeRes := weixin_model.ProAuthCodeRes{}
	gjson.DecodeTo(proAuthCode, &proAuthCodeRes)
	/*
		{
			"pre_auth_code": "preauthcode@@@pxvu7JW0hDQqNf38HcEXF6ejB4pnzVnA_GXlqqb1XcSmS3GjEhy-TfJOIqjAODk3MmmTZpNHi7Brgc_ugz0RCg",
			"expires_in": 1800
		}
	*/

	//redirect_url := gurl.Encode("https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.authRes")
	redirect_url := "https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.authRes"

	//authUrl := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" +

	// 5.引导用户进入授权页面
	authUrl, _ := buildAppAuthURL(redirect_url, appId, proAuthCodeRes.PreAuthCode)
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n应用授权全链接： "+authUrl, "cAppAuth")

	g.RequestFromCtx(ctx).Response.Header().Set("referer", "https://www.kauimk.com/weixin/wx56j8q12l89h99/gateway.services") // https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

	// g.RequestFromCtx(ctx).Response.RedirectTo(authUrl) // 会报错：说请确认授权入口页所在域名和授权回调页所在域名相同

	g.RequestFromCtx(ctx).Response.Header().Set("Content-Type", "text/html; charset=UTF-8")
	//g.RequestFromCtx(ctx).Response.WriteTplContent(`<html lang="zh"><head><meta charset="utf-8"></head><body>测试页面：<a href="{{.url}}">{{.label}}</a></body></html>`, g.Map{
	//	"url":   authUrl,
	//	"label": "授权",
	//})
	g.RequestFromCtx(ctx).Response.WriteTplContent(`<script>window.location.href = "{{.url}}"</script>`, g.Map{
		"url": authUrl,
	})

	return "success", nil
}

// AuthRes 商家应用授权变更等消息推送
func (c *cMerchantService) AuthRes(ctx context.Context, req *weixin_merchant_app_v1.AuthResReq) (v1.StringRes, error) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	fmt.Println("认证code：", req.AuthCode)

	fmt.Println("推送消息：", req.ExpiresIn)

	weixin_service.AppAuth().AppAuth(ctx, g.Map{
		"appId":      appId,
		"auth_code":  req.AuthCode,
		"expires_in": req.ExpiresIn,
	})

	return "success", nil
}
