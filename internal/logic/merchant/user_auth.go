package merchant

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
	"log"
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

// TODO ----------------------------------------------------公众号---------------------------------------------------------------------------------------------

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
	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, sysUserId)
	if err != nil {
		return false
	}

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
				wConsumerData.SysUserId = sysUserId     // TODO 后期想办法将sysUserId传递
				wConsumerData.UserType = sysUser.Type   // TODO 后期通过sysUserId拿到user，拿到type
				wConsumerData.UserState = sysUser.State // TODO 用户状态,拿到user，拿到type
				wConsumerData.Avatar = userInfo.HeadImgURL

				wConsumerData.AccessToken = accessToken.AccessToken
				wConsumerData.RefreshToken = accessToken.RefreshToken
				wConsumerData.ExpiresIn = gtime.New(accessToken.ExpiresIn)

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
								UserId:        openID,       // 平台账户唯一标识
								Type:          sysUser.Type, // TODO 后期通过sysUserId拿到user，拿到type
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

// TODO -----------------------------------------------小程序--------------------------------------------------------------------------------------------------

// GetMiniAppUserInfo 获取小程序用户唯一标识，用于检查是否注册,如果已经注册，返会openId
func (s *sUserAuth) GetMiniAppUserInfo(ctx context.Context, authCode string, appId string, getDetail bool) (*weixin_model.UserInfoRes, error) {
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------获取用户openID---- ", "sUserAuth")

	//openId := ""
	code := authCode
	userInfo := &weixin_model.UserInfoRes{}

	// 1.根据AppId获取商家相关配置，包括AppAuthToken
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil || merchantApp == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "appId错误", "sUserAuth")
	}

	// 2.获取商家应用对应的第三方应用
	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	if err != nil || merchantApp == nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该appId没有对应的第三方应用", "sUserAuth")
	}

	// 3.获取session_key和openId，非静默授权的情况下能拿到unionId
	res, err := getOpenIDAndSessionKey(code, merchantApp.AppId, merchantApp.ThirdAppId, thirdApp.AppAuthToken)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "获取用户唯一标识openId失败，请检查", "sUserAuth")
	}

	userInfo.OpenID = res.Openid

	//userInfo.AccessToken = res.AccessToken
	//userInfo.RefreshToken = res.RefreshToken
	//userInfo.ExpiresIn = gtime.New(res.ExpiresIn)

	//userInfo.RefreshToken = res.RefreshToken
	userInfo.SessionKey = res.SessionKey

	if res.Unionid != "" { // openId和unionId都拿到了
		userInfo.UnionID = res.Unionid
	}

	// 静默授权的情况下拿unionId，需要通过获取userInfo，从而拿到unionId
	if !getDetail {
		return userInfo, nil
	}

	// TODO 获取用户信息失败。token有误
	userInfo, err = getSNSUserInfo(ctx, userInfo.OpenID, userInfo.AccessToken) // TODO token为空

	if err != nil {
		return userInfo, sys_service.SysLogs().ErrorSimple(ctx, err, "获取用户unionId失败，请检查", "sUserAuth")
	}

	if userInfo.OpenID == "" {
		userInfo.OpenID = res.Openid
	}

	//userInfo.AccessToken = res.AccessToken
	//userInfo.RefreshToken = res.RefreshToken
	//userInfo.ExpiresIn = gtime.New(res.ExpiresIn)

	if userInfo.SessionKey == "" {
		userInfo.SessionKey = res.SessionKey
	}

	if userInfo.UnionID == "" { // openId和unionId都拿到了
		userInfo.UnionID = res.Unionid
	}

	return userInfo, nil
}

