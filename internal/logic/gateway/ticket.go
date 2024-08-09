package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
	"net/http"
)

type sTicket struct{}

func NewTicket() weixin_service.ITicket {
	// 初始化文件内容

	result := &sTicket{}

	result.injectHook()
	return result
}

func (s *sTicket) injectHook() {
	serviceHook := weixin_service.Gateway().GetServiceNotifyTypeHook()

	serviceHook.InstallHook(weixin_enum.Info.ServiceType.Ticket, s.Ticket)
}

// Ticket 票据具体服务
func (s *sTicket) Ticket(ctx context.Context, info g.Map) bool {
	if info["MsgType"] != weixin_enum.Info.ServiceType.Ticket.Code() {
		return false
	}
	//	Services 解密后的内容： &{wxb64aa49959fa359c 1720167532 component_verify_ticket ticket@@@LsDX8DrrnTcYrr0nk7VndG9ickwwcupzE1MPhGAp4zPSgCrdEkXupN2mVXxsZhguYaQNGQ7Yb5S0Ghls-YzhXw    }

	data := weixin_model.EventMessageBody{}
	_ = gconv.Struct(info["info"], &data)
	res := getComponentAccessToken(ctx, &data)

	return res
}

// getComponentAccessToken 获取第三方接口调用凭据 component_access_token
func getComponentAccessToken(ctx context.Context, data *weixin_model.EventMessageBody) bool {
	appConfig, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, data.AppId)
	if appConfig == nil || err != nil {
		return false
	}

	appId := gconv.String(appConfig.AppId)
	// 3.获取token令牌 (第三方平台接口的调用凭据 access_token)
	tokenUrl := "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	fmt.Println(tokenUrl)

	//66_5JMMvvOh0aT5KP7okeJ-ipV40If7M4ymowIOrILxxbdv5firRTiONfi6DouWvf6Nz0InYT7kI0tg9OLaZg2ZCwks0nJ1JCK6y7VjBu36_tMPlnwbumO-HAjqOe8LHVfAEAFBP
	tokenReq := weixin_model.ComponentAccessTokenReq{
		ComponentAppid:        appId,
		ComponentAppsecret:    appConfig.AppSecret,
		ComponentVerifyTicket: data.ComponentVerifyTicket,
	}
	tokenReqJson, _ := json.Marshal(tokenReq) // post请求参数不能直接拼接在URL，应该使用json序列化数据

	componentAccessToken := g.Client().PostContent(ctx, tokenUrl, tokenReqJson)
	fmt.Println("获取令牌返回数据：", componentAccessToken)

	componentAccessTokenRes := weixin_model.ComponentAccessTokenRes{}
	_ = gjson.DecodeTo(componentAccessToken, &componentAccessTokenRes)
	fmt.Println("令牌：", componentAccessTokenRes)

	// 找出服务商
	if componentAccessTokenRes.ComponentAccessToken != "" {
		thirdAppConfigInfo, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, data.AppId)
		if err != nil {
			fmt.Println(" 服务商应用不存在")
			return false
		}
		thirdAppConfigInfo.ReExpiresIn.Format("")
		//t := gtime.New(thirdAppConfigInfo.ReExpiresIn)
		//timestamp := t.Timestamp()
		timestamp := gtime.Now().Timestamp()

		// 更新Token
		tokenData := weixin_model.UpdateAppAuthToken{
			AppId:        data.AppId,
			AppAuthToken: componentAccessTokenRes.ComponentAccessToken,
			//ExpiresIn:    gtime.New(componentAccessTokenRes.ExpiresIn),
			ExpiresIn:   gtime.New(gconv.Int64(componentAccessTokenRes.ExpiresIn) + timestamp),
			ReExpiresIn: gtime.New(gconv.Int64(componentAccessTokenRes.ExpiresIn) + timestamp), // 0 应该替换成原来时间的时间戳
		}

		// 修改数据库中服务商Token
		weixin_service.ThirdAppConfig().UpdateAppAuthToken(ctx, &tokenData)
	}

	return true
}

// GetTicket 获取票据
func (s *sTicket) GetTicket(ctx context.Context, appId string) (string, error) {
	appConfig, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if appConfig == nil || err != nil {
		return "", err
	}

	var url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", appConfig.AppAuthToken)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var result weixin_model.TicketResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Ticket, nil
}
