package info_type

import "github.com/kysion/base-library/utility/enum"

type ServiceNotifyTypeEnum enum.IEnumCode[string]

// 各种应用通知类型  - 平台主动发送
type serviceNotifyType struct {
	Ticket ServiceNotifyTypeEnum
}

var ServiceNotifyType = serviceNotifyType{
	Ticket: enum.New[ServiceNotifyTypeEnum]("component_verify_ticket", "票据"),

	// 分账通知

}
