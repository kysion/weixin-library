// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WeixinPayMerchant is the golang structure of table weixin_pay_merchant for DAO operations like Where/Data.
type WeixinPayMerchant struct {
	g.Meta            `orm:"table:weixin_pay_merchant, do:true"`
	Id                interface{} // ID
	Mchid             interface{} // 微信支付商户号
	MerchantName      interface{} // 商户号公司名称
	MerchantShortName interface{} // 商户号简称
	MerchantType      interface{} // 商户号类型：1服务商、2商户
	ApiV3Key          interface{} // 用于ApiV3平台证书解密、回调信息解密
	ApiV2Key          interface{} // 用于ApiV2平台证书解密、回调信息解密
	PayCertP12        interface{} // 支付证书文件
	PayPublicKeyPem   interface{} // 公钥文件
	PayPrivateKeyPem  interface{} // 私钥文件
	CertSerialNumber  interface{} // 证书序列号
	JsapiAuthPath     interface{} // JSAPI支付授权目录
	SysUserId         interface{} // 用户ID
	UnionMainId       interface{} // 用户关联主体
	UnionMainType     interface{} // 用户类型
	BankcardAccount   interface{} // 银行结算账户,用于交易和提现
	UnionAppid        interface{} // 该商户号关联的AppId，微信支付接入模式属于直连模式，限制只能是同一主体下的App列表
	UpdatedAt         *gtime.Time //
	AppId             interface{} // 商户号 对应的公众号的服务号APPID
}