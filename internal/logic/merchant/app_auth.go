package merchant

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/format_utils"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
)

/*
应用授权
*/

/*
	应用授权通知类型InfoType ：
		authorized 授权成功
		updateauthorized 更新授权
		unauthorized 取消授权
*/

type sAppAuth struct {
}

func init() {
	//weixin_service.RegisterAppAuth(NewAppAuth())
}

func (s *sAppAuth) injectHook() {
	//callHook := weixin_service.Gateway().GetCallbackMsgHook()
	notifyHook := weixin_service.Gateway().GetServiceNotifyTypeHook()

	//callHook.InstallHook(weixin_enum.Info.CallbackType.ComponentAccessToken, s.AppAuth) // 应用授权

	notifyHook.InstallHook(weixin_enum.Info.ServiceType.Authorized, s.Authorized)             // 授权成功
	notifyHook.InstallHook(weixin_enum.Info.ServiceType.UpdateAuthorized, s.UpdateAuthorized) // 授权更新
	notifyHook.InstallHook(weixin_enum.Info.ServiceType.Unauthorized, s.Unauthorized)         // 取消授权
}

func NewAppAuth() weixin_service.IAppAuth {

	result := &sAppAuth{}

	result.injectHook()
	return result
}

// GetAuthorizerAccessToken 获取商家授权应用 authorizer_access_token
func GetAuthorizerAccessToken(ctx context.Context, thirdAppId, componentAccessToken string, merchantAppId, refreshToken string) (bool, error) {
	// 2.获取授权方商家的authorizer_access_token
	queryTokenUrl := "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=" + componentAccessToken
	tokenReq := weixin_model.AuthorizerAccessTokenReq{
		ComponentAppid:         thirdAppId,
		AuthorizerAppid:        merchantAppId,
		AuthorizerRefreshToken: refreshToken,
	}

	tokenReqJson, _ := gjson.Encode(tokenReq)

	tokenRes := g.Client().PostContent(ctx, queryTokenUrl, tokenReqJson)
	tokenResData := weixin_model.AuthorizerAccessTokenRes{}

	_ = gjson.DecodeTo(tokenRes, &tokenResData)
	fmt.Println("商家接口调用Token：", tokenResData.AuthorizerAccessToken)

	// 有了authorizer_access_token就又能调用各种商家的API接口了
	if tokenResData.AuthorizerAccessToken != "" {
		// 将秒数转化为 duration 对象
		tokenReExpiresIn := format_utils.SecondToDuration(tokenResData.ExpiresIn)
		// 计算未来时间
		tokenReTime := gtime.Now().Add(tokenReExpiresIn).Format("Y-m-d H:i:s")
		fmt.Println("增加后的时间：", tokenReTime)

		// 存储authorizer_access_token至数据库
		isFullProxy := 1
		_, err := weixin_service.MerchantAppConfig().UpdateAppAuthToken(ctx, &weixin_model.UpdateMerchantAppAuthToken{
			AppId:        &merchantAppId, // 授权商家应用APPID
			AppAuthToken: &tokenResData.AuthorizerAccessToken,
			//ExpiresIn:    gtime.Now().Add(time.Hour * 2), // token 有效期两小时 （7200 = 2小时）
			ExpiresIn:    gtime.NewFromStr(tokenReTime), // token 有效期两小时 （7200 = 2小时）
			ReExpiresIn:  gtime.NewFromStr(tokenReTime),
			RefreshToken: &tokenResData.AuthorizerRefreshToken,
			ThirdAppId:   &thirdAppId,  // 第三方平台应用appId
			IsFullProxy:  &isFullProxy, // 是否全权待开发
		})

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "token令牌刷新失败", "AppAuth")
		}
	}

	return true, nil
}

// RefreshToken 刷新授权应用的Token
func (s *sAppAuth) RefreshToken(ctx context.Context, merchantAppId, thirdAppId, refreshToken string) (bool, error) {

	//merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, merchantAppId)
	//if err != nil {
	//	return false, err
	//}

	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, thirdAppId)
	if err != nil {
		return false, err
	}

	// 2.获取授权方商家的应用：authorizer_access_token
	return GetAuthorizerAccessToken(ctx, thirdApp.AppId, thirdApp.AppAuthToken, merchantAppId, refreshToken)

}

