package merchant

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// 小程序开发管理
type sAppVersion struct{}

func init() {
	weixin_service.RegisterAppVersion(NewAppVersion())

}
func NewAppVersion() *sAppVersion {
	return &sAppVersion{}
}

// 上传代码  (前端)

// 提交审核

// 审核撤销

// 版本回退

// 获取上传列表

// 查询指定版本审核状态

// 最新一次提审单的审核状态

// 代码审核结果推送 （配置的事件接收 URL）

// SubmitAppVersionAudit 提交应用版本审核
func (s *sAppVersion) SubmitAppVersionAudit(ctx context.Context, appId string, info *weixin_model.SubmitAppVersionAuditReq) (*weixin_model.AppVersionAuditRes, error) {
	// POST https://api.weixin.qq.com/wxa/submit_audit?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}

	url := "https://api.weixin.qq.com/wxa/submit_audit?access_token=" + merchantApp.AppAuthToken

	reqData, _ := gjson.Encode(info)

	result := g.Client().PostContent(ctx, url, reqData)

	res := weixin_model.AppVersionAuditRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

// CancelAppVersionAudit 撤销应用版本审核
func (s *sAppVersion) CancelAppVersionAudit(ctx context.Context, appId string) (*weixin_model.CancelAppVersionAuditRes, error) {
	// GET https://api.weixin.qq.com/wxa/undocodeaudit?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/wxa/undocodeaudit?access_token=" + merchantApp.AppAuthToken
	result := g.Client().GetContent(ctx, url)

	res := weixin_model.CancelAppVersionAuditRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

// CancelAppVersion 退回开发版本
func (s *sAppVersion) CancelAppVersion(ctx context.Context, appId string, info *weixin_model.CancelAppVersionReq) (*weixin_model.CancelAppVersionRes, error) {
	// GET https://api.weixin.qq.com/wxa/revertcoderelease?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/wxa/revertcoderelease?access_token=" + merchantApp.AppAuthToken
	result := g.Client().GetContent(ctx, url)

	res := weixin_model.CancelAppVersionRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

// QueryAppVersionList 查询小程序版本列表,获取已上传的代码页面列表
func (s *sAppVersion) QueryAppVersionList(ctx context.Context, appId string) (*weixin_model.QueryAppVersionListRes, error) {
	// GET https://api.weixin.qq.com/wxa/get_page?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/wxa/get_page?access_token=" + merchantApp.AppAuthToken
	result := g.Client().GetContent(ctx, url)

	res := weixin_model.QueryAppVersionListRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

// GetAppVersionDetail 查询小程序版本详情
func (s *sAppVersion) GetAppVersionDetail(ctx context.Context, appId string) (*weixin_model.QueryAppVersionDetailRes, error) {
	// POST https://api.weixin.qq.com/wxa/getversioninfo?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}

	url := "https://api.weixin.qq.com/wxa/getversioninfo?access_token=" + merchantApp.AppAuthToken

	encode, _ := gjson.Encode(g.Map{})
	result := g.Client().PostContent(ctx, url, encode)

	res := weixin_model.QueryAppVersionDetailRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

// GetAppLatestVersionAudit 最新一次提审单的审核状态
func (s *sAppVersion) GetAppLatestVersionAudit(ctx context.Context, appId string) (*weixin_model.GetAppLatestVersionAuditRes, error) {
	//GET https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=" + merchantApp.AppAuthToken
	result := g.Client().GetContent(ctx, url)

	res := weixin_model.GetAppLatestVersionAuditRes{}
	gjson.DecodeTo(result, &res)

	return &res, err

}

// GetAllCategory 获取应用所有类目
func (s *sAppVersion) GetAllCategory(ctx context.Context, appId string) (*weixin_model.AppCategoryInfoRes, error) {
	// GET https://api.weixin.qq.com/wxa/get_category?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/wxa/get_category?access_token=" + merchantApp.AppAuthToken
	result := g.Client().GetContent(ctx, url)

	res := weixin_model.AppCategoryInfoRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

// GetAccountVBasicInfo 获取小程序基本信息
func (s *sAppVersion) GetAccountVBasicInfo(ctx context.Context, appId string) (*weixin_model.AccountVBasicInfoRes, error) {
	// GET https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=" + merchantApp.AppAuthToken
	result := g.Client().GetContent(ctx, url)

	res := weixin_model.AccountVBasicInfoRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

// UploadAppMediaToAudit 应用提审素材上传接口
func (s *sAppVersion) UploadAppMediaToAudit(ctx context.Context, appId string, mediaPath string) (*weixin_model.UploadAppMediaToAuditRes, error) {
	// POST https://api.weixin.qq.com/wxa/uploadmedia?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	//url := "https://api.weixin.qq.com/wxa/uploadmedia?access_token=" + merchantApp.AppAuthToken
	//result := g.Client().PostContent(ctx, url)

	mediaid, err := UploaImage(ctx, merchantApp.AppAuthToken, mediaPath)
	// img1.jpg nXZPp3Jc2FitVGuiCBYvyApsY0F4m9i9TiWaNEEvbrZt12B4r6VxjOSbsM_5PziGjR5OHwG9JoVMM9LHZWH44Q
	// img2.jpg nXZPp3Jc2FitVGuiCBYvyLLgBpcvY-K8t2Ujrc2wiznXRL0CJJOZK1TkCdv4H7UO75xfTgS9SgjS5BNYvj4LCQ
	// img3.jpg nXZPp3Jc2FitVGuiCBYvyKRuFaqLoLdHbDhv8tjxO-7rAPJm-yx8yoODFUwX379lOJy_iUINj2moHlncHlPQRw
	// testVicdeo.mp4 nXZPp3Jc2FitVGuiCBYvyD-T2v4Rc9QbBmBSJBBhugkmjiM2EqvwCXim7qYpBg4weCwV13bbv17gEKSLatj2jA
	return mediaid, err
}

const (
	WechatUploadMediaAPI = "https://api.weixin.qq.com/wxa/uploadmedia"
)

func UploaImage(ctx context.Context, token string, imagePath string) (*weixin_model.UploadAppMediaToAuditRes, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(imagePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", WechatUploadMediaAPI, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	urlQuery := req.URL.Query()

	if err != nil {
		return nil, err
	}
	urlQuery.Add("access_token", token)
	//urlQuery.Add("type", "image")
	req.URL.RawQuery = urlQuery.Encode()
	fmt.Println(req.URL)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	jsonbody, _ := ioutil.ReadAll(res.Body)
	media := weixin_model.UploadAppMediaToAuditRes{}
	err = json.Unmarshal(jsonbody, &media)
	if err != nil {
		return nil, err
	}
	if media.Mediaid == "" {
		err = sys_service.SysLogs().ErrorSimple(ctx, err, "素材上传失败！", "WeiXin-App-Version-Manager")
	}

	return &media, err
}

// CommitAppAuditCode 上传代码并生成体验版
func (s *sAppVersion) CommitAppAuditCode(ctx context.Context, appId string, info *weixin_model.CommitAppAuditCodeReq) (*weixin_model.CommitAppAuditCodeRes, error) {
	// POST https://api.weixin.qq.com/wxa/commit?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/wxa/commit?access_token=" + merchantApp.AppAuthToken

	reqData, _ := gjson.Encode(info)

	result := g.Client().PostContent(ctx, url, reqData)

	res := weixin_model.CommitAppAuditCodeRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}

//  GET https://api.weixin.qq.com/wxa/get_qrcode?access_token=ACCESS_TOKEN

// GetQrcode 获取小程序体验版二维码
func (s *sAppVersion) GetQrcode(ctx context.Context, appId string) (*weixin_model.CommitAppAuditCodeRes, error) {
	// GET https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/wxa/get_qrcode?access_token=" + merchantApp.AppAuthToken
	result := g.Client().GetContent(ctx, url)

	response := g.RequestFromCtx(ctx).Response
	response.Header().Set("Content-Type", "image/jpeg")
	response.Header().Set("Cache-Control", "no-cache")
	response.Header().Set("Connection", "keep-alive")

	//response.Write(result)

	response.WriteExit(result)

	res := weixin_model.CommitAppAuditCodeRes{}
	gjson.DecodeTo(result, &res)

	return &res, err
}