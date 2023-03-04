package merchant

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
)

type sAppAuth struct {
}

// gateway 主要用于记录和服务商相关操作

// merchant 主要记录和商家有关，例如一些商家消息的hook注册，

// internal 主要用于拓展SDK所不具备。票据例外

func init() {
	weixin_service.RegisterAppAuth(NewAppAuth())
}

func NewAppAuth() *sAppAuth {

	result := &sAppAuth{}

	result.injectHook()
	return result
}

func (s *sAppAuth) injectHook() {
	weixin_service.Gateway().InstallHook(weixin_enum.Info.Type.ComponentAccessToken, s.AppAuth)
}

// AppAuth 应用授权具体服务
func (s *sAppAuth) AppAuth(ctx context.Context, info g.Map) bool {
	//getComponentAccessToken(ctx, gconv.String(info))
	return true
}
