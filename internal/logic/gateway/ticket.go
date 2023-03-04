package gateway

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
	"time"
)

type sTicket struct{}

func NewTicket() *sTicket {
	// 初始化文件内容

	result := &sTicket{}

	result.injectHook()
	return result
}

func (s *sTicket) injectHook() {
	weixin_service.Gateway().InstallHook(weixin_enum.Info.Type.Ticket, s.Ticket)
}

// Ticket 票据具体服务
func (s *sTicket) Ticket(ctx context.Context, info g.Map) bool {
	if info["MsgType"] != weixin_enum.Info.Type.Ticket.Code() {
		return false
	}

	data := weixin_model.EventMessageBody{}
	gconv.Struct(info["info"], &data)
	res := getComponentAccessToken(ctx, &data)

	return res
}

// getComponentAccessToken 获取第三方接口调用凭据 access_token
func getComponentAccessToken(ctx context.Context, data *weixin_model.EventMessageBody) bool {
	appConfig, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, data.AppId)
	appId := gconv.String(appConfig.AppId)
	// 3.获取token令牌 (第三方平台接口的调用凭据 access_token)
	tokenUrl := "https://api.weixin.qq.com/cgi-bin/component/api_component_token?component_appid=" + appId +
		"&component_appsecret=" + appConfig.AppSecret +
		"&component_verify_ticket=" + data.ComponentVerifyTicket
	fmt.Println(tokenUrl)
	componentAccessToken := g.Client().PostContent(ctx, tokenUrl)

	componentAccessTokenRes := weixin_model.ComponentAccessTokenRes{}
	gjson.DecodeTo(componentAccessToken, &componentAccessTokenRes)
	fmt.Println("令牌：", componentAccessTokenRes)

	// 缓存componentAccessToken 第三方接口调用凭据
	if componentAccessTokenRes.ComponentAccessToken != "" {
		gcache.Set(ctx, appId+"_component_access_token", componentAccessTokenRes.ComponentAccessToken, time.Duration(componentAccessTokenRes.ExpiresIn))
		return true
	}

	// 找出服务商
	thirdAppConfigInfo, err := weixin_service.ThirdAppConfig().GetThirdAppConfigByAppId(ctx, data.AppId)
	if err != nil {
		fmt.Println(" 服务商应用不存在")
		return false
	}
	thirdAppConfigInfo.ReExpiresIn.Format("")
	t := gtime.New(thirdAppConfigInfo.ReExpiresIn)
	timestamp := t.Timestamp()

	// 更新Token
	tokenData := weixin_model.UpdateAppAuthToken{
		AppId:        data.AppId,
		AppAuthToken: componentAccessTokenRes.ComponentAccessToken,
		ExpiresIn:    gtime.New(componentAccessTokenRes.ExpiresIn),
		ReExpiresIn:  gtime.New(gconv.Int64(componentAccessTokenRes.ExpiresIn) + timestamp), // 0 应该替换成原来时间的时间戳
	}

	// 修改数据库中服务商Token
	weixin_service.ThirdAppConfig().UpdateAppAuthToken(ctx, &tokenData)

	return false
}
