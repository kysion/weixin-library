package weixin_enum

import (
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/consumer"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/info_type"
)

type (
	// CallbackMsgType 回调消息
	CallbackMsgType info_type.CallBackMsgTypeEnum

	// ServiceNotifyType 应用通知
	ServiceNotifyType info_type.ServiceNotifyTypeEnum

	// ConsumerAction 消费者相关
	ConsumerAction consumer.ActionEnum
	Category       consumer.CategoryEnum
)

var (
	// Info 消息
	Info = info_type.Info

	// Consumer 消费者
	Consumer = consumer.Consumer
)
