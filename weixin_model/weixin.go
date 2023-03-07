package weixin_model

// authorizer_access_token  商家应用Token

// component_access_token 服务商应用Token  - 服务商

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
	AppId                        string `xml:"AppId" json:"app_id"   dc:"第三方平台 appid"`
	CreateTime                   int    `xml:"CreateTime" json:"create_time" dc:"时间戳" `
	InfoType                     string `xml:"InfoType" json:"info_type"  dc:"通知类型" `
	ComponentVerifyTicket        string `xml:"ComponentVerifyTicket" json:"component_verify_ticket" dc:"票据内容"`
	AuthorizerAppid              string `xml:"AuthorizerAppid" json:"authorizer_appid" dc:"公众号或小程序的appid"`
	AuthorizationCode            string `xml:"AuthorizationCode" json:"authorization_code" dc:"授权码，可用于获取授权信息"`
	AuthorizationCodeExpiredTime string `xml:"AuthorizationCodeExpiredTime" json:"authorization_code_expired_time" dc:"授权码过期时间 单位秒"`
	PreAuthCode                  string `xml:"PreAuthCode" json:"pre_auth_code" dc:"预授权码"`
}

/*
	AppId -> wx534d1a08aa84c529
	Encrypt -> OpuMbY5x5IAId+jfCQTYFCC7p3JarrbJCW6tzDTW8k0xwVfq/is1OIEWQB0oMvZ7gNg+0/W/zhzeEnAS8QkpywHLLHpcVu/QGkk7
*/

// ComponentAccessTokenReq 获取第三方平台接口的调用凭据Req
type ComponentAccessTokenReq struct {
	ComponentAppid        string `json:"component_appid" dc:"第三方平台 appid"`
	ComponentAppsecret    string `json:"component_appsecret" dc:"第三方平台 appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket" dc:"微信后台推送的 ticket"`
}

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

// ProAuthCodeReq 获取预授权码Req
type ProAuthCodeReq struct {
	ComponentAppid       string `json:"component_appid" dc:"第三方平台 appid"`
	ComponentAccessToken string `json:"component_access_token" dc:"第三方平台接口的调用凭据 component_access_token "`
}
