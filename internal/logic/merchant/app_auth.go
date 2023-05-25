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
	"time"
)

// 应用授权
type sAppAuth struct {
}

// gateway 主要用于记录和服务商相关操作

// merchant 主要记录和商家有关，例如一些商家消息的hook注册，

// internal 主要用于拓展SDK所不具备。票据例外

func init() {
	weixin_service.RegisterAppAuth(NewAppAuth())
}

func (s *sAppAuth) injectHook() {
	callHook := weixin_service.Gateway().GetCallbackMsgHook()
	//notifyHook := weixin_service.Gateway().GetServiceNotifyTypeHook()

	callHook.InstallHook(weixin_enum.Info.CallbackType.ComponentAccessToken, s.AppAuth) // 应用授权

	callHook.InstallHook(weixin_enum.Info.CallbackType.Authorized, s.Authorized)             // 授权成功
	callHook.InstallHook(weixin_enum.Info.CallbackType.UpdateAuthorized, s.UpdateAuthorized) // 授权更新
	callHook.InstallHook(weixin_enum.Info.CallbackType.Unauthorized, s.Unauthorized)         // 取消授权
}

func NewAppAuth() *sAppAuth {

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

	gjson.DecodeTo(tokenRes, &tokenResData)
	fmt.Println("商家接口调用Token：", tokenResData.AuthorizerAccessToken)

	// 有了authorizer_access_token就又能调用各种商家的API接口了

	// 将秒数转化为 duration 对象
	tokenReExpiresIn := format_utils.SecondToDuration(tokenResData.ExpiresIn)

	// 计算未来时间
	tokenReTime := gtime.Now().Add(tokenReExpiresIn).Format("Y-m-d H:i:s")
	fmt.Println("增加后的时间：", tokenReTime)

	if tokenResData.AuthorizerAccessToken != "" {
		// 存储authorizer_access_token至数据库
		_, err := weixin_service.MerchantAppConfig().UpdateAppAuthToken(ctx, &weixin_model.UpdateMerchantAppAuthToken{
			AppId:        merchantAppId,
			AppAuthToken: tokenResData.AuthorizerAccessToken,
			//expiresIn := gtime.NewFromTimeStamp(tokenResData.ExpiresIn)
			ExpiresIn:    gtime.Now().Add(time.Hour * 2), // token 有效期两小时
			ReExpiresIn:  gtime.NewFromStr(tokenReTime),
			RefreshToken: tokenResData.AuthorizerRefreshToken,
			ThirdAppId:   thirdAppId, // 第三方应用appId
		})

		if err != nil {
			return false, sys_service.SysLogs().ErrorSimple(ctx, err, "token令牌刷新失败", "AppAuth")
		}
	}

	return true, nil
}

// RefreshToken 刷新Token
func (s *sAppAuth) RefreshToken(ctx context.Context, merchantAppId, thirdAppId, refreshToken string) (bool, error) {

	//merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, merchantAppId)
	//if err != nil {
	//	return false, err
	//}

	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, thirdAppId)
	if err != nil {
		return false, err
	}

	// 2.获取授权方商家的authorizer_access_token
	return GetAuthorizerAccessToken(ctx, thirdApp.AppId, thirdApp.AppAuthToken, merchantAppId, refreshToken)

}

// AppAuth 应用授权具体服务
func (s *sAppAuth) AppAuth(ctx context.Context, info g.Map) bool {
	//getComponentAccessToken(ctx, gconv.String(info))
	from := gmap.NewStrAnyMapFrom(info)

	appId := gconv.String(from.Get("appId"))
	authCode := gconv.String(from.Get("auth_code"))
	//expiresIn := gconv.String(from.Get("expires_in"))

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
	gjson.DecodeTo(authorizerInfo, &authorizerInfoRes)
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

// Authorized 授权成功
func (s *sAppAuth) Authorized(ctx context.Context, info g.Map) bool {
	if info["MsgType"] != weixin_enum.Info.CallbackType.Authorized.Code() {
		return false
	}

	// 返回的信息
	data := weixin_model.EventMessageBody{}
	gconv.Struct(info["info"], &data)

	// 授权码 过期时间 authorization_code + 时间
	fmt.Println("商家AppID", data.AppId)
	fmt.Println("授权码：", data.AuthorizationCode)

	// 第三方应用信息
	thirdInfo := info["thirdInfo"]
	thirdData := weixin_model.WeixinThirdAppConfig{}
	gconv.Struct(thirdInfo, &thirdData)

	// 商家应用信息
	//merchantInfo := info["info"]
	//merchantData := entity.WeixinMerchantAppConfig{}
	//gconv.Struct(merchantInfo, &merchantData)

	//https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token= (获取授权方账号基本信息)
	//https: //api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=ACCESS_TOKEN (获取授权Token)
	//https: //api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token= （查询授权）

	// 获取授权方的帐号基本信息 http请求方式: POST（请使用https协议）可选
	authorizerInfoUrl := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=" + thirdData.AppAuthToken
	authorizerInfoReq := weixin_model.AuthorizerInfoReq{
		ComponentAppid:    thirdData.AppId,
		AuthorizationCode: data.AuthorizationCode,
	}
	encode, _ := gjson.Encode(authorizerInfoReq)
	authorizerInfo := g.Client().PostContent(ctx, authorizerInfoUrl, encode)
	authorizerInfoRes := weixin_model.AuthorizationInfoRes{}
	gjson.DecodeTo(authorizerInfo, &authorizerInfoRes)
	fmt.Println("授权方基本信息：", authorizerInfoRes)

	// 获取商家接口调用凭据 authorizer_access_token  + authorizer_refresh_token + 时间
	//	POST https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN
	queryAuthUrl := "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=" + thirdData.AppAuthToken

	tokenReq := weixin_model.QueryAuthReq{
		ComponentAppid:    thirdData.AppId,
		AuthorizationCode: data.AuthorizationCode,
	}

	tokenReqJson, _ := gjson.Encode(tokenReq)

	tokenRes := g.Client().PostContent(ctx, queryAuthUrl, tokenReqJson)

	tokenResData := weixin_model.AuthorizationInfoRes{}
	gjson.DecodeTo(tokenRes, &tokenResData)
	fmt.Println("商家接口调用Token：", tokenResData.AuthorizationInfo.AuthorizerAccessToken)

	// 有了authorizer_access_token就又能调用各种商家的API接口了

	// 存储authorizer_access_token至数据库
	_, err := weixin_service.MerchantAppConfig().UpdateAppAuthToken(ctx, &weixin_model.UpdateMerchantAppAuthToken{
		AppId:        data.AppId,
		AppAuthToken: tokenResData.AuthorizationInfo.AuthorizerAccessToken,
		ExpiresIn:    gtime.New(tokenResData.AuthorizationInfo.ExpiresIn),
		//ReExpiresIn:  gtime.New(tokenResData.AuthorizationInfo.AuthorizerRefreshToken),
		RefreshToken: tokenResData.AuthorizationInfo.AuthorizerRefreshToken,
	})

	if err != nil {
		return false
	}

	return true
}

// UpdateAuthorized 授权更新
func (s *sAppAuth) UpdateAuthorized(ctx context.Context, info g.Map) bool {
	return true
}

// Unauthorized 授权取消
func (s *sAppAuth) Unauthorized(ctx context.Context, info g.Map) bool {
	return true
}
