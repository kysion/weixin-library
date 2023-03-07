package gateway

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	do "github.com/kysion/weixin-library/weixin_model/weixin_do"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

// 关于微信商家的应用配置信息
type sMerchantAppConfig struct {
	redisCache *gcache.Cache
	Duration   time.Duration
}

func NewMerchantAppConfig() *sMerchantAppConfig {
	return &sMerchantAppConfig{
		redisCache: gcache.New(),
	}
}

// GetMerchantAppConfigById 根据id查找商家应用配置信息
func (s *sMerchantAppConfig) GetMerchantAppConfigById(ctx context.Context, id int64) (*weixin_model.WeixinMerchantAppConfig, error) {
	return daoctl.GetByIdWithError[weixin_model.WeixinMerchantAppConfig](dao.WeixinMerchantAppConfig.Ctx(ctx), id)
}

// GetMerchantAppConfigBySysUserId  根据商家id查询商家应用配置信息
func (s *sMerchantAppConfig) GetMerchantAppConfigBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.WeixinMerchantAppConfig, error) {
	result := weixin_model.WeixinMerchantAppConfig{}
	model := dao.WeixinMerchantAppConfig.Ctx(ctx)

	err := model.Where(dao.WeixinMerchantAppConfig.Columns().SysUserId, sysUserId).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateMerchantAppConfig  创建商家应用配置信息
func (s *sMerchantAppConfig) CreateMerchantAppConfig(ctx context.Context, info weixin_model.WeixinMerchantAppConfig) (*weixin_model.WeixinMerchantAppConfig, error) {
	data := do.WeixinMerchantAppConfig{}

	gconv.Struct(info, &data)

	data.Id = idgen.NextId()
	data.AuthState = 1 // 授权状态默认正常
	affected, err := daoctl.InsertWithError(
		dao.WeixinMerchantAppConfig.Ctx(ctx),
		data,
	)

	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用配置信息创建失败", dao.WeixinMerchantAppConfig.Table())
	}

	return s.GetMerchantAppConfigById(ctx, gconv.Int64(data.Id))
}

// UpdateMerchantAppConfig 更新商家应用配置信息
func (s *sMerchantAppConfig) UpdateMerchantAppConfig(ctx context.Context, id int64, info weixin_model.UpdateMerchantAppConfig) (bool, error) {
	// 首先判断商家应用配置信息是否存在
	consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinMerchantAppConfig](dao.WeixinMerchantAppConfig.Ctx(ctx), id)
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该商家应用配置不存在", dao.WeixinMerchantAppConfig.Table())
	}
	data := do.WeixinMerchantAppConfig{}
	gconv.Struct(info, &data)

	model := dao.WeixinMerchantAppConfig.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(model).OmitNilData().Where(do.WeixinMerchantAppConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用配置信息更新失败", dao.WeixinMerchantAppConfig.Table())
	}

	return affected > 0, nil
}

// UpdateMerchantAppConfigAuthState 修改商家授权状态
func (s *sMerchantAppConfig) UpdateMerchantAppConfigAuthState(ctx context.Context, id int64, authState int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinMerchantAppConfig.Ctx(ctx).Data(do.WeixinMerchantAppConfig{
		AuthState: authState,
	}).OmitNilData().Where(do.WeixinMerchantAppConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用配置状态修改失败", dao.WeixinMerchantAppConfig.Table())
	}
	return affected > 0, err
}
