package info_type

import "github.com/kysion/base-library/utility/enum"

type ServiceNotifyTypeEnum enum.IEnumCode[string]

// 各种应用通知类型  - 平台主动发送
type serviceNotifyType struct {
	// 第三方平台票据
	Ticket ServiceNotifyTypeEnum
	Event  ServiceNotifyTypeEnum

	// 应用授权通知类型 TODO 需要知道到底是通过servcei 发送的还是 callback
	Authorized       CallBackMsgTypeEnum
	UpdateAuthorized CallBackMsgTypeEnum
	Unauthorized     CallBackMsgTypeEnum
}

var ServiceNotifyType = serviceNotifyType{
	// 微信推送给第三方平台的票据
	Ticket: enum.New[ServiceNotifyTypeEnum]("component_verify_ticket", "票据"),

	// 分账通知

	// 用户-公众号相关的消息推送 （有关注/取消关注事件、扫描带参数二维码事件、上报地理位置事件、自定义菜单事件、点击菜单拉取消息时的事件推送、点击菜单跳转链接时的事件推送）
	Event: enum.New[ServiceNotifyTypeEnum]("event", "消息类型为：事件"),

	// 应用授权 ---- servcei 接收
	Authorized:       enum.New[CallBackMsgTypeEnum]("authorized", "授权成功"),
	UpdateAuthorized: enum.New[CallBackMsgTypeEnum]("updateauthorized", "更新授权"),
	Unauthorized:     enum.New[CallBackMsgTypeEnum]("unauthorized", "取消授权"),
}
