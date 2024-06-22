package merchant

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
	"strconv"
)

/*
	公众号用户授权相关
*/

// UserAuthCallback 处理网页授权回调请求 （公众号登录）
func (s *sUserAuth) UserAuthCallback(ctx context.Context, info g.Map) (int64, error) {
	from := gmap.NewStrAnyMapFrom(info)

	// 1.拿到code
	code := gconv.String(from.Get("code"))
	appId := gconv.String(from.Get("app_id"))
	sysUserId := gconv.Int64(from.Get("sys_user_id"))
	merchantId := gconv.Int64(from.Get("merchant_id"))

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil && merchantApp == nil {
		return sysUserId, err
	}
	// TODO 获取用户信息
	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, sysUserId)
	//if err != nil {
	//	return nil,err
	//}

	// 2.获取access_token  (能拿到openId和access_token)
	accessToken, err := getAccessTokenByH5(code, appId, merchantApp.AppSecret)
	if err != nil {
		sys_service.SysLogs().ErrorSimple(ctx, err, "获取AccessToken失败："+err.Error(), "WeiXin")
		return sysUserId, err
	}
	g.Dump(accessToken)

	sys_service.SysLogs().InfoSimple(ctx, nil, "\nOpenId："+accessToken.Openid, "sUserAuth")
	sys_service.SysLogs().InfoSimple(ctx, nil, "\nAccessToken： "+accessToken.AccessToken, "sUserAuth")

	openID := accessToken.Openid
	err = dao.WeixinConsumerConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if openID != "" {
			var weiXinConsumer *weixin_model.WeixinConsumerConfig
			if accessToken.Unionid != "" {
				//weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenId(ctx, openID, accessToken.Unionid)
				weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenIdAndAppId(ctx, openID, appId, accessToken.Unionid)
			} else {
				//weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenId(ctx, openID)
				weiXinConsumer, err = weixin_service.Consumer().GetConsumerByOpenIdAndAppId(ctx, openID, appId)
			}

			if weiXinConsumer == nil {
				id := sysUserId
				passwordLen := len(gconv.String(id))
				pwd := gstr.SubStr(gconv.String(id), passwordLen-6, 6) // 密码为id后六位

				data := sys_model.UserInnerRegister{
					Username:        strconv.FormatInt(gconv.Int64(id), 36),
					Password:        pwd,
					ConfirmPassword: pwd,
				}
				sysUser, err = sys_service.SysUser().CreateUser(ctx, data, sys_enum.User.State.Normal, sys_enum.User.Type.New(0, "匿名"), id)
				if err != nil {
					return err
				}
				sysUserId = sysUser.Id

			}
			if weiXinConsumer != nil {
				sysUserId = weiXinConsumer.SysUserId
			}
			// 3.获取用户信息userInfo
			userInfo, err := getUserInfoByH5(accessToken.AccessToken, openID)
			g.Dump(userInfo)

			if err != nil {
				return sys_service.SysLogs().ErrorSimple(ctx, err, "获取用户信息失败："+err.Error(), "WeiXin")
			}

			// 4.处理微信消费者数据
			if weiXinConsumer == nil { // 创建
				wConsumerData := kconv.Struct(userInfo, &weixin_model.WeixinConsumerConfig{})

				wConsumerData.OpenId = openID
				wConsumerData.SysUserId = sysUserId
				wConsumerData.UserType = sysUser.Type
				wConsumerData.UserState = sysUser.State
				wConsumerData.Avatar = userInfo.HeadImgURL

				wConsumerData.AccessToken = accessToken.AccessToken
				wConsumerData.RefreshToken = accessToken.RefreshToken
				wConsumerData.ExpiresIn = gtime.NewFromTimeStamp(gtime.Now().Timestamp() + accessToken.ExpiresIn)

				wConsumerData.UnionId = accessToken.Unionid
				wConsumerData.OpenId = accessToken.Openid
				wConsumerData.SessionKey = ""
				wConsumerData.AuthState = weixin_enum.Consumer.AuthState.Auth.Code()
				wConsumerData.AppType = weixin_enum.AppManage.AppType.PublicAccount.Code()
				wConsumerData.AppId = appId

				_, err = weixin_service.Consumer().CreateConsumer(ctx, wConsumerData)
				if err != nil {
					return err
				}

			} else { // 修改
				wConsumerData := kconv.Struct(userInfo, &weixin_model.UpdateConsumerReq{}) // 修改用户基本数据
				wConsumerData.Avatar = userInfo.HeadImgURL
				//_, err = weixin_service.Consumer().UpdateConsumer(ctx, weiXinConsumer.Id, wConsumerData)
				_, err = weixin_service.Consumer().UpdateConsumerByUserId(ctx, weiXinConsumer.SysUserId, wConsumerData)
				if err != nil {
					return err
				}

				_, err = weixin_service.Consumer().UpdateConsumerToken(ctx, openID, &weixin_model.UpdateConsumerTokenReq{ // 修改用户Token
					AccessToken:  &accessToken.AccessToken,
					SessionKey:   nil,
					RefreshToken: &accessToken.RefreshToken,
					ExpiresIn:    gtime.NewFromTimeStamp(gtime.Now().Timestamp() + accessToken.ExpiresIn),
				})

				if err != nil {
					return err
				}
			}

			// 5.存储consumer消费者数据
			g.Try(ctx, func(ctx context.Context) {
				s.ConsumerHook.Iterator(func(key hook.ConsumerKey, value hook.ConsumerHookFunc) {
					if key.ConsumerAction.Code() == weixin_enum.Consumer.ActionEnum.Auth.Code() && key.Category.Code() == weixin_enum.Consumer.Category.Consumer.Code() { // 如果订阅者是订阅授权,并且是操作kmk_consumer表
						g.Try(ctx, func(ctx context.Context) {
							data := hook.UserInfo{
								SysUserId:   sysUserId, // (消费者id = sys_User_id)
								UserInfoRes: *userInfo,
							}
							sys_service.SysLogs().InfoSimple(ctx, nil, "\n广播-------存储消费者数据 xxx-consumer", "sUserAuth")

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
								Id:             idgen.NextId(),
								FacilitatorId:  0,
								OperatorId:     0,
								SysUserId:      sysUserId,                                    // EmployeeId  == consumerId == sysUserId   三者相等
								MerchantId:     merchantId,                                   // 商家id，就是消费者首次扫码的商家
								PlatformType:   pay_enum.Order.TradeSourceType.Weixin.Code(), // 平台类型：1支付宝、2微信、4抖音、8银联
								ThirdAppId:     merchantApp.ThirdAppId,
								MerchantAppId:  merchantApp.AppId,
								PlatformUserId: openID,       // 平台账户唯一标识
								SysUserType:    sysUser.Type, // TODO 后期通过sysUserId拿到user，拿到type
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
		return sysUserId, err
	}

	return sysUserId, err
}

// 获取微信AccessToken （网页授权）
func getAccessTokenByH5(code string, appID, appSecret string) (*weixin_model.AccessTokenRes, error) {
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
func getUserInfoByH5(accessToken string, openID string) (*weixin_model.UserInfoRes, error) {
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
