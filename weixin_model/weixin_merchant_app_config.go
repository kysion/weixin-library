package weixin_model

import "github.com/gogf/gf/v2/os/gtime"

type WeixinMerchantAppConfig struct {
	Id           int64       `json:"id"           dc:"授权商家id"`
	Name         string      `json:"name"         dc:"授权商家name"`
	AppId        int64       `json:"appId"        dc:"商家授权应用Id"`
	AppName      string      `json:"appName"      dc:"商家授权应用名称"`
	AppType      string      `json:"appType"      dc:"应用类型"`
	AppAuthToken string      `json:"appAuthToken" dc:"授权应用token"`
	IsFullProxy  int         `json:"isFullProxy"  dc:"是否全权委托待开发：0否 1是"`
	AuthState    int         `json:"authState"    dc:"授权状态"`
	ExpiresIn    *gtime.Time `json:"expiresIn"    dc:"生效时间"`
	ReExpiresIn  *gtime.Time `json:"reExpiresIn"  dc:"失效时间"`
	UserId       int64       `json:"userId"       dc:"用户账号id"`
	UnionMainId  int64       `json:"unionMainId"  dc:"关联主体id"`
	SysUserId    int64       `json:"sysUserId"    dc:"用户id"`
	Tokens       string      `json:"tokens"       dc:"token列表"`
	ExtJson      string      `json:"extJson"      dc:"拓展字段"`
}

type UpdateMerchantAppConfig struct {
	Id           int64  `json:"id"           dc:"授权商家id"`
	AppId        int64  `json:"appId"        description:"商家授权应用Id"`
	AppName      string `json:"appName"      description:"商家授权应用名称"`
	AppAuthToken string `json:"appAuthToken" description:"商家授权应用token"`
	AuthState    int    `json:"authState"    dc:"授权状态"`
	ExtJson      string `json:"extJson"      dc:"拓展字段"`
}
