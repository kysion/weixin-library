package weixin_enum

import (
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/consumer"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/info_type"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/notify"
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
)

var (
	// Info 消息
	Info = info_type.Info

	// Notify 通知
	Notify = notify.Notify

	// Consumer 消费者
	Consumer = consumer.Consumer
)
