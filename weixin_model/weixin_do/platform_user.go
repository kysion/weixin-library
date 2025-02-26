// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package weixin_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlatformUser is the golang structure of table platform_user for DAO operations like Where/Data.
type PlatformUser struct {
	g.Meta         `orm:"table:platform_user, do:true"`
	Id             interface{} //
	FacilitatorId  interface{} // 服务商id
	OperatorId     interface{} // 运营商id
	MerchantId     interface{} // 商户id
	SysUserId      interface{} // 系统用户id
	PlatformType   interface{} // 平台类型：1支付宝、2微信、4抖音、8银联
	ThirdAppId     interface{} // 第三方平台AppId
	MerchantAppId  interface{} // 商家应用AppId
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	PlatformUserId interface{} // 平台用户唯一标识
	SysUserType    interface{} // 系统用户类型：0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心，64后台
}
