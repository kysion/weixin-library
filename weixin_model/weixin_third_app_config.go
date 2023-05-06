package weixin_model

import "github.com/gogf/gf/v2/os/gtime"

type WeixinThirdAppConfig struct {
	Id             int64       `json:"id"             description:"服务商id"`
	Name           string      `json:"name"           description:"服务商name"`
	AppId          string      `json:"appId"          description:"服务商应用Id"`
	AppName        string      `json:"appName"        description:"服务商应用名称"`
	AppType        int         `json:"appType"        description:"服务商应用类型"`
	AppAuthToken   string      `json:"appAuthToken"   description:"服务商应用授权token"`
	ExpiresIn      *gtime.Time `json:"expiresIn"      description:"Token过期时间"`
	ReExpiresIn    *gtime.Time `json:"reExpiresIn"    description:"Token限期刷新时间"`
	UnionMainId    int64       `json:"unionMainId"    description:"关联主体id"`
	SysUserId      int64       `json:"sysUserId"      description:"用户id"`
	ExtJson        string      `json:"extJson"        description:"拓展字段"`
	AppGatewayUrl  string      `json:"appGatewayUrl"  description:"网关地址"`
	AppCallbackUrl string      `json:"appCallbackUrl" description:"回调地址"`
	AppSecret      string      `json:"appSecret"      description:"服务商应用密钥"`
	MsgVerfiyToken string      `json:"msgVerfiyToken" description:"消息校验Token"`
	MsgEncryptKey  string      `json:"msgEncryptKey"  description:"消息加密解密密钥"`
	AuthInitUrl    string      `json:"authInitUrl"    description:"授权发起页域名"`
	ServerDomain   string      `json:"serverDomain"   description:"服务器域名"`
	BusinessDomain string      `json:"businessDomain" description:"业务域名"`
	AuthTestAppIds string      `json:"authTestAppIds" description:"授权测试应用列表"`
	PlatformSite   string      `json:"platformSite"   description:"平台官方"`
	Logo           string      `json:"logo"           description:"服务商logo"`
	State          int         `json:"state"          description:"状态：0禁用 1启用"`
	ReleaseState   int         `json:"releaseState"   description:"发布状态：0未发布 1已发布"`
	HttpsCert      string      `json:"httpsCert"      description:"域名证书"`
	HttpsKey       string      `json:"httpsKey"       description:"域名私钥"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
	AppIdMd5       string      `json:"appIdMd5"       description:"应用id加密md5后的结果"`
	UserId         int64       `json:"userId"         description:"应用所属账号"`
	RefreshToken   string      `json:"refreshToken"   description:"刷新应用Token"`
}

type UpdateThirdAppConfig struct {
	Id             int64       `json:"id"             dc:"服务商id"`
	Name           string      `json:"name"           dc:"服务商name" v:"required#服务商"`
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
	RefreshToken string      `json:"refreshToken"   description:"刷新应用Token"`
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
