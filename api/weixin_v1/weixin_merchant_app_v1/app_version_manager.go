package weixin_merchant_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

// 小程序开发管理

type SubmitAppVersionAuditReq struct {
	g.Meta `path:"/:appId/submitAppVersionAudit" method:"post" summary:"提交应用版本代码审核" tags:"WeiXin小程序管理"`

	weixin_model.SubmitAppVersionAuditReq
}

type CancelAppVersionAuditReq struct {
	g.Meta `path:"/:appId/cancelAppVersionAudit" method:"get" summary:"撤销应用版本审核" tags:"WeiXin小程序管理"`
}

type CancelAppVersionReq struct {
	g.Meta `path:"/:appId/cancelAppVersion" method:"get" summary:"退回开发版本" tags:"WeiXin小程序管理"`
	weixin_model.CancelAppVersionReq
}

type QueryAppVersionListReq struct {
	g.Meta `path:"/:appId/queryAppVersionList" method:"post" summary:"小程序已上传的代码页面｜列表" tags:"WeiXin小程序管理"`
}

type GetAppVersionDetailReq struct {
	g.Meta `path:"/:appId/getAppVersionDetail" method:"post" summary:"小程序版本详情查询" tags:"WeiXin小程序管理"`
}

type GetAppLatestVersionAuditReq struct {
	g.Meta `path:"/:appId/getAppLatestVersionAudit" method:"post" summary:"最新一次提审单的审核状态" tags:"WeiXin小程序管理"`
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
	g.Meta `path:"/:appId/commitAppAuditCode" method:"post" summary:"上传代码并生成体验版" tags:"WeiXin小程序管理"`

	weixin_model.CommitAppAuditCodeReq
}

type GetQrcodeReq struct {
	g.Meta `path:"/:appId/getQrcode" method:"post" summary:"获取体验版二维码" tags:"WeiXin小程序管理"`
}

type ReleaseAppReq struct {
	g.Meta `path:"/:appId/releaseApp" method:"post" summary:"发布已过审的小程序" tags:"WeiXin小程序管理"`
}
