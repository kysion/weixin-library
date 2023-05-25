package merchant

import (
	"context"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_model"
	service "github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
)

// AppVersionManager 小程序开发管理
var AppVersionManager = cAppVersionManager{}

type cAppVersionManager struct{}

// 上传代码

// 提交审核

// 审核撤销

// 版本回退

// 获取上传列表

// 查询指定版本审核状态

// 最新一次提审单的审核状态

// 代码审核结果推送 （配置的事件接收 URL）

// SubmitAppVersionAudit 提交应用版本审核
func (c *cAppVersionManager) SubmitAppVersionAudit(ctx context.Context, req *weixin_merchant_app_v1.SubmitAppVersionAuditReq) (*weixin_model.AppVersionAuditRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().SubmitAppVersionAudit(ctx, appId, &req.SubmitAppVersionAuditReq)

	return ret, err
}

// CancelAppVersionAudit 撤销应用版本审核
func (c *cAppVersionManager) CancelAppVersionAudit(ctx context.Context, req *weixin_merchant_app_v1.CancelAppVersionAuditReq) (*weixin_model.CancelAppVersionAuditRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	// TODO  需要检查是否需要撤销具体的版本
	ret, err := service.AppVersion().CancelAppVersionAudit(ctx, appId)

	return ret, err
}

// CancelAppVersion 退回开发版本
func (c *cAppVersionManager) CancelAppVersion(ctx context.Context, req *weixin_merchant_app_v1.CancelAppVersionReq) (*weixin_model.CancelAppVersionRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	// TODO  需要检查是否需要退回具体的版本
	ret, err := service.AppVersion().CancelAppVersion(ctx, appId, &req.CancelAppVersionReq)

	return ret, err
}

// QueryAppVersionList 查询小程序版本列表
func (c *cAppVersionManager) QueryAppVersionList(ctx context.Context, req *weixin_merchant_app_v1.QueryAppVersionListReq) (*weixin_model.QueryAppVersionListRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().QueryAppVersionList(ctx, appId)

	return ret, err
}

// GetAppVersionDetail 查询小程序版本详情
func (c *cAppVersionManager) GetAppVersionDetail(ctx context.Context, req *weixin_merchant_app_v1.GetAppVersionDetailReq) (*weixin_model.QueryAppVersionDetailRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().GetAppVersionDetail(ctx, appId)

	return ret, err
}

// GetAppLatestVersionAudit 最新一次提审单的审核状态
func (c *cAppVersionManager) GetAppLatestVersionAudit(ctx context.Context, req *weixin_merchant_app_v1.GetAppLatestVersionAuditReq) (*weixin_model.GetAppLatestVersionAuditRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().GetAppLatestVersionAudit(ctx, appId)

	return ret, err
}

// GetAllCategory  获取所有类目
func (c *cAppVersionManager) GetAllCategory(ctx context.Context, req *weixin_merchant_app_v1.GetAllCategoryReq) (*weixin_model.AppCategoryInfoRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().GetAllCategory(ctx, appId)

	return ret, err
}

// GetAccountVBasicInfo 获取小程序基本信息
func (c *cAppVersionManager) GetAccountVBasicInfo(ctx context.Context, req *weixin_merchant_app_v1.GetAccountVBasicInfoReq) (*weixin_model.AccountVBasicInfoRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().GetAccountVBasicInfo(ctx, appId)

	return ret, err
}

// UploadAppMediaToAudit 应用提审素材上传接口
func (c *cAppVersionManager) UploadAppMediaToAudit(ctx context.Context, req *weixin_merchant_app_v1.UploadAppMediaToAuditReq) (*weixin_model.UploadAppMediaToAuditRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().UploadAppMediaToAudit(ctx, appId, req.MediaPath)

	return ret, err
}

// CommitAppAuditCode 应用提审素材上传接口
func (c *cAppVersionManager) CommitAppAuditCode(ctx context.Context, req *weixin_merchant_app_v1.CommitAppAuditCodeReq) (*weixin_model.CommitAppAuditCodeRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().CommitAppAuditCode(ctx, appId, &req.CommitAppAuditCodeReq)

	return ret, err
}

// GetQrcode 获取小程序体验版二维码
func (c *cAppVersionManager) GetQrcode(ctx context.Context, _ *weixin_merchant_app_v1.GetQrcodeReq) (*weixin_model.ErrorCommonRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().GetQrcode(ctx, appId)

	return ret, err
}

// ReleaseApp 发布已通过审核的小程序
func (c *cAppVersionManager) ReleaseApp(ctx context.Context, _ *weixin_merchant_app_v1.ReleaseAppReq) (*weixin_model.ErrorCommonRes, error) {
	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.AppVersion().ReleaseApp(ctx, appId)

	return ret, err
}