// AppAuth 应用授权具体服务
func (s *sAppAuth) AppAuth(ctx context.Context, info g.Map) bool {
	fmt.Println("自建商家应用授权的Res处理逻辑：AppAuth------>")
	g.Dump(info)

	from := gmap.NewStrAnyMapFrom(info)
	appId := gconv.String(from.Get("appId"))
	authCode := gconv.String(from.Get("auth_code"))

	thirdData, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, appId)
	if err != nil {
		return false
	}

	// 1.使用授权码获取授权方信息
	authorizerInfoUrl := "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=" + thirdData.AppAuthToken
	authorizerInfoReq := weixin_model.AuthorizerInfoReq{
		ComponentAppid:    thirdData.AppId,
		AuthorizationCode: authCode,
	}
	encode, _ := gjson.Encode(authorizerInfoReq)
	authorizerInfo := g.Client().PostContent(ctx, authorizerInfoUrl, encode)
	authorizerInfoRes := weixin_model.AuthorizationInfoRes{}
	_ = gjson.DecodeTo(authorizerInfo, &authorizerInfoRes)
	fmt.Println("授权方基本信息：", authorizerInfoRes)

	// 2.获取授权方商家的authorizer_access_token

	merchantId := authorizerInfoRes.AuthorizationInfo.AuthorizerAppid
	authorizerRefreshToken := authorizerInfoRes.AuthorizationInfo.AuthorizerRefreshToken

	result, err := GetAuthorizerAccessToken(ctx, thirdData.AppId, thirdData.AppAuthToken, merchantId, authorizerRefreshToken)

	if err != nil || result == false {
		sys_service.SysLogs().ErrorSimple(ctx, err, "应用授权失败", "AppAuth")
	}

	return result
}

// Authorized 授权成功 （应用授权成功微信会推送service一次，但是我们自建授权/:appId/gateway.auth自建授权的req中指定了res的地址，就是/:appId/gateway.authRes， 所以要避免重复处理的情况出现）
func (s *sAppAuth) Authorized(ctx context.Context, info g.Map) bool {
	fmt.Println("AppAuth商家应用授权成功，Authorized----->：")
	g.Dump(info)
	/*
		{
		    "MsgType": "authorized",
		    "info":    {
		        AppId:                        "wxb64aa49959fa359c",
		        CreateTime:                   1718857157,
		        InfoType:                     "authorized",
		        ComponentVerifyTicket:        "",
		        AuthorizerAppid:              "wx983095e89f741d4b",
		        AuthorizationCode:            "queryauthcode@@@q765E-81endY0_aRrgBPked4EekH4OlIpXFZ42i6erV54gGHgn-JRpwX09qXDR3QaKOq930LqvTq5cek4FO7bQ",
		        AuthorizationCodeExpiredTime: "1718860757",
		        PreAuthCode:                  "preauthcode@@@igO97u2IBNxk0Xc4MmFCJC2PZ8qP0A7X3VnIPIB-arKBT4MB9kZDz5MchnKOmarxrCHITRtpz_GsNFBHpO66vg",
		    },
		    "appId":   "wxb64aa49959fa359c",
		}
	*/
	if info["MsgType"] != weixin_enum.Info.ServiceType.Authorized.Code() {
		return false
	}

	// 返回的信息
	data := weixin_model.EventMessageBody{}
	_ = gconv.Struct(info["info"], &data)

	// 授权码 过期时间 authorization_code + 时间
	fmt.Println("授权的商家AppID：", data.AuthorizerAppid)
	fmt.Println("发起授权的第三方应用AppID：", data.AppId)

	// 第三方应用信息
	thirdInfo, _ := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, data.AppId)
	if thirdInfo == nil {
		return false
	}

	// 这是微信推送过来的，不会有 info["thirdInfo"] 的信息
	//thirdInfo := info["thirdInfo"]
	//thirdData := weixin_model.WeixinThirdAppConfig{}
	//gconv.Struct(thirdInfo, &thirdData)

	// 商家应用信息
	//merchantInfo := info["info"]
	//merchantData := entity.WeixinMerchantAppConfig{}
	//gconv.Struct(merchantInfo, &merchantData)

	//https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=ACCESS_TOKEN (获取授权方账号基本信息) (Pass 已过期）
	//https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?access_token=ACCESS_TOKEN (获取授权方账号基本信息)
	//https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=ACCESS_TOKEN (获取授权Token)
	//https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token= （查询授权）

	// 1、获取授权方的帐号基本信息 http请求方式: POST（请使用https协议）可选
	//authorizerInfoUrl := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=" + thirdInfo.AppAuthToken // 此API已过期

	// POST https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?access_token=ACCESS_TOKEN
	authorizerInfoUrl := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?access_token=" + thirdInfo.AppAuthToken
	fmt.Println(authorizerInfoUrl)
	authorizerInfoReq := weixin_model.AuthorizerInfoReq{
		ComponentAppid: thirdInfo.AppId,
		//AuthorizationCode: data.AuthorizationCode,
		AuthorizerAppid: data.AuthorizerAppid,
	}

	encode, _ := gjson.Encode(authorizerInfoReq)
	authorizerInfo := g.Client().PostContent(ctx, authorizerInfoUrl, encode)
	fmt.Println("Authorized-authorizerInfo: ----->")
	g.Dump(authorizerInfo)
	var authorizerInfoRes weixin_model.AuthorizationInfoRes
	_ = gjson.DecodeTo(authorizerInfo, &authorizerInfoRes)
	fmt.Println("授权方基本信息：", authorizerInfoRes)

	if &authorizerInfoRes == nil || &authorizerInfoRes.AuthorizationInfo == nil || authorizerInfoRes.AuthorizationInfo.AuthorizerAppid == "0" {
		return false
	}

	// 2、获取商家接口调用凭据 authorizer_access_token  + authorizer_refresh_token + 时间
	// POST https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN
	// POST https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN
	queryAuthUrl := "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=" + thirdInfo.AppAuthToken
	fmt.Println(queryAuthUrl)
	tokenReq := weixin_model.QueryAuthReq{
		ComponentAppid:    thirdInfo.AppId,
		AuthorizationCode: data.AuthorizationCode,
	}
	tokenReqJson, _ := gjson.Encode(tokenReq)
	tokenRes := g.Client().PostContent(ctx, queryAuthUrl, tokenReqJson)
	fmt.Println("Authorized-tokenRes: ----->")
	g.Dump(tokenRes)
	if tokenRes == "" {
		return false
	}

	tokenResData := weixin_model.AuthorizationInfoRes{}
	_ = gjson.DecodeTo(tokenRes, &tokenResData)
	fmt.Println("商家接口调用Token：", tokenResData.AuthorizationInfo.AuthorizerAccessToken)

	if &tokenResData == nil || tokenResData.AuthorizationInfo.AuthorizerRefreshToken == "" {
		return false
	}

	// 有了authorizer_access_token就又能调用各种商家的API接口了
	if tokenResData.AuthorizationInfo.AuthorizerAccessToken != "" {
		// 将秒数转化为 duration 对象
		tokenReExpiresIn := format_utils.SecondToDuration(tokenResData.AuthorizationInfo.ExpiresIn)
		// 计算未来时间
		tokenReTime := gtime.Now().Add(tokenReExpiresIn).Format("Y-m-d H:i:s")
		fmt.Println("增加后的时间：", tokenReTime)

		// 存储authorizer_access_token至数据库
		isFullProxy := 1
		_, err := weixin_service.MerchantAppConfig().UpdateAppAuthToken(ctx, &weixin_model.UpdateMerchantAppAuthToken{
			AppId:        &tokenResData.AuthorizationInfo.AuthorizerAppid, // 授权商家应用APPID
			AppAuthToken: &tokenResData.AuthorizationInfo.AuthorizerAccessToken,
			//ExpiresIn: gtime.Now().Add(time.Hour * 2), // token 有效期两小时 （7200 = 2小时）
			ExpiresIn:    gtime.NewFromStr(tokenReTime), // token 有效期两小时 （7200 = 2小时）
			ReExpiresIn:  gtime.NewFromStr(tokenReTime),
			RefreshToken: &tokenResData.AuthorizationInfo.AuthorizerRefreshToken,
			ThirdAppId:   &thirdInfo.AppId, // 第三方平台应用appId
			IsFullProxy:  &isFullProxy,
		})

		if err != nil {
			return false
		}
	}

	return true
}

