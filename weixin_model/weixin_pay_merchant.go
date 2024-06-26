package weixin_model

import "github.com/gogf/gf/v2/os/gtime"

type PayMerchant struct {
	Id                int64       `json:"id"                description:"ID"`
	Mchid             int         `json:"mchid"             description:"微信支付商户号"`
	MerchantName      string      `json:"merchantName"      description:"商户号公司名称"`
	MerchantShortName string      `json:"merchantShortName" description:"商户号简称"`
	MerchantType      int         `json:"merchantType"      description:"商户号类型：1服务商、2商户、4门店商家"`
	ApiV3Key          string      `json:"apiV3Key"          description:"用于ApiV3平台证书解密、回调信息解密"`
	ApiV2Key          string      `json:"apiV2Key"          description:"用于ApiV2平台证书解密、回调信息解密"`
	PayCertP12        string      `json:"payCertP12"        description:"支付证书文件"`
	PayPublicKeyPem   string      `json:"payPublicKeyPem"   description:"公钥文件"`
	PayPrivateKeyPem  string      `json:"payPrivateKeyPem"  description:"私钥文件"`
	CertSerialNumber  string      `json:"certSerialNumber"  description:"证书序列号"`
	JsapiAuthPath     string      `json:"jsapiAuthPath"     description:"JSAPI支付授权目录"`
	SysUserId         int64       `json:"sysUserId"         description:"用户ID"`
	UnionMainId       int64       `json:"unionMainId"       description:"用户关联主体"`
	UnionMainType     int         `json:"unionMainType"     description:"用户类型"`
	BankcardAccount   string      `json:"bankcardAccount"   description:"银行结算账户,用于交易和提现"`
	UnionAppid        []string    `json:"unionAppid"        description:"该商户号关联的AppId，微信支付接入模式属于直连模式，限制只能是同一主体下的App列表"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         description:""`
	AppId             string      `json:"appId"             description:"商户号 对应的公众号的服务号APPID"`
}

type PayMerchantRes PayMerchant

type UpdatePayMerchant struct {
	Id                int64  `json:"id"             description:"商家id"`
	MerchantName      string `json:"merchantName"      description:"商户号公司名称"`
	MerchantShortName string `json:"merchantShortName" description:"商户号简称"`
}

type SetPayMerchantUnionId struct {
	Mchid      int      `json:"mchid"             description:"微信支付商户号"`
	UnionAppid []string `json:"unionAppid"        description:"该商户号关联的AppId，微信支付接入模式属于直连模式，限制只能是同一主体下的App列表"`
}

type SetCertAndKey struct {
	ApiV3Key         *string `json:"apiV3Key"          description:"用于ApiV3平台证书解密、回调信息解密"`
	ApiV2Key         *string `json:"apiV2Key"          description:"用于ApiV2平台证书解密、回调信息解密"`
	PayCertP12       string  `json:"payCertP12"        description:"支付证书文件"`
	PayPublicKeyPem  string  `json:"payPublicKeyPem"   description:"公钥文件"`
	PayPrivateKeyPem string  `json:"payPrivateKeyPem"  description:"私钥文件"`
	CertSerialNumber *string `json:"certSerialNumber"  description:"证书序列号"`
}

type SetAuthPath struct {
	Mchid         int    `json:"mchid"             description:"微信支付商户号"`
	JsapiAuthPath string `json:"jsapiAuthPath"     description:"JSAPI支付授权目录"`
}

type SetBankcardAccount struct {
	Mchid           int    `json:"mchid"             description:"微信支付商户号"`
	BankcardAccount string `json:"bankcardAccount"   description:"银行结算账户,用于交易和提现"`
}
