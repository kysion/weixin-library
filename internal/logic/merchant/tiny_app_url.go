package merchant

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

/*
小程序码&小程序链接 ：

- 小程序码： 文档 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getQRCode.html
- URL Scheme ： 文档 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/queryScheme.html
- URL Link ： 文档 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-link/generateUrlLink.html
- Short Link ： 文档 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/short-link/generateShortLink.html
*/

type sTinyAppUrl struct{}

func NewTinyAppUrl() weixin_service.ITinyAppUrl {
	result := &sTinyAppUrl{}

	return result
}

// GenerateScheme 获取scheme码 【加密 URL Scheme】
func (s *sTinyAppUrl) GenerateScheme(ctx context.Context, appId string, info *weixin_model.JumpWxa) (*weixin_model.GetSchemeRes, error) {
	// POST https://api.weixin.qq.com/wxa/generatescheme?access_token=ACCESS_TOKEN

	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}

	url := "https://api.weixin.qq.com/wxa/generatescheme?access_token=" + merchantApp.AppAuthToken

	encode, _ := gjson.Encode(g.Map{
		"jump_wxa":        info,
		"is_expire":       true,
		"expire_type":     1,
		"expire_interval": 1,
	})

	result := g.Client().PostContent(ctx, url, encode)

	res := weixin_model.GetSchemeRes{}
	_ = gjson.DecodeTo(result, &res)

	return &res, err
}

// GeneratePubScheme 获取scheme码 【明文 URL Scheme】
func (s *sTinyAppUrl) GeneratePubScheme(ctx context.Context, appId string, info *weixin_model.JumpWxa) (*weixin_model.GetSchemeRes, error) {
	// weixin://dl/business/?appid=*APPID*&path=*PATH*&query=*QUERY*&env_version=*ENV_VERSION*

	link := "weixin://dl/business/?appid=" + appId + "&path=" + info.Path + "&query=" + info.Query + "&env_version=" + info.EnvVersion

	res := &weixin_model.GetSchemeRes{
		Errcode:  0,
		Errmsg:   "ok",
		Openlink: link,
	}

	return res, nil
}
