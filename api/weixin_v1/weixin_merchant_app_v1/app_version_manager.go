package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

// 小程序开发管理

type SubmitAppVersionAuditReq struct {
	g.Meta `path:"/:appId/submitAppVersionAudit" method:"post" summary:"提交应用版本审核" tags:"WeiXin小程序管理"`

	// TODO 需要修改微信的

	weixin_model.SubmitAppVersionAuditReq
}

type CancelAppVersionAuditReq struct {
	g.Meta `path:"/:appId/cancelAppVersionAudit" method:"get" summary:"撤销应用版本审核" tags:"WeiXin小程序管理"`
	//AppVersion string `json:"app_version" dc:"版本号"`
}

type CancelAppVersionReq struct {
	g.Meta `path:"/:appId/cancelAppVersion" method:"get" summary:"退回开发版本" tags:"WeiXin小程序管理"`
	//AppVersion string `json:"app_version" dc:"版本号"`
	weixin_model.CancelAppVersionReq
}

type QueryAppVersionListReq struct {
	g.Meta `path:"/:appId/queryAppVersionList" method:"post" summary:"小程序版本列表查询" tags:"WeiXin小程序管理"`
	// query请求URL传参参数
	//BundleId      string `json:"bundle_id" dc:"端参数"`

	// TODO 需要修改微信的
	VersionStatus string `json:"version_status" dc:"版本状态列表，用英文逗号,分割，不填默认不返回，说明如下：INIT: 开发中, AUDITING: 审核中, AUDIT_REJECT: 审核驳回, WAIT_RELEASE: 待上架, BASE_AUDIT_PASS: 准入不可营销, GRAY: 灰度中, RELEASE: 已上架, OFFLINE: 已下架, AUDIT_OFFLINE: 已下架;"`
}

type GetAppVersionDetailReq struct {
	g.Meta     `path:"/:appId/getAppVersionDetail" method:"post" summary:"小程序版本详情查询" tags:"WeiXin小程序管理"`
	AppVersion string `json:"app_version" dc:"版本号"`
}

type GetAppLatestVersionAuditReq struct {
	g.Meta     `path:"/:appId/getAppLatestVersionAudit" method:"post" summary:"最新一次提审单的审核状态" tags:"WeiXin小程序管理"`
	AppVersion string `json:"app_version" dc:"版本号"`
}

type GetAllCategoryReq struct {
	g.Meta `path:"/:appId/getAllCategory" method:"post" summary:"获取小程序所有类目" tags:"WeiXin小程序管理"`
}

type GetAccountVBasicInfoReq struct {
	g.Meta `path:"/:appId/getAccountVBasicInfo" method:"post" summary:"获取小程序基本信息" tags:"WeiXin小程序管理"`
}

type UploadAppMediaToAuditReq struct {
	g.Meta    `path:"/:appId/uploadAppMediaToAudit" method:"post" summary:"App提审素材上传" tags:"WeiXin小程序管理"`
	MediaPath string `json:"media_path" dc:"素材路径"`
}

type CommitAppAuditCodeReq struct {
	g.Meta `path:"/:appId/commitAppAuditCode" method:"post" summary:"上传代码" tags:"WeiXin小程序管理"`

	weixin_model.CommitAppAuditCodeReq
}

type GetQrcodeReq struct {
	g.Meta `path:"/:appId/getQrcode" method:"post" summary:"获取体验版二维码" tags:"WeiXin小程序管理"`
}
