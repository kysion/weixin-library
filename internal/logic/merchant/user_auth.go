package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/kysion/base-library/base_hook"
	hook "github.com/kysion/weixin-library/weixin_model/weixin_hook"
	"github.com/kysion/weixin-library/weixin_service"
)

// 用户授权 （静默、手动授权）

/*
网页授权：
   	1、构建授权连接，在回调中拿到code (前端)
	2、通过code拿到接口调用凭据access_token
	3、通过access_token拿到用户信息user_info
	4、通过refresh_token 进行刷新access_token
*/

/*
小程序授权流程：
	1.wx.login()拿到登陆凭据code （Ok）
	2.通过code拿到openId和session_key会话密钥  （Ok）
	3.后端实现自定义登陆态token
*/

type sUserAuth struct {
	// 消费者Hook
	ConsumerHook base_hook.BaseHook[hook.ConsumerKey, hook.ConsumerHookFunc]
}

//func init() {
//weixin_service.RegisterUserAuth(NewUserAuth())
//}

func NewUserAuth() weixin_service.IUserAuth {

	result := &sUserAuth{}

	//result.injectHook()
	return result
}

func (s *sUserAuth) injectHook() {
	//notifyHook := weixin_service.Gateway().GetServiceNotifyTypeHook()
	//callHook := weixin_service.Gateway().GetCallbackMsgHook()

	//callHook.InstallHook(weixin_enum.Info.CallbackType.UserAuth, s.UserAuthCallback)
	//
	//serviceHook := weixin_service.Gateway().GetServiceNotifyTypeHook()
	//
	//serviceHook.InstallHook(weixin_enum.Info.ServiceType.Event, s.UserEvent)

}

func (s *sUserAuth) InstallConsumerHook(infoType hook.ConsumerKey, hookFunc hook.ConsumerHookFunc) {
	sys_service.SysLogs().InfoSimple(context.Background(), nil, "\n-------订阅Alipay的sUserAuth用户授权Hook： ------- ", "sUserAuth")

	s.ConsumerHook.InstallHook(infoType, hookFunc)
}

func (s *sUserAuth) GetHook() base_hook.BaseHook[hook.ConsumerKey, hook.ConsumerHookFunc] {
	return s.ConsumerHook
}