// UpdateAuthorized 授权更新
func (s *sAppAuth) UpdateAuthorized(ctx context.Context, info g.Map) bool {
	return true
}

// Unauthorized 授权取消 （渠道1:解决操作可以在微信公众平台登陆后，然后解除第三方应用的授权、渠道2：...）
func (s *sAppAuth) Unauthorized(ctx context.Context, info g.Map) bool {
	fmt.Println("AppAuth商家应用取消授权，Unauthorized----->：")
	g.Dump(info)
	/*
		{
		    "MsgType": "unauthorized",
		    "info":    {
		        AppId:                        "wxb64aa49959fa359c",
		        CreateTime:                   1718856108,
		        InfoType:                     "unauthorized",
		        ComponentVerifyTicket:        "",
		        AuthorizerAppid:              "wx983095e89f741d4b",
		        AuthorizationCode:            "",
		        AuthorizationCodeExpiredTime: "",
		        PreAuthCode:                  "",
		    },
		    "appId":   "wxb64aa49959fa359c",
		}
	*/

	if info["MsgType"] != weixin_enum.Info.ServiceType.Unauthorized.Code() {
		return false
	}

	// 返回的信息
	data := weixin_model.EventMessageBody{}
	_ = gconv.Struct(info["info"], &data)

	// 授权码 过期时间 authorization_code + 时间
	fmt.Println("解除授权的商家AppID：", data.AuthorizerAppid)
	fmt.Println("发起授权的第三方应用AppID：", data.AppId)

	// 1、解除授权：将原本绑定的服务商第三方应用ID重置为0，全权代开发也重置为0
	_, err := weixin_service.MerchantAppConfig().UpdateAppAuth(ctx, data.AuthorizerAppid, 0, 0)
	if err != nil {
		return false
	}

	return true
}
