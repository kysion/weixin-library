package gateway

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	do "github.com/kysion/weixin-library/weixin_model/weixin_do"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

/*
消息的模板：
- 订阅消息 【小程序】
*/

type sSubscribeMessageTemplate struct {
	redisCache *gcache.Cache
	Duration   time.Duration
}

func NewSubscribeMessageTemplate() *sSubscribeMessageTemplate {
	return &sSubscribeMessageTemplate{
		redisCache: gcache.New(),
	}
}

// GetSubscribeMessageTemplateByTemplateId 根据模板ID查找消息模板
func (s *sSubscribeMessageTemplate) GetSubscribeMessageTemplateByTemplateId(ctx context.Context, templateId string) (*weixin_model.WeixinSubscribeMessageTemplateRes, error) {
	data := weixin_model.WeixinSubscribeMessageTemplateRes{}

	err := dao.WeixinSubscribeMessageTemplate.Ctx(ctx).Where(do.WeixinSubscribeMessageTemplate{TemplateId: templateId}).Scan(&data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetSubscribeMessageTemplateById 根据id查找消息模板信息
func (s *sSubscribeMessageTemplate) GetSubscribeMessageTemplateById(ctx context.Context, id int64) (*weixin_model.WeixinSubscribeMessageTemplateRes, error) {
	result, err := daoctl.GetByIdWithError[weixin_model.WeixinSubscribeMessageTemplateRes](dao.WeixinSubscribeMessageTemplate.Ctx(ctx), id)
	if err != nil {
		return nil, err
	}
	return result, err
}

// CreateSubscribeMessageTemplate  创建消息模板信息
func (s *sSubscribeMessageTemplate) CreateSubscribeMessageTemplate(ctx context.Context, info *weixin_model.WeixinSubscribeMessageTemplate) (*weixin_model.WeixinSubscribeMessageTemplateRes, error) {
	data := do.WeixinSubscribeMessageTemplate{}

	gconv.Struct(info, &data)

	data.Id = idgen.NextId()

	affected, err := daoctl.InsertWithError(
		dao.WeixinSubscribeMessageTemplate.Ctx(ctx),
		data,
	)

	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "消息模板信息创建失败", dao.WeixinSubscribeMessageTemplate.Table())
	}

	return s.GetSubscribeMessageTemplateById(ctx, gconv.Int64(data.Id))
}

// UpdateSubscribeMessageTemplate 更新消息模板信息
func (s *sSubscribeMessageTemplate) UpdateSubscribeMessageTemplate(ctx context.Context, id int64, info *weixin_model.UpdateWeixinSubscribeMessageTemplate) (bool, error) {
	// 首先判断消息模板信息是否存在
	consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinSubscribeMessageTemplate](dao.WeixinSubscribeMessageTemplate.Ctx(ctx), id)
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该消息模板不存在", dao.WeixinSubscribeMessageTemplate.Table())
	}
	data := do.WeixinSubscribeMessageTemplate{}
	gconv.Struct(info, &data)

	model := dao.WeixinSubscribeMessageTemplate.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(data).OmitNilData().Where(do.WeixinSubscribeMessageTemplate{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消息模板信息更新失败", dao.WeixinSubscribeMessageTemplate.Table())
	}

	return affected > 0, nil
}

// DeleteSubscribeMessageTemplate 删除模板
func (s *sSubscribeMessageTemplate) DeleteSubscribeMessageTemplate(ctx context.Context, appId string, templateId string) (bool, error) {
	info, err := s.GetSubscribeMessageTemplateByTemplateId(ctx, templateId)
	if info == nil || err != nil {
		return false, err
	}

	affected, err := daoctl.DeleteWithError(dao.WeixinSubscribeMessageTemplate.Ctx(ctx).Where(do.WeixinSubscribeMessageTemplate{MerchantAppId: appId, TemplateId: templateId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "消息模板信息删除失败", dao.WeixinSubscribeMessageTemplate.Table())
	}

	return affected > 0, nil
}

// QuerySubscribeMessageTemplate 查询模板｜列表
func (s *sSubscribeMessageTemplate) QuerySubscribeMessageTemplate(ctx context.Context, params *base_model.SearchParams, isExport bool) (*weixin_model.WeixinSubscribeMessageTemplateListRes, error) {
	result := &weixin_model.WeixinSubscribeMessageTemplateListRes{}

	model := dao.WeixinSubscribeMessageTemplate.Ctx(ctx)

	response, err := daoctl.Query[weixin_model.WeixinSubscribeMessageTemplateRes](model, params, isExport)
	if err != nil {
		return result, err
	}
	result.Records = response.Records
	result.PaginationRes = response.PaginationRes

	return result, err
}

// GetSubscribeMessageTemplateByAppAndSceneTypeAndMessageType 查询模板
func (s *sSubscribeMessageTemplate) GetSubscribeMessageTemplateByAppAndSceneTypeAndMessageType(ctx context.Context, appId string, appType int, sceneType, messageType int) (*weixin_model.WeixinSubscribeMessageTemplateRes, error) {
	result := weixin_model.WeixinSubscribeMessageTemplateRes{}
	err := dao.WeixinSubscribeMessageTemplate.Ctx(ctx).Where(do.WeixinSubscribeMessageTemplate{
		MerchantAppId:   appId,
		MerchantAppType: appType,
		SceneType:       sceneType,
		MessageType:     messageType,
	}).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
