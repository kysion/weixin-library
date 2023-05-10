package merchant

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	v1 "github.com/kysion/weixin-library/api/weixin_v1"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/utility"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

// AlipayAuthUserInfo 用户登录信息

// 获取用户openId

// 获取用户unionID

// 用户网页授权

// 获取用户基本信息

var MerchantService = cMerchantService{}

type cMerchantService struct{}

// UserAuth 用户授权 （代商家管理小程序）  --- 公众号网页授权圈
func (c *cMerchantService) UserAuth(ctx context.Context, _ *weixin_merchant_app_v1.UserAuthReq) (api_v1.StringRes, error) {
	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	if err != nil {
		return "", err
	}
	// 问题：小程序不能正常用户授权，提示此公众号并没有这些scope的权限，错误码:10005，
	//      公众号可以

	// 1.用户授权，拿到登陆凭据code

	// 2.通过code拿到openId和session_key (基本)

	// 3.获取用户信息 (详细)

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

	redirect_url := gurl.Encode("https://www.kuaimk.com/weixin/" + appId + "/gateway.userAuthRes")

	// 2.引导用户进入用户授权页面
	authUrl := "https://open.weixin.qq.com/connect/oauth2/authorize?" +
		"appid=" + appId +
		"&redirect_uri=" + redirect_url +
		"&response_type=code" +
		//"&scope=snsapi_userinfo" +
		"&scope=snsapi_base" +
		"&state=STATE" +
		"#wechat_redirect"
	fmt.Println("用户授权全链接：\n", authUrl)

	g.RequestFromCtx(ctx).Response.Header().Set("referer", "https://www.kauimk.com/weixin/wx56j8q12l89h99/gateway.services") // https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

	g.RequestFromCtx(ctx).Response.Header().Set("Content-Type", "text/html; charset=UTF-8")

	g.RequestFromCtx(ctx).Response.WriteTplContent(`<script>window.location.href = "{{.url}}"</script>`, g.Map{
		"url": authUrl,
	})

	//
	//redirect_url := gurl.Encode("https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.authRes")
	//
	////authURL, err := buildAuthURL(merchantApp.AppCallbackUrl, appId)
	//authURL, err := buildAuthURL(redirect_url, appId)
	//if err != nil {
	//	return "", err
	//}
	//// https:www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback
	//// https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback
	//
	//fmt.Println(authURL)
	//
	//g.RequestFromCtx(ctx).Response.RedirectTo(authURL)

	return "success", nil
}

// UserAuthRes 用户授权接收地址
func (c *cMerchantService) UserAuthRes(ctx context.Context, req *weixin_merchant_app_v1.UserAuthResReq) (v1.StringRes, error) {
	appId := g.RequestFromCtx(ctx).Get("appId").String()
	//appIdLen := len(pathAppId)
	//subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00
	//
	//appId := "wx" + utility.Base32ToHex(subAppId)

	// 1.拿到登陆凭据code
	fmt.Println("appId：" + appId)
	fmt.Println("登陆凭据code：", req.Code) // 登陆凭据code： 011Koy200CX7UP12uD100PgJFk3Koy2P

	// 2.获得用户openId和session_key
	weixin_service.UserAuth().UserAuthCallback(ctx, g.Map{
		"code":   req.Code,
		"app_id": appId,
	})

	return "success", nil
}

// AppAuthReq 应用授权
func (c *cMerchantService) AppAuthReq(ctx context.Context, _ *weixin_merchant_app_v1.AppAuthReq) (v1.StringRes, error) {
	// 通过appId将具体第三方应用配置信息从数据库获取出来

	//appId := g.RequestFromCtx(ctx).Get("appId").String()

	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	//https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.services
	//https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.services
	//
	//https://www.kuaimk.com/weixin/wx534d1a08aa84c529/wx56j8q12l89h99/gateway.callback
	//https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

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

	redirect_url := gurl.Encode("https://www.kuaimk.com/weixin/wx56j8q12l89h99/gateway.authRes")
	//authUrl := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" +
	//	//authUrl := "https://mp.weixin.qq.com/safe/bindcomponent?" +
	//	"component_appid=" + appId +
	//	"&pre_auth_code=" + proAuthCodeRes.PreAuthCode +
	//	"&redirect_url=" + redirect_url

	// 5.引导用户进入授权页面
	authUrl := "https://open.weixin.qq.com/wxaopen/safe/bindcomponent?" +
		"action=bindcomponent&no_scan=1&component_appid=" + appId +
		"&pre_auth_code=" + proAuthCodeRes.PreAuthCode +
		"&redirect_uri=" + redirect_url +
		"&auth_type=3" +
		//"&biz_appid=xxxx" +
		"#wechat_redirect"
	fmt.Println("授权全链接：\n", authUrl)

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

	return "", nil
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
