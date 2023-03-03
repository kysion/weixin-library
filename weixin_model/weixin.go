package weixin_model

import "encoding/xml"

type TicketRes struct {
	AppId                 string `json:"app_id" dc:"第三方平台 appid"`
	CreateTime            int64  `json:"create_time" dc:"时间戳，单位：s"`
	InfoType              string `json:"info_type" dc:"固定为：component_verify_ticket"`
	ComponentVerifyTicket string `json:"component_verify_ticket" dc:"Ticket 内容"`
}

// EventEncryptMsgReq  微信事件推送结构体
type EventEncryptMsgReq struct {
	AppId   string `json:"app_id" dc:"第三方平台 appid"`
	Encrypt string `json:"weixin_encrypt" dc:""`
}

// MessageEncryptReq  推送过来的微信消息加密结构体数据 例如：票据Ticket... (上面结构体解密后)
type MessageEncryptReq struct {
	Nonce        string `json:"Nonce" dc:"随机数" `
	Encrypt      string `json:"Encrypt" dc:"密文"`
	MsgSignature string `json:"MsgSignature" dc:"签名"`
	TimeStamp    string `json:"TimeStamp" dc:"时间戳"`
}

type MessageRRequest struct {
	EventEncryptMsgReq
	MessageEncryptReq
}

// EventMessageBody 事件推送
type EventMessageBody struct {
	XMLName                      xml.Name `xml:"xml"`
	AppId                        string   `xml:"AppId" json:"app_id"`
	CreateTime                   int      `xml:"CreateTime" json:"create_time"`
	InfoType                     string   `xml:"InfoType" json:"info_type"`
	ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket" json:"component_verify_ticket"`
	AuthorizerAppid              string   `xml:"AuthorizerAppid" json:"authorizer_appid"`
	AuthorizationCode            string   `xml:"AuthorizationCode" json:"authorization_code"`
	AuthorizationCodeExpiredTime string   `xml:"AuthorizationCodeExpiredTime" json:"authorization_code_expired_time"`
	PreAuthCode                  string   `xml:"PreAuthCode" json:"pre_auth_code"`
}

/*
	AppId -> wx534d1a08aa84c529
	Encrypt -> OpuMbY5x5IAId+jfCQTYFCC7p3JarrbJCW6tzDTW8k0xwVfq/is1OIEWQB0oMvZ7gNg+0/W/zhzeEnAS8QkpywHLLHpcVu/QGkk7
*/

// ComponentAccessTokenRes 获取第三方平台接口的调用凭据Res
type ComponentAccessTokenRes struct {
	ComponentAccessToken string `json:"component_access_token" dc:"第三方平台 access_token"`
	ExpiresIn            int    `json:"expires_in" dc:"有效期，单位：秒"`
}

// ProAuthCodeRes 获取预授权码Res
type ProAuthCodeRes struct {
	PreAuthCode string `json:"pre_auth_code" dc:"预授权码"`
	ExpiresIn   int    `json:"expires_in" dc:"有效期，单位：秒"`
}
