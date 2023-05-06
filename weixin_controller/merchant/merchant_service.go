package merchant

import (
	"context"
	"fmt"
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

// AppAuthReq 应用授权
func (s *cMerchantService) AppAuthReq(ctx context.Context, _ *weixin_merchant_app_v1.AppAuthReq) (v1.StringRes, error) {
	// 通过appId将具体第三方应用配置信息从数据库获取出来

	//appId := g.RequestFromCtx(ctx).Get("appId").String()

	pathAppId := g.RequestFromCtx(ctx).Get("appId").String()
	appIdLen := len(pathAppId)
	subAppId := gstr.SubStr(pathAppId, 2, appIdLen) // caf4b7b8d6620f00

	appId := "wx" + utility.Base32ToHex(subAppId)

	app, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)

	// 4.获取与授权码
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

	// 5.引导用户进入授权页面
	redirect_url := gurl.Encode(app.AppCallbackUrl)
	authUrl := "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" +
		//authUrl := "https://mp.weixin.qq.com/safe/bindcomponent?" +
		"component_appid=" + appId +
		"&pre_auth_code=" + proAuthCodeRes.PreAuthCode +
		"&redirect_url=" + redirect_url
	fmt.Println("授权全链接：\n", authUrl)

	g.RequestFromCtx(ctx).Response.Header().Set("referer", "https://www.kauimk.com/weixin/wx56j8q12l89h99/gateway.services") // https://www.kuaimk.com/weixin/$APPID$/wx56j8q12l89h99/gateway.callback

	// g.RequestFromCtx(ctx).Response.RedirectTo(authUrl) // 会报错：说请确认授权入口页所在域名和授权回调页所在域名相同

	g.RequestFromCtx(ctx).Response.Header().Set("Content-Type", "text/html; charset=UTF-8")
	g.RequestFromCtx(ctx).Response.WriteTplContent(`<html lang="zh"><head><meta charset="utf-8"></head><body>测试页面：<a href="{{.url}}">{{.label}}</a></body></html>`, g.Map{
		"url":   authUrl,
		"label": "授权",
	})

	return "", nil
}
