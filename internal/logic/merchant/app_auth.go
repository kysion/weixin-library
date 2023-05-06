package merchant

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
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

// AppAuth 应用授权具体服务
func (s *sAppAuth) AppAuth(ctx context.Context, info g.Map) bool {
	//getComponentAccessToken(ctx, gconv.String(info))
	// TODO 需要补充应用授权相关代码
	return false
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

	// 获取授权方的帐号基本信息 http请求方式: POST（请使用https协议）可选
	authorizerInfoUrl := "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=" + thirdData.AppAuthToken
	authorizerInfoReq := weixin_model.AuthorizerInfoReq{
		ComponentAppid:    thirdData.AppId,
		AuthorizationCode: "授权码",
	}
	encode, _ := gjson.Encode(authorizerInfoReq)
	g.Client().PostContent(ctx, authorizerInfoUrl, encode)

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
	fmt.Println("商家接口调用Token：", tokenResData.AuthorizerAccessToken)

	// 有了authorizer_access_token就又能调用各种商家的API接口了

	// 存储authorizer_access_token至数据库
	_, err := weixin_service.MerchantAppConfig().UpdateAppAuthToken(ctx, &weixin_model.UpdateMerchantAppAuthToken{
		AppId:        data.AppId,
		AppAuthToken: tokenResData.AuthorizerAccessToken,
		ExpiresIn:    gtime.New(tokenResData.ExpiresIn),
		//ReExpiresIn:  gtime.New(tokenResData.AuthorizerRefreshToken),
		RefreshToken: tokenResData.AuthorizerRefreshToken,
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
