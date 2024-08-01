package gateway

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/idgen"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	do "github.com/kysion/weixin-library/weixin_model/weixin_do"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
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
	result, err := daoctl.GetByIdWithError[weixin_model.WeixinMerchantAppConfig](dao.WeixinMerchantAppConfig.Ctx(ctx), id)
	if err != nil {
		return nil, err
	}

	return result, err
}

// GetMerchantAppConfigByAppId 根据AppId查找商家应用配置信息
func (s *sMerchantAppConfig) GetMerchantAppConfigByAppId(ctx context.Context, id string) (*weixin_model.WeixinMerchantAppConfig, error) {
	data := weixin_model.WeixinMerchantAppConfig{}

	err := dao.WeixinMerchantAppConfig.Ctx(ctx).Where(do.WeixinMerchantAppConfig{AppId: id}).Scan(&data)
	if err != nil {
		return nil, err
	}

	if gtime.Now().After(data.ExpiresIn) { // 如果Token已经过期
		_, err := weixin_service.AppAuth().RefreshToken(ctx, id, data.ThirdAppId, data.RefreshToken)
		if err != nil {
			return &data, err
		}
	}

	return s.GetMerchantAppConfigById(ctx, data.Id)
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
func (s *sMerchantAppConfig) CreateMerchantAppConfig(ctx context.Context, info *weixin_model.WeixinMerchantAppConfig) (*weixin_model.WeixinMerchantAppConfig, error) {
	data := do.WeixinMerchantAppConfig{}

	// wxcaf4b7b8d6620f00
	appLen := len(info.AppId)
	subAppId := gstr.SubStr(info.AppId, 2, appLen)      // caf4b7b8d6620f00
	appIdBase32 := weixin_utility.HexToBase32(subAppId) // 十六进制转32进制
	appId := "wx" + appIdBase32                         // wxclt5nn3b643o0

	if info.ServerDomain != "" {
		info.AppGatewayUrl = info.ServerDomain + "/weixin/" + appId + "/gateway.services"
		info.AppCallbackUrl = info.ServerDomain + "/weixin/$APPID$/" + appId + "/gateway.callback"
		info.NotifyUrl = info.ServerDomain + "/weixin/" + appId + "/gateway.notify"
	} else if info.ServerDomain == "" {
		// 没指定服务器域名，默认使用当前服务器域名
		info.ServerDomain = "https://www.kuaimk.com"
		info.AppGatewayUrl = "https://www.kuaimk.com/weixin/" + appId + "/gateway.services"
		info.AppCallbackUrl = "https://www.kuaimk.com/weixin/$APPID$/" + appId + "/gateway.callback"
		info.NotifyUrl = "https://www.kuaimk.com/weixin/" + appId + "/gateway.notify"
	}

	gconv.Struct(info, &data)

	data.Id = idgen.NextId()
	if data.ExtJson == "" {
		data.ExtJson = nil
	}

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
func (s *sMerchantAppConfig) UpdateMerchantAppConfig(ctx context.Context, id int64, info *weixin_model.UpdateMerchantAppConfig) (bool, error) {
	// 首先判断商家应用配置信息是否存在
	consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinMerchantAppConfig](dao.WeixinMerchantAppConfig.Ctx(ctx), id)
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该商家应用配置不存在", dao.WeixinMerchantAppConfig.Table())
	}
	data := do.WeixinMerchantAppConfig{}
	gconv.Struct(info, &data)

	model := dao.WeixinMerchantAppConfig.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(data).OmitNilData().Where(do.WeixinMerchantAppConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用配置信息更新失败", dao.WeixinMerchantAppConfig.Table())
	}

	return affected > 0, nil
}

// UpdateState 修改应用状态
func (s *sMerchantAppConfig) UpdateState(ctx context.Context, id int64, state int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinMerchantAppConfig.Ctx(ctx).Data(do.WeixinMerchantAppConfig{
		State: state,
	}).OmitNilData().Where(do.WeixinMerchantAppConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用配置状态修改失败", dao.WeixinMerchantAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateAppAuth 修改应用授权信息 (绑定&解绑第三方服务商)
func (s *sMerchantAppConfig) UpdateAppAuth(ctx context.Context, appId string, thirdAppId, isFullProxy int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinMerchantAppConfig.Ctx(ctx).
		Where(do.WeixinMerchantAppConfig{
			AppId: appId,
		}).
		Data(do.WeixinMerchantAppConfig{
			ThirdAppId:  thirdAppId,
			IsFullProxy: isFullProxy,
		}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用配置授权信息修改失败", dao.WeixinMerchantAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateAppAuthToken 更新Token  商家应用授权token
func (s *sMerchantAppConfig) UpdateAppAuthToken(ctx context.Context, info *weixin_model.UpdateMerchantAppAuthToken) (bool, error) {
	data := do.WeixinMerchantAppConfig{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinMerchantAppConfig.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinMerchantAppConfig{AppId: info.AppId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用Token修改失败", dao.WeixinMerchantAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateAppConfig 修改商家基础信息
func (s *sMerchantAppConfig) UpdateAppConfig(ctx context.Context, info *weixin_model.UpdateMerchantAppConfigReq) (bool, error) {
	data := do.WeixinMerchantAppConfig{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinMerchantAppConfig.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinMerchantAppConfig{Id: info.Id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用基础修改失败", dao.WeixinMerchantAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateAppConfigHttps 修改商家应用Https配置
func (s *sMerchantAppConfig) UpdateAppConfigHttps(ctx context.Context, info *weixin_model.UpdateMerchantAppConfigHttpsReq) (bool, error) {
	data := do.WeixinMerchantAppConfig{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinMerchantAppConfig.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinMerchantAppConfig{Id: info.Id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商家应用基础修改失败", dao.WeixinMerchantAppConfig.Table())
	}
	return affected > 0, err
}

// GetPolicy 获取协议
func (s *sMerchantAppConfig) GetPolicy(ctx context.Context, appId string) (*weixin_model.GetPolicyRes, error) {
	res := weixin_model.GetPolicyRes{}

	err := dao.WeixinMerchantAppConfig.Ctx(ctx).Fields(dao.WeixinMerchantAppConfig.Columns().PrivacyPolicy, dao.WeixinMerchantAppConfig.Columns().UserPolicy).Where(do.WeixinMerchantAppConfig{
		AppId: appId,
	}).Scan(&res)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, nil, "该AppId商家应用不存在", dao.WeixinMerchantAppConfig.Table())
	}

	return &res, nil
}
