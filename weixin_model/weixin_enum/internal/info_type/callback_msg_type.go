package info_type

import "github.com/kysion/base-library/utility/enum"

type CallBackMsgTypeEnum enum.IEnumCode[string]

// 各种回调消息类型  - 某人某行为产生
type callBackMsgType struct {
	ComponentAccessToken CallBackMsgTypeEnum

	// 应用授权通知类型
	Authorized       CallBackMsgTypeEnum
	UpdateAuthorized CallBackMsgTypeEnum
	Unauthorized     CallBackMsgTypeEnum

	UserAuth CallBackMsgTypeEnum
}

var CallBackMsgType = callBackMsgType{

	ComponentAccessToken: enum.New[CallBackMsgTypeEnum]("ComponentAccessToken", "第三方平台接口的调用凭据"),

	// 应用授权
	Authorized:       enum.New[CallBackMsgTypeEnum]("authorized", "授权成功"),
	UpdateAuthorized: enum.New[CallBackMsgTypeEnum]("updateauthorized", "更新授权"),
	Unauthorized:     enum.New[CallBackMsgTypeEnum]("unauthorized", "取消授权"),

	// 用户授权   （！！！！！需要替换的，根据微信用户授权后实际返回的type进行修改）
	UserAuth: enum.New[CallBackMsgTypeEnum]("userAuth", "用户授权"),
}
