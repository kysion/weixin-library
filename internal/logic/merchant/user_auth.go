package merchant

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/pay-share-library/pay_model/pay_enum"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	hook "github.com/kysion/weixin-library/weixin_model/weixin_hook"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/yitter/idgenerator-go/idgen"
	"io/ioutil"
	"net/http"
)

// 用户授权 （静默、手动授权）

/*
   	1、构建授权连接，在回调中拿到code
	2、通过code拿到接口调用凭据access_token
	3、通过access_token拿到用户信息user_info
	4、通过refresh_token 进行刷新access_token
*/

type sUserAuth struct {
	// 消费者Hook
	ConsumerHook base_hook.BaseHook[hook.ConsumerKey, hook.ConsumerHookFunc]
}

func init() {
	weixin_service.RegisterUserAuth(NewUserAuth())
}

func NewUserAuth() *sUserAuth {

	result := &sUserAuth{}

	result.injectHook()
	return result
}

func (s *sUserAuth) injectHook() {
	//notifyHook := weixin_service.Gateway().GetServiceNotifyTypeHook()
	callHook := weixin_service.Gateway().GetCallbackMsgHook()

	callHook.InstallHook(weixin_enum.Info.CallbackType.UserAuth, s.UserAuthCallback)
}

func (s *sUserAuth) InstallConsumerHook(infoType hook.ConsumerKey, hookFunc hook.ConsumerHookFunc) {
	sys_service.SysLogs().InfoSimple(context.Background(), nil, "\n-------订阅sUserAuth-Hook： ------- ", "sUserAuth")

	s.ConsumerHook.InstallHook(infoType, hookFunc)
}

func (s *sUserAuth) GetHook() base_hook.BaseHook[hook.ConsumerKey, hook.ConsumerHookFunc] {
	return s.ConsumerHook
}

// UserAuthCallback 处理网页授权回调请求 （公众号登录）
func (s *sUserAuth) UserAuthCallback(ctx context.Context, info g.Map) bool {
	from := gmap.NewStrAnyMapFrom(info)

	// 1.拿到code
	code := gconv.String(from.Get("code"))
	appId := gconv.String(from.Get("app_id"))
	sysUserId := gconv.Int64(from.Get("sys_user_id"))
	merchantId := gconv.Int64(from.Get("merchant_id"))

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil || merchantApp == nil {
		return false
	}
	// TODO 获取用户信息
	//sysUser, err := sys_service.SysUser().GetSysUserById(ctx, sysUserId)
	//if err != nil {
	//	return err
	//}

	// 2.获取access_token  (能拿到openId和access_token)
	accessToken, err := getAccessToken(code, appId, merchantApp.AppSecret)
	if err != nil {
		sys_service.SysLogs().ErrorSimple(ctx, err, "获取AccessToken失败："+err.Error(), "WeiXin")
		return false
	}
	sys_service.SysLogs().InfoSimple(ctx, nil, "\nOpenId："+accessToken.Openid, "sUserAuth")
	sys_service.SysLogs().InfoSimple(ctx, nil, "\nAccessToken： "+accessToken.AccessToken, "sUserAuth")

	openID := accessToken.Openid
	err = dao.WeixinConsumerConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if openID != "" {

			// 3.获取用户信息userInfo
			userInfo, err := getUserInfo(accessToken.AccessToken, openID)
			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "获取用户信息失败："+err.Error(), "WeiXin")
			}

			var weiXinConsumer *weixin_model.WeixinConsumerConfig
			if accessToken.Unionid != "" {
				weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenId(ctx, openID, accessToken.Unionid)
			} else {
				weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenId(ctx, openID)
			}

			// 4.处理微信消费者数据
			if weiXinConsumer == nil { // 创建
				wConsumerData := kconv.Struct(userInfo, &weixin_model.WeixinConsumerConfig{})

				wConsumerData.OpenId = openID
				wConsumerData.SysUserId = sysUserId // TODO 后期想办法将sysUserId传递
				wConsumerData.UserType = 0          // TODO 后期通过sysUserId拿到user，拿到type
				wConsumerData.UserState = 0         // TODO 用户状态,拿到user，拿到type
				wConsumerData.Avatar = userInfo.HeadImgURL

				wConsumerData.AccessToken = accessToken.AccessToken
				wConsumerData.UnionId = accessToken.Unionid
				wConsumerData.OpenId = accessToken.Openid
				wConsumerData.SessionKey = "" // TODO 获取用户信息的时候补全sessionKey

				_, err = weixin_service.Consumer().CreateConsumer(ctx, wConsumerData)
				if err != nil {
					return err
				}

			} else { // 修改
				wConsumerData := kconv.Struct(userInfo, &weixin_model.UpdateConsumerReq{}) // 修改用户基本数据
				wConsumerData.Avatar = userInfo.HeadImgURL
				_, err = weixin_service.Consumer().UpdateConsumer(ctx, weiXinConsumer.Id, wConsumerData)
				if err != nil {
					return err
				}

				_, err = weixin_service.Consumer().UpdateConsumerToken(ctx, openID, &weixin_model.UpdateConsumerTokenReq{ // 修改用户Token
					AccessToken: &accessToken.AccessToken,
					SessionKey:  nil,
				})
				if err != nil {
					return err
				}
			}

			// 5.存储kmk_consumer消费者数据
			g.Try(ctx, func(ctx context.Context) {
				s.ConsumerHook.Iterator(func(key hook.ConsumerKey, value hook.ConsumerHookFunc) {
					if key.ConsumerAction.Code() == weixin_enum.Consumer.ActionEnum.Auth.Code() && key.Category.Code() == weixin_enum.Consumer.Category.Consumer.Code() { // 如果订阅者是订阅授权,并且是操作kmk_consumer表
						g.Try(ctx, func(ctx context.Context) {
							data := hook.UserInfo{
								SysUserId:   sysUserId, // (消费者id = sys_User_id)
								UserInfoRes: *userInfo,
							}
							sys_service.SysLogs().InfoSimple(ctx, nil, "\n广播-------存储消费者数据 kmk-consumer", "sUserAuth")

							value(ctx, data)
						})
					}
				})
			})

			// 6.存储第三方应用和用户关系记录  plat_form_user
			g.Try(ctx, func(ctx context.Context) {
				s.ConsumerHook.Iterator(func(key hook.ConsumerKey, value hook.ConsumerHookFunc) { // 这会同时走两个Hook，kmk_consumer  + platform_user,所以加上了category类型
					if key.ConsumerAction.Code() == weixin_enum.Consumer.ActionEnum.Auth.Code() && key.Category.Code() == weixin_enum.Consumer.Category.PlatFormUser.Code() { // 如果订阅者是订阅授权
						g.Try(ctx, func(ctx context.Context) {
							platformUser := entity.PlatformUser{
								Id:            idgen.NextId(),
								FacilitatorId: 0,
								OperatorId:    0,
								EmployeeId:    sysUserId,                                    // EmployeeId  == consumerId == sysUserId   三者相等
								MerchantId:    merchantId,                                   // 商家id，就是消费者首次扫码的商家
								Platform:      pay_enum.Order.TradeSourceType.Weixin.Code(), // 平台类型：1支付宝、2微信、4抖音、8银联
								ThirdAppId:    merchantApp.ThirdAppId,
								MerchantAppId: merchantApp.AppId,
								UserId:        openID, // 平台账户唯一标识
								Type:          0,      // TODO 后期通过sysUserId拿到user，拿到type
							}

							sys_service.SysLogs().InfoSimple(ctx, nil, "\n广播-------存储第三方应用和用户关系记录 kmk-plat_form_user", "sMerchantService")

							value(ctx, platformUser) // 调用Hook
						})
					}
				})
			})

		} else {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "缺少OpenID参数："+err.Error(), "WeiXin")
		}

		return nil
	})

	if err != nil {
		return false
	}

	return true
}

