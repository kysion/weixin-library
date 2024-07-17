package weixin_controller

import (
	"context"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/weixin-library/api/weixin_v1/wexin_subscribe_message_template_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_dao"
	"github.com/kysion/weixin-library/weixin_service"
)

// SubscribeMessageTemplate 小程序订阅消息模板管理
var SubscribeMessageTemplate = cSubscribeMessageTemplate{}

type cSubscribeMessageTemplate struct{}

// QueryMessageTemplate 获取订阅消息模板｜列表
func (c *cSubscribeMessageTemplate) QueryMessageTemplate(ctx context.Context, req *wexin_subscribe_message_template_v1.QueryMessageTemplateReq) (*weixin_model.WeixinSubscribeMessageTemplateListRes, error) {

	req.SearchParams.Filter = append(req.SearchParams.Filter, base_model.FilterInfo{
		Field: weixin_dao.WeixinSubscribeMessageTemplate.Columns().MerchantAppId,
		Where: "=",
		Value: req.AppId,
	})

	ret, err := weixin_service.SubscribeMessageTemplate().QuerySubscribeMessageTemplate(ctx, &req.SearchParams, req.IsExport)

	return ret, err
}

// GetMessageTemplateById 获取订阅消息模板｜列表
func (c *cSubscribeMessageTemplate) GetMessageTemplateById(ctx context.Context, req *wexin_subscribe_message_template_v1.GetMessageTemplateByIdReq) (*weixin_model.WeixinSubscribeMessageTemplateRes, error) {
	ret, err := weixin_service.SubscribeMessageTemplate().GetSubscribeMessageTemplateById(ctx, req.Id)

	return ret, err
}
