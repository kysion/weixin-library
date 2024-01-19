package info_type

import "github.com/kysion/base-library/utility/enum"

type ServiceNotifyTypeEnum enum.IEnumCode[string]

// 各种应用通知类型  - 平台主动发送
type serviceNotifyType struct {
	Ticket ServiceNotifyTypeEnum
	Event  ServiceNotifyTypeEnum
}

var ServiceNotifyType = serviceNotifyType{
	Ticket: enum.New[ServiceNotifyTypeEnum]("component_verify_ticket", "票据"),

	// 分账通知

	// 用户-公众号相关的消息推送 （有关注/取消关注事件、扫描带参数二维码事件、上报地理位置事件、自定义菜单事件、点击菜单拉取消息时的事件推送、点击菜单跳转链接时的事件推送）
	Event: enum.New[ServiceNotifyTypeEnum]("event", "消息类型为：事件"),
}