// 获取微信AccessToken
func getAccessToken(code string, appID, appSecret string) (*weixin_model.AccessTokenRes, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?" +
		"appid=" + appID +
		"&secret=" + appSecret +
		"&code=" + code +
		"&grant_type=authorization_code"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	accessToken := weixin_model.AccessTokenRes{}
	if err := json.Unmarshal(body, &accessToken); err != nil {
		return nil, err
	}

	if accessToken.AccessToken == "" || accessToken.ErrCode != 0 {
		return nil, fmt.Errorf("获取AccessToken失败：%s", accessToken.ErrMsg)
	}

	return &accessToken, nil
}

// 获取微信用户信息
func getUserInfo(accessToken string, openID string) (*weixin_model.UserInfoRes, error) {
	// GET https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN

	url := "https://api.weixin.qq.com/sns/userinfo?" +
		"access_token=" + accessToken +
		"&openid=" + openID

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userInfo := weixin_model.UserInfoRes{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	if userInfo.OpenID == "" || userInfo.ErrCode != 0 {
		return nil, fmt.Errorf("获取用户信息失败：%s", userInfo.ErrMsg)
	}

	return &userInfo, nil
}

// UserLogin 获取微信用户openId和sessionKey会话key 进行login  （小程序登录）
func (s *sUserAuth) UserLogin(ctx context.Context, info g.Map) (string, error) {
	from := gmap.NewStrAnyMapFrom(info)

	// 1.拿到code
	code := gconv.String(from.Get("code"))
	appId := gconv.String(from.Get("app_id"))

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil || merchantApp == nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "appId错误", "UserLogin")
	}

	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	if err != nil || merchantApp == nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "该appId没有对应的第三方应用", "UserLogin")
	}

	// 3.获取session_key和openId
	//openID := gconv.String(info["code"])
	res, err := getOpenIDAndSessionKey(code, merchantApp.AppId, merchantApp.ThirdAppId, thirdApp.AppAuthToken)
	if err != nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "", "WeiXin")
	}
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------OpenId： ------- "+res.Openid, "sUserAuth")
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------SessionKey： ------- "+res.SessionKey, "sUserAuth")

	return "success", nil
}

// 获取微信用户openId和sessionKey会话key
func getOpenIDAndSessionKey(code string, appID, thirdAppId string, componentAccessToken string) (*weixin_model.OpenIdAndSessionKeyRes, error) {
	//https: //api.weixin.qq.com/sns/component/jscode2session?appid=APPID&js_code=JSCODE&grant_type=authorization_code&component_appid=COMPONENT_APPID&component_access_token=COMPONENT_ACCESS_TOKEN

	url := "https://api.weixin.qq.com/sns/component/jscode2session?" +
		"appid=" + appID +
		"&js_code=" + code +
		"&grant_type=authorization_code" +
		"&component_appid=" + thirdAppId +
		"&component_access_token=" + componentAccessToken
	//
	//req := weixin_model.OpenIdAndSessionKeyReq{
	//	Appid:          appID,
	//	GrantType:      "authorization_code",
	//	ComponentAppid: thirdAppId,
	//	JsCode:         code,
	//}
	//
	//data, _ := gjson.Encode(req)
	//content := g.Client().GetContent(context.Background(), url, data)
	//
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := weixin_model.OpenIdAndSessionKeyRes{}
	//gjson.DecodeTo(content, &res)

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if res.Openid == "" || res.ErrCode != 0 {
		return nil, fmt.Errorf("获取用户openId和session_key失败：%s", res.ErrMsg)
	}

	return &res, nil
}
