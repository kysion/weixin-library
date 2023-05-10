package merchant

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
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
}

func init() {
	weixin_service.RegisterUserAuth(NewUserAuth())
}

func (s *sUserAuth) injectHook() {
	//notifyHook := weixin_service.Gateway().GetServiceNotifyTypeHook()
	callHook := weixin_service.Gateway().GetCallbackMsgHook()

	callHook.InstallHook(weixin_enum.Info.CallbackType.UserAuth, s.UserAuthCallback)
}

func NewUserAuth() *sUserAuth {

	result := &sUserAuth{}

	result.injectHook()
	return result
}

// UserAuthCallback 处理授权回调请求
func (s *sUserAuth) UserAuthCallback(ctx context.Context, info g.Map) bool {
	from := gmap.NewStrAnyMapFrom(info)

	// 1.拿到code
	code := gconv.String(from.Get("code"))
	appId := gconv.String(from.Get("app_id"))

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil || merchantApp == nil {
		return false
	}

	// 2.获取access_token
	accessToken, err := getAccessToken(code, appId, merchantApp.AppSecret)
	if err != nil {
		sys_service.SysLogs().ErrorSimple(ctx, err, "获取AccessToken失败："+err.Error(), "WeiXin")
		return false
	}

	openID := gconv.String(info["code"])

	// 3.获取用户信息userInfo
	if openID == "" {
		userInfo, err := getUserInfo(accessToken, openID)
		if err != nil {
			sys_service.SysLogs().ErrorSimple(ctx, err, "获取用户信息失败："+err.Error(), "WeiXin")
			return false
		}
		// TODO: 处理用户信息
		fmt.Println("用户信息：", userInfo)
	} else {
		sys_service.SysLogs().ErrorSimple(ctx, err, "缺少OpenID参数："+err.Error(), "WeiXin")
		return false
	}

	return true
}

// 获取微信AccessToken
func getAccessToken(code string, appID, appSecret string) (string, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?" +
		"appid=" + appID +
		"&secret=" + appSecret +
		"&code=" + code +
		"&grant_type=authorization_code"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	accessToken := weixin_model.AccessTokenRes{}
	if err := json.Unmarshal(body, &accessToken); err != nil {
		return "", err
	}

	if accessToken.AccessToken == "" || accessToken.ErrCode != 0 {
		return "", fmt.Errorf("获取AccessToken失败：%s", accessToken.ErrMsg)
	}

	return accessToken.AccessToken, nil
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
