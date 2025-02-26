// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package weixin_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WeixinThirdAppConfig is the golang structure of table weixin_third_app_config for DAO operations like Where/Data.
type WeixinThirdAppConfig struct {
	g.Meta         `orm:"table:weixin_third_app_config, do:true"`
	Id             interface{} // 服务商id
	Name           interface{} // 服务商name
	AppId          interface{} // 服务商应用Id
	AppName        interface{} // 服务商应用名称
	AppType        interface{} // 服务商应用类型
	AppAuthToken   interface{} // 服务商应用授权token
	ExpiresIn      *gtime.Time // Token过期时间
	ReExpiresIn    *gtime.Time // Token限期刷新时间
	UnionMainId    interface{} // 关联主体id
	SysUserId      interface{} // 用户id
	ExtJson        interface{} // 拓展字段
	AppGatewayUrl  interface{} // 网关地址
	AppCallbackUrl interface{} // 回调地址
	AppSecret      interface{} // 服务商应用密钥
	MsgVerifyToken interface{} // 消息校验Token
	MsgEncryptKey  interface{} // 消息加密解密密钥
	AuthInitUrl    interface{} // 授权发起页域名
	ServerDomain   interface{} // 服务器域名
	BusinessDomain interface{} // 业务域名
	AuthTestAppids interface{} // 授权测试应用列表
	PlatformSite   interface{} // 平台官方
	Logo           interface{} // 服务商logo
	State          interface{} // 状态：0禁用 1启用
	ReleaseState   interface{} // 发布状态：0未发布 1已发布
	HttpsCert      interface{} // 域名证书
	HttpsKey       interface{} // 域名私钥
	UpdatedAt      *gtime.Time //
	AppIdMd5       interface{} // 应用id加密md5后的结果
	UserId         interface{} // 应用所属账号
	RefreshToken   interface{} // 刷新应用Token
}
