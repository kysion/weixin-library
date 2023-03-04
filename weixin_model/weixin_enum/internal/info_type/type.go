package info_type

import "github.com/kysion/base-library/utility/enum"

type InfoTypeEnum enum.IEnumCode[string]

// 各种消息类型
type infoType struct {
	Ticket               InfoTypeEnum
	ComponentAccessToken InfoTypeEnum
	WeixinAppAuth        InfoTypeEnum
	WeixinWallet         InfoTypeEnum
}

var InfoType = infoType{
	Ticket:               enum.New[InfoTypeEnum]("component_verify_ticket", "票据"),
	ComponentAccessToken: enum.New[InfoTypeEnum]("ComponentAccessToken", "第三方平台接口的调用凭据"),
	//WeixinAppAuth:        enum.New[InfoTypeEnum]("alipay_app_auth", "应用认证授权"),
	//WeixinWallet:         enum.New[InfoTypeEnum]("alipay_wallet", "用户登录"),
}
