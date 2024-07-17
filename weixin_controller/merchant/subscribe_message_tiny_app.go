package merchant

import (
	"context"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_merchant_app_v1"
	"github.com/kysion/weixin-library/weixin_model"
	service "github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
)

// SubscribeMessage 小程序订阅消息
var SubscribeMessage = cSubscribeMessage{}

type cSubscribeMessage struct{}

// SendMessage 发送订阅消息
func (c *cSubscribeMessage) SendMessage(ctx context.Context, req *weixin_merchant_app_v1.SendMessageReq) (*weixin_model.SendMessageRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.SubscribeMessage().SendMessage(ctx, appId, &req.SendMessage)

	return ret, err
}

// GetCategory 获取小程序账号的类目
func (c *cSubscribeMessage) GetCategory(ctx context.Context, _ *weixin_merchant_app_v1.GetCategoryReq) (*weixin_model.GetCategoryRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.SubscribeMessage().GetCategory(ctx, appId)

	return ret, err
}

// GetMyTemplateList 获取个人模板列表
func (c *cSubscribeMessage) GetMyTemplateList(ctx context.Context, _ *weixin_merchant_app_v1.GetMyTemplateListReq) (*weixin_model.GetMyTemplateListRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.SubscribeMessage().GetMyTemplateList(ctx, appId)

	return ret, err
}

// DeleteTemplate 删除模板
func (c *cSubscribeMessage) DeleteTemplate(ctx context.Context, req *weixin_merchant_app_v1.DeleteTemplateReq) (*weixin_model.DeleteTemplateRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.SubscribeMessage().DeleteTemplate(ctx, appId, &weixin_model.DeleteTemplate{PriTmplId: req.TemplateId})

	return ret, err
}

// GetPubTemplateKeyWords 获取模板的关键词列表
func (c *cSubscribeMessage) GetPubTemplateKeyWords(ctx context.Context, req *weixin_merchant_app_v1.GetPubTemplateKeyWordsReq) (*weixin_model.GetPubTemplateKeyWordsRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.SubscribeMessage().GetPubTemplateKeyWords(ctx, appId, req.TemplateId)

	return ret, err
}

// GetPubTemplateTitleList 获取指定类目下的公共模板列表
func (c *cSubscribeMessage) GetPubTemplateTitleList(ctx context.Context, _ *weixin_merchant_app_v1.GetPubTemplateTitleListReq) (*weixin_model.GetPubTemplateTitleListRes, error) {

	appId := weixin_utility.GetAppIdFormContext(ctx)

	ret, err := service.SubscribeMessage().GetPubTemplateTitleList(ctx, appId)

	return ret, err
}
