package weixin_model

import "github.com/gogf/gf/v2/os/gtime"

type WeixinPaySubMerchant struct {
	Id                   int64       `json:"id"                   description:"ID"`
	SubMchid             int         `json:"subMchid"             description:"特约商户商户号"`
	SpMchid              int         `json:"spMchId"              description:"服务商商户号"`
	SubAppid             string      `json:"subAppid"             description:"特约商户App唯一标识ID"`
	SubAppName           string      `json:"subAppName"           description:"特约商户App名称"`
	SubAppType           int         `json:"subAppType"           description:"特约商户App类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店"`
	SubMerchantName      string      `json:"subMerchantName"      description:"特约商户公司名称"`
	SubMerchantShortName string      `json:"subMerchantShortName" description:"特约商户商家简称"`
	SysUserId            int64       `json:"sysUserId"            description:"特约商户用户ID"`
	UnionMainId          int64       `json:"unionMainId"          description:"特约商户用户主体"`
	UnionMainType        int         `json:"unionMainType"        description:"特约商户主体类型"`
	JsapiAuthPath        string      `json:"jsapiAuthPath"        description:"JSAPI支付授权目录"`
	H5AuthPath           string      `json:"h5AuthPath"           description:"H5支付授权目录"`
	UpdatedAt            *gtime.Time `json:"updatedAt"            description:""`
}

type PaySubMerchantRes WeixinPaySubMerchant

type UpdatePaySubMerchant struct {
	Id                   int64  `json:"id"                   description:"ID"`
	SpMchId              int    `json:"spMchId"              description:"服务商户号"`
	SubAppid             string `json:"subAppid"             description:"特约商户App唯一标识ID"`
	SubAppName           string `json:"subAppName"           description:"特约商户App名称"`
	SubAppType           int    `json:"subAppType"           description:"特约商户App类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店"`
	SubMerchantName      string `json:"subMerchantName"      description:"特约商户公司名称"`
	SubMerchantShortName string `json:"subMerchantShortName" description:"特约商户商家简称"`
	SysUserId            int64  `json:"sysUserId"            description:"特约商户用户ID"`
	UnionMainId          int64  `json:"unionMainId"          description:"特约商户用户主体"`
	UnionMainType        int    `json:"unionMainType"        description:"特约商户主体类型"`
}

type SetSubMerchantAuthPath struct {
	SubMchid      int     `json:"subMchid"             description:"特约商户商户号"`
	JsapiAuthPath *string `json:"jsapiAuthPath"     description:"JSAPI支付授权目录"`
	H5AuthPath    *string `json:"h5AuthPath"           description:"H5支付授权目录"`
}
