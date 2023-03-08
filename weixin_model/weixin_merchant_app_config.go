package weixin_model

import "github.com/gogf/gf/v2/os/gtime"

type WeixinMerchantAppConfig struct {
	Id             int64       `json:"id"             description:"商家id"`
	Name           string      `json:"name"           description:"商家name"`
	AppId          string      `json:"appId"          description:"商家应用Id"`
	AppName        string      `json:"appName"        description:"商家应用名称"`
	AppType        int         `json:"appType"        description:"应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店"`
	AppAuthToken   string      `json:"appAuthToken"   description:"应用token"`
	IsFullProxy    int         `json:"isFullProxy"    description:"是否全权委托待开发：0否 1是"`
	State          int         `json:"state"          description:"状态： 0禁用 1启用"`
	ExpiresIn      *gtime.Time `json:"expiresIn"      description:"Token过期时间"`
	ReExpiresIn    *gtime.Time `json:"reExpiresIn"    description:"Token限期刷新时间"`
	UserId         int64       `json:"userId"         description:"用户账号id"`
	UnionMainId    int64       `json:"unionMainId"    description:"关联主体id"`
	SysUserId      int64       `json:"sysUserId"      description:"用户id"`
	ExtJson        string      `json:"extJson"        description:"拓展字段"`
	AppGatewayUrl  string      `json:"appGatewayUrl"  description:"网关地址"`
	AppCallbackUrl string      `json:"appCallbackUrl" description:"回调地址"`
	AppSecret      string      `json:"appSecret"      description:"服务器应用密钥"`
	MsgVerfiyToken string      `json:"msgVerfiyToken" description:"消息校验Token"`
	MsgEncryptKey  string      `json:"msgEncryptKey"  description:"消息加密解密密钥"`
	MsgEncryptType int         `json:"msgEncryptType" description:"消息加密模式：1兼容模式 2明文模式 4安全模式"`
	BusinessDomain string      `json:"businessDomain" description:"业务域名"`
	JsDomain       string      `json:"jsDomain"       description:"JS接口安全域名"`
	AuthDomain     string      `json:"authDomain"     description:"网页授权域名"`
	Logo           string      `json:"logo"           description:"商家logo"`
	HttpsCert      string      `json:"httpsCert"      description:"域名证书"`
	HttpsKey       string      `json:"httpsKey"       description:"域名私钥"`
}

type UpdateMerchantAppConfig struct {
	Id             int64       `json:"id"             description:"商家id"`
	Name           string      `json:"name"           description:"商家name"`
	AppAuthToken   string      `json:"appAuthToken"   description:"应用token"`
	ExpiresIn      *gtime.Time `json:"expiresIn"      description:"Token过期时间"`
	AppGatewayUrl  string      `json:"appGatewayUrl"  description:"网关地址"`
	AppCallbackUrl string      `json:"appCallbackUrl" description:"回调地址"`
	AppSecret      string      `json:"appSecret"      description:"服务器应用密钥"`
	MsgVerfiyToken string      `json:"msgVerfiyToken" description:"消息校验Token"`
	MsgEncryptKey  string      `json:"msgEncryptKey"  description:"消息加密解密密钥"`
	MsgEncryptType int         `json:"msgEncryptType" description:"消息加密模式：1兼容模式 2明文模式 4安全模式"`
	BusinessDomain string      `json:"businessDomain" description:"业务域名"`
	JsDomain       string      `json:"jsDomain"       description:"JS接口安全域名"`
	AuthDomain     string      `json:"authDomain"     description:"网页授权域名"`
	Logo           string      `json:"logo"           description:"商家logo"`
}

type UpdateMerchantAppAuthToken struct {
	AppId        string      `json:"appId"          description:"商家应用Id"`
	AppAuthToken string      `json:"appAuthToken"   description:"应用token"`
	ExpiresIn    *gtime.Time `json:"expiresIn"      description:"Token过期时间"`
	ReExpiresIn  *gtime.Time `json:"reExpiresIn"    description:"Token限期刷新时间"`
}

// UpdateMerchantAppConfigReq 修改商家应用基础信息
type UpdateMerchantAppConfigReq struct {
	Id             int64  `json:"id"             description:"商家id"`
	Name           string `json:"name"           description:"商家name"`
	ExtJson        string `json:"extJson"        description:"拓展字段"`
	AppGatewayUrl  string `json:"appGatewayUrl"  description:"网关地址"`
	AppCallbackUrl string `json:"appCallbackUrl" description:"回调地址"`
	AppSecret      string `json:"appSecret"      description:"服务器应用密钥"`
	BusinessDomain string `json:"businessDomain" description:"业务域名"`
	JsDomain       string `json:"jsDomain"       description:"JS接口安全域名"`
	AuthDomain     string `json:"authDomain"     description:"网页授权域名"`
	Logo           string `json:"logo"           description:"商家logo"`
}

// UpdateMerchantAppConfigHttpsReq 修改Https文件
type UpdateMerchantAppConfigHttpsReq struct {
	Id        int64  `json:"id"             description:"商家id"`
	HttpsCert string `json:"httpsCert"      description:"域名证书"`
	HttpsKey  string `json:"httpsKey"       description:"域名私钥"`
}
