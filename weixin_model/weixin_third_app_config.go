package weixin_model

import "github.com/gogf/gf/v2/os/gtime"

type WeixinThirdAppConfig struct {
	Id             int64       `json:"id"             dc:"授权商家id"`
	Name           string      `json:"name"           dc:"服务商名称"`
	AppId          string      `json:"appId"          dc:"服务商应用Id"`
	AppName        string      `json:"appName"        dc:"服务商应用名称"`
	AppType        int         `json:"appType"        dc:"服务商应用类型"`
	AppAuthToken   string      `json:"appAuthToken"   dc:"服务商应用授权token"`
	ExpiresIn      *gtime.Time `json:"expiresIn"      dc:"Token过期时间"`
	ReExpiresIn    *gtime.Time `json:"reExpiresIn"    dc:"Token限期刷新时间"`
	UnionMainId    int64       `json:"unionMainId"    dc:"关联主体id"`
	SysUserId      int64       `json:"sysUserId"      dc:"用户id"`
	ExtJson        string      `json:"extJson"        dc:"拓展字段"`
	AppGatewayUrl  string      `json:"appGatewayUrl"  dc:"网关地址"`
	AppCallbackUrl string      `json:"appCallbackUrl" dc:"回调地址"`
	AppSecret      string      `json:"appSecret"      dc:"服务商应用密钥"`
	MsgVerfiyToken string      `json:"msgVerfiyToken" dc:"消息校验Token"`
	MsgEncryptKey  string      `json:"msgEncryptKey"  dc:"消息加密解密密钥"`
	AuthInitUrl    string      `json:"authInitUrl"    dc:"授权发起页域名"`
	ServerDomain   string      `json:"serverDomain"   dc:"服务器域名"`
	BusinessDomain string      `json:"businessDomain" dc:"业务域名"`
	AuthTestAppIds string      `json:"authTestAppIds" dc:"授权测试应用列表"`
	PlatformSite   string      `json:"platformSite"   dc:"平台官方"`
	Logo           string      `json:"logo"           dc:"服务商logo"`
	State          int         `json:"state"          dc:"状态：0禁用 1启用"`
	ReleaseState   int         `json:"releaseState"   dc:"发布状态：0未发布 1已发布"`
	HttpsCert      string      `json:"httpsCert"      dc:"域名证书"`
	HttpsKey       string      `json:"httpsKey"       dc:"域名私钥"`
}

type UpdateThirdAppConfig struct {
	Id             int64       `json:"id"             dc:"授权商家id"`
	Name           string      `json:"name"           dc:"服务商名称" v:"required#服务商"`
	AppAuthToken   string      `json:"appAuthToken"   dc:"服务商应用授权token"`
	ExpiresIn      *gtime.Time `json:"expiresIn"      dc:"Token过期时间"`
	ReExpiresIn    *gtime.Time `json:"reExpiresIn"    dc:"Token限期刷新时间"`
	ExtJson        string      `json:"extJson"        dc:"拓展字段"`
	AppGatewayUrl  string      `json:"appGatewayUrl"  dc:"网关地址"`
	AppCallbackUrl string      `json:"appCallbackUrl" dc:"回调地址"`
	AppSecret      string      `json:"appSecret"      dc:"服务商应用密钥"`
	MsgVerfiyToken string      `json:"msgVerfiyToken" dc:"消息校验Token"`
	MsgEncryptKey  string      `json:"msgEncryptKey"  dc:"消息加密解密密钥"`
	AuthInitUrl    string      `json:"authInitUrl"    dc:"授权发起页域名"`
	ServerDomain   string      `json:"serverDomain"   dc:"服务器域名"`
	BusinessDomain string      `json:"businessDomain" dc:"业务域名"`
	AuthTestAppIds string      `json:"authTestAppIds" dc:"授权测试应用列表"`
	PlatformSite   string      `json:"platformSite"   dc:"平台官方"`
	Logo           string      `json:"logo"           description:"服务商logo"`
}
type UpdateAppAuthToken struct {
	AppId        string      `json:"appId"             dc:"服务商应用id"`
	AppAuthToken string      `json:"appAuthToken"   dc:"服务商应用授权token"`
	ExpiresIn    *gtime.Time `json:"expiresIn"      dc:"Token过期时间"`
	ReExpiresIn  *gtime.Time `json:"reExpiresIn"    dc:"Token限期刷新时间"`
}

// UpdateThirdAppConfigReq 修改第三方应用基础信息
type UpdateThirdAppConfigReq struct {
	Id             int64  `json:"id"             dc:"授权商家id"`
	Name           string `json:"name"           dc:"服务商名称"`
	ExtJson        string `json:"extJson"        dc:"拓展字段"`
	AppGatewayUrl  string `json:"appGatewayUrl"  dc:"网关地址"`
	AppCallbackUrl string `json:"appCallbackUrl" dc:"回调地址"`
	AppSecret      string `json:"appSecret"      dc:"服务商应用密钥"`
	AuthInitUrl    string `json:"authInitUrl"    dc:"授权发起页域名"`
	ServerDomain   string `json:"serverDomain"   dc:"服务器域名"`
	BusinessDomain string `json:"businessDomain" dc:"业务域名"`
	AuthTestAppIds string `json:"authTestAppIds" dc:"授权测试应用列表"`
	PlatformSite   string `json:"platformSite"   dc:"平台官方"`
	Logo           string `json:"logo"           description:"服务商logo"`
}

// UpdateThirdAppConfigHttpsReq 修改
type UpdateThirdAppConfigHttpsReq struct {
	Id        int64  `json:"id"             dc:"授权商家id"`
	HttpsCert string `json:"httpsCert"      description:"域名证书"`
	HttpsKey  string `json:"httpsKey"       description:"域名私钥"`
}
