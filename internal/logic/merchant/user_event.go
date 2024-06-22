package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_dao"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	hook "github.com/kysion/weixin-library/weixin_model/weixin_hook"
	"github.com/kysion/weixin-library/weixin_service"
)

/* 用户关注公众号的消息通知：
"MsgType": "event",
	{
        ToUserName:   "gh_3fa93e09aa18",
        FromUserName: "ot4C45uA1iium5jntPGyTnlkY_EA",
        CreateTime:   "1705480023",
        MsgType:      "event",
        Event:        "subscribe",
	}

*/

/* 用户取消关注公众号的消息通知：
"MsgType": "event",
	{
		ToUserName:   "gh_3fa93e09aa18",
       FromUserName: "ot4C45uA1iium5jntPGyTnlkY_EA",
       CreateTime:   "1705479865",
       MsgType:      "event",
       Event:        "unsubscribe",
     }
*/

/* 用户通知删除授权的昵称和头像的消息通知：
"MsgType": "event",
	{
        ToUserName:   "gh_3fa93e09aa18",
        FromUserName: "ot4C45mKCz4m90ZetlLKoM4jZvvU",
        CreateTime:   "1705480106",
        MsgType:      "event",
        Event:        "user_authorization_revoke",
		RevokeInfo:   "205",
	}
*/
/* 用户给公众号发送消息的消息通知：
"MsgType": "text",
{
		ToUserName: "gh_3fa93e09aa18",
		FromUserName : "ot4C45uA1iium5jntPGyTnlkY_EA",
		CreateTime : "1705480266",
		MsgType : "text",
		Event : "",
		Url : "",
		PicUrl : "",
		MediaId : "",
		ThumbMediaId : "",
		Content : "刚刚",
		MsgId :24416605642226393,
}
*/

/*
	拓展：还有其他的Event类型
		user_info_modified：用户资料变更，
		user_authorization_revoke：用户撤回，
		user_authorization_cancellation：用户完成注销；

	拓展：RevokeInfo代表：
		用户撤回的H5授权信息：201:地址,202:发票信息,203:卡券信息,204:麦克风,205:昵称和头像,206:位置信息,207:选中的图片或视频
*/

type sUserEvent struct {
	// 消费者Hook
	ConsumerHook base_hook.BaseHook[hook.ConsumerKey, hook.ConsumerHookFunc]
}

func init() {
	//weixin_service.RegisterUserEvent(NewUserEvent())
}

func NewUserEvent() *sUserEvent {

	result := &sUserEvent{}

	result.injectHook()
	return result
}

func (s *sUserEvent) injectHook() {
	//notifyHook := weixin_service.Gateway().GetServiceNotifyTypeHook()
	//callHook := weixin_service.Gateway().GetCallbackMsgHook()

	//callHook.InstallHook(weixin_enum.Info.CallbackType.UserEvent, s.UserEventCallback)

	serviceHook := weixin_service.Gateway().GetServiceNotifyTypeHook()

	serviceHook.InstallHook(weixin_enum.Info.ServiceType.Event, s.UserEvent)

}

// UserEvent 用户相关事件
func (s *sUserEvent) UserEvent(ctx context.Context, info g.Map) bool {
	g.Dump("收到的用户消息事件：", info)
	from := gmap.NewStrAnyMapFrom(info)
	infoValue := from.Get("info")
	appId := gconv.String(from.Get("appId"))
	messageInfo := kconv.Struct(infoValue, &weixin_model.MessageBodyDecrypt{})

	//  注册事件处理
	if messageInfo.Event != "" {
		s.Subscribe(ctx, appId, messageInfo) // subscribe(订阅)、

		s.UnSubscribe(ctx, appId, messageInfo) // unsubscribe(取消订阅)

		s.UserAuthorizationRevoke(ctx, appId, messageInfo) // user_authorization_revoke（用户撤回）
	}

	return true
}

// Subscribe 用户关注公众号
func (s *sUserEvent) Subscribe(ctx context.Context, appId string, info *weixin_model.MessageBodyDecrypt) (bool, error) {
	if info.Event != "subscribe" {
		return true, nil
	}
	/*
			{
		        ToUserName:   "gh_3fa93e09aa18",
		        FromUserName: "ot4C45uA1iium5jntPGyTnlkY_EA",
		        CreateTime:   "1705480023",
		        MsgType:      "event",
		        Event:        "subscribe",
			}
	*/

	// 不是公众号退出
	appConfig, _ := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if appConfig != nil && appConfig.AppType != weixin_enum.AppManage.AppType.PublicAccount.Code() {
		return true, nil
	}

	// 设置用户是否关注公众号
	_, err := weixin_service.Consumer().SetIsFollowPublic(ctx, info.FromUserName, appId, weixin_enum.Consumer.IsFollowPublic.Subscribe.Code())
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "修改用户是够关注公众号失败", weixin_dao.WeixinConsumerConfig.Table())
	}

	return true, nil
}

// UnSubscribe 用户取消关注公众号
func (s *sUserEvent) UnSubscribe(ctx context.Context, appId string, info *weixin_model.MessageBodyDecrypt) (bool, error) {
	/*
			{
				ToUserName:   "gh_3fa93e09aa18",
		       FromUserName: "ot4C45uA1iium5jntPGyTnlkY_EA",
		       CreateTime:   "1705479865",
		       MsgType:      "event",
		       Event:        "unsubscribe",
		     }
	*/
	if info.Event != "unsubscribe" {
		return true, nil
	}
	// 不是公众号退出
	appConfig, _ := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if appConfig != nil && appConfig.AppType != weixin_enum.AppManage.AppType.PublicAccount.Code() {
		return true, nil
	}

	// 设置用户是否关注公众号
	_, err := weixin_service.Consumer().SetIsFollowPublic(ctx, info.FromUserName, appId, weixin_enum.Consumer.IsFollowPublic.UnSubscribe.Code())
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "修改用户是够关注公众号失败", weixin_dao.WeixinConsumerConfig.Table())
	}

	return true, nil
}

// UserAuthorizationRevoke 用户撤回事件  -- 取消授权
func (s *sUserEvent) UserAuthorizationRevoke(ctx context.Context, appId string, info *weixin_model.MessageBodyDecrypt) (bool, error) {
	if info.Event != "user_authorization_revoke" {
		return true, nil
	}
	/*
			{
		        ToUserName:   "gh_3fa93e09aa18",
		        FromUserName: "ot4C45mKCz4m90ZetlLKoM4jZvvU",
		        CreateTime:   "1705480106",
		        MsgType:      "event",
		        Event:        "user_authorization_revoke",
				RevokeInfo:   "205",
			}
	*/

	configConfig, _ := weixin_service.Consumer().GetConsumerByOpenIdAndAppId(ctx, info.OpenID, appId)

	if configConfig != nil && configConfig.Id != 0 {
		if info.RevokeInfo == "205" { // 205撤回头像和昵称
			_, err := weixin_service.Consumer().UpdateConsumerAuthState(ctx, configConfig.Id, weixin_enum.Consumer.AuthState.UnAuth.Code())
			if err != nil {
				return false, sys_service.SysLogs().ErrorSimple(ctx, err, "修改用户授权状态失败", weixin_dao.WeixinConsumerConfig.Table())
			}
		}
	}

	return true, nil
}

// TODO user_info_modified：用户资料变更，

// TODO user_authorization_revoke：用户撤回，

// TODO user_authorization_cancellation：用户完成注销；
