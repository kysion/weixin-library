package info_type

import "github.com/kysion/base-library/utility/enum"

type CallBackMsgTypeEnum enum.IEnumCode[string]

// 各种回调消息类型  - 某人某行为产生
type callBackMsgType struct {
	ComponentAccessToken CallBackMsgTypeEnum

	// 用户授权
	UserAuth CallBackMsgTypeEnum

	// 用户发消息
	UserSendMessage CallBackMsgTypeEnum
}

var CallBackMsgType = callBackMsgType{

	ComponentAccessToken: enum.New[CallBackMsgTypeEnum]("ComponentAccessToken", "第三方平台接口的调用凭据"),

	// 用户授权   （！！！！！需要替换的，根据微信用户授权后实际返回的type进行修改）
	UserAuth: enum.New[CallBackMsgTypeEnum]("userAuth", "用户授权"),

	// 用户发消息
	UserSendMessage: enum.New[CallBackMsgTypeEnum]("text", "用户发送消息"),
}
