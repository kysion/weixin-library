package weixin_enum

import (
	"github.com/kysion/weixin-library/weixin_model/weixin_enum/internal/info_type"
)

type (
	// CallbackMsgType 回调消息
	CallbackMsgType info_type.CallBackMsgTypeEnum

	// ServiceNotifyType 应用通知
	ServiceNotifyType info_type.ServiceNotifyTypeEnum
)

var (
	// Info 消息
	Info = info_type.Info
)
