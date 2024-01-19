package weixin_enum

import (
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/app_manager"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/consumer"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/info_type"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/notify"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/weixin_pay"
)

type (
	// CallbackMsgType 回调消息
	CallbackMsgType info_type.CallBackMsgTypeEnum

	// ServiceNotifyType 应用通知
	ServiceNotifyType info_type.ServiceNotifyTypeEnum

	// NotifyType 异步通知类型
	NotifyType notify.NotifyTypeEnum

	// ConsumerAction 消费者相关
	ConsumerAction consumer.ActionEnum
	Category       consumer.CategoryEnum

	// MerchantType 微信支付
	MerchantType weixin_pay.MerchantTypeEnum

	// AppType 应用类型
	AppType app_manager.AppTypeEnum
)

var (
	// Info 消息
	Info = info_type.Info

	// Notify 通知
	Notify = notify.Notify

	// Consumer 消费者
	Consumer = consumer.Consumer

	// Pay 微信支付
	Pay = weixin_pay.Pay

	// AppManage 应用管理
	AppManage = app_manager.AppManager
)