// UserLogin 获取微信用户openId和sessionKey会话key 进行login  （小程序登录）
func (s *sUserAuth) UserLogin(ctx context.Context, info g.Map) (string, error) {
	sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------小程序用户登陆---- ", "sUserAuth")

	from := gmap.NewStrAnyMapFrom(info)

	// 1.拿到code
	code := gconv.String(from.Get("code"))
	appId := gconv.String(from.Get("app_id"))
	sysUserId := gconv.Int64(from.Get("sys_user_id"))
	merchantId := gconv.Int64(from.Get("merchant_id"))
	userInfo := weixin_model.UserInfoRes{}
	gconv.Struct(from.Get("user_info"), &userInfo)

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil || merchantApp == nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "appId错误", "UserLogin")
	}

	thirdApp, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, merchantApp.ThirdAppId)
	if err != nil || merchantApp == nil {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "该appId没有对应的第三方应用", "UserLogin")
	}

	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, sysUserId)
	if err != nil {
		return "", err
	}

	// 3.获取session_key和openId
	//openID := gconv.String(info["code"])

	if userInfo.OpenID == "" { // 说明没有传递
		res, err := getOpenIDAndSessionKey(code, merchantApp.AppId, merchantApp.ThirdAppId, thirdApp.AppAuthToken)
		if err != nil {
			return "", sys_service.SysLogs().ErrorSimple(ctx, err, "获取用户唯一标识openId失败，请检查", "WeiXin")
		}

		sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------OpenId： ------- "+res.Openid, "sUserAuth")
		sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------SessionKey： ------- "+res.SessionKey, "sUserAuth")

		gconv.Struct(res, &userInfo)
	}

	// 3.获取用户信息userInfo

	// 4.将用户信息存储
	openID := userInfo.OpenID
	var weiXinConsumer *weixin_model.WeixinConsumerConfig

	if openID != "" {
		if userInfo.UnionID != "" {
			weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenId(ctx, openID, userInfo.UnionID)
		} else {
			weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenId(ctx, openID)
		}

	} else {
		return "", sys_service.SysLogs().ErrorSimple(ctx, err, "缺少OpenID参数："+err.Error(), "WeiXin")
	}

	err = dao.WeixinConsumerConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userInfoRes := weixin_model.UserInfoRes{}
		gconv.Struct(userInfo, &userInfoRes)

		if weiXinConsumer == nil { // 创建
			wConsumerData := weixin_model.WeixinConsumerConfig{}
			gconv.Struct(userInfo, &wConsumerData)

			wConsumerData.OpenId = openID
			wConsumerData.SysUserId = sysUserId     // TODO 后期想办法将sysUserId传递
			wConsumerData.UserType = sysUser.Type   // TODO 后期通过sysUserId拿到user，拿到type
			wConsumerData.UserState = sysUser.State // TODO 用户状态,拿到user，拿到type

			wConsumerData.UnionId = userInfo.UnionID
			wConsumerData.OpenId = userInfo.OpenID
			wConsumerData.SessionKey = userInfo.SessionKey

			_, err = weixin_service.Consumer().CreateConsumer(ctx, &wConsumerData)
			if err != nil {
				return err
			}

		} else { // 修改
			// TODO,因为小程序wx.login是没有userInfo的，所以我们将session_key存储好，供后续接口使用，例如wx.getUserInfo接口
			// wConsumerData := weixin_model.UpdateConsumerReq{}
			//_, err = weixin_service.Consumer().UpdateConsumer(ctx, weiXinConsumer.Id, &wConsumerData)
			//if err != nil {
			//	return err
			//}

			// 修改sessionKey
			_, err = weixin_service.Consumer().UpdateConsumerToken(ctx, openID, &weixin_model.UpdateConsumerTokenReq{
				//AccessToken:  &userInfo.AccessToken,
				//RefreshToken: &userInfo.RefreshToken,
				//ExpiresIn:    userInfo.ExpiresIn,
				SessionKey: &userInfo.SessionKey,
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
							UserInfoRes: userInfoRes,
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
							UserId:        openID,       // 平台账户唯一标识
							Type:          sysUser.Type, // TODO 后期通过sysUserId拿到user，拿到type
						}

						sys_service.SysLogs().InfoSimple(ctx, nil, "\n广播-------存储第三方应用和用户关系记录 kmk-plat_form_user", "sMerchantService")

						value(ctx, platformUser) // 调用Hook
					})
				}
			})
		})

		return nil
	})

	// 5.自定义实现微信登陆态 （JwtToken）

	// getTinyAppUserInfo() 获取用户信息

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

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if res.Openid == "" || res.ErrCode != 0 {
		return nil, fmt.Errorf("获取用户openId和session_key失败：%s", res.ErrMsg)
	}

	return &res, nil
}

// TODO ----------------------------------------小程序用户信息，还没测试---------------------------------------------------------------------------------------------------------
type wxUserInfo struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	AvatarURL string `json:"avatarUrl"`
}

// 小程序获取用户信息
func miniAppGetUserInfo(sessionKey, encryptedData, iv string, openId string) (*wxUserInfo, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/getuserinfo?access_token=%s&openid=%s&lang=zh_CN", sessionKey, openId)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, err
	}
	q := req.URL.Query()
	q.Add("encryptedData", encryptedData)
	q.Add("iv", iv)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make HTTP request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo wxUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return nil, err
	}

	return &userInfo, nil
}

// GetTinyAppUserInfo 小程序获取用户数据
func (s *sUserAuth) GetTinyAppUserInfo(ctx context.Context, sessionKey, encryptedData, iv, appId string, openId string) (*weixin_model.UserInfoRes, error) {
	// 假设这里已经获取到了用户的 sessionKey 和 encryptedData
	//sessionKey := "<your_session_key>"
	//encryptedData := "<your_encrypted_data>"
	//iv := "<your_iv>"

	userInfo, err := miniAppGetUserInfo(sessionKey, encryptedData, iv, openId)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		return nil, err
	}

	log.Printf("User info: %+v", userInfo)

	res := kconv.Struct(userInfo, &weixin_model.UserInfoRes{})

	return res, nil
}

// 开放平台获取userInfo
func getSNSUserInfo(ctx context.Context, openid string, accessToken string) (*weixin_model.UserInfoRes, error) {

	// 获取access_token和openid
	//access_token := "YOUR ACCESS_TOKEN"
	//openid := "YOUR OPENID"

	// 根据access_token和openid获取用户信息  // https://api.weixin.qq.com/sns/userinfo?access_token=fu6joqgW/fyS3qp/9WRisg==&openid=oc9L-5Kig8S10IaBYOiOcyUp4EbQ&lang=zh_CN
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", accessToken, openid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	// 解析结果
	var user_info weixin_model.UserInfoRes
	if err := json.NewDecoder(resp.Body).Decode(&user_info); err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 打印用户信息
	fmt.Printf("openid=%s, nickname=%s\n", user_info.OpenID, user_info.NickName)

	return &user_info, nil
}
