package gateway

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/weixin-library/utility"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"log"
	"sort"
	"strings"

	do "github.com/kysion/weixin-library/weixin_model/weixin_do"
	"github.com/yitter/idgenerator-go/idgen"
	"time"
)

type sThirdAppConfig struct {
	redisCache *gcache.Cache
	Duration   time.Duration
}

func NewThirdAppConfig() *sThirdAppConfig {
	return &sThirdAppConfig{
		redisCache: gcache.New(),
	}
}

// GetThirdAppConfigByAppId 根据AppId查找第三方应用配置信息
func (s *sThirdAppConfig) GetThirdAppConfigByAppId(ctx context.Context, id string) (*weixin_model.WeixinThirdAppConfig, error) {
	data := weixin_model.WeixinThirdAppConfig{}

	err := dao.WeixinThirdAppConfig.Ctx(ctx).Where(do.WeixinThirdAppConfig{AppId: id}).Scan(&data)

	return &data, err
}

// GetThirdAppConfigById 根据id查找第三方应用配置信息
func (s *sThirdAppConfig) GetThirdAppConfigById(ctx context.Context, id int64) (*weixin_model.WeixinThirdAppConfig, error) {
	return daoctl.GetByIdWithError[weixin_model.WeixinThirdAppConfig](dao.WeixinThirdAppConfig.Ctx(ctx), id)
}

// CreateThirdAppConfig  创建第三方应用配置信息
func (s *sThirdAppConfig) CreateThirdAppConfig(ctx context.Context, info *weixin_model.WeixinThirdAppConfig) (*weixin_model.WeixinThirdAppConfig, error) {
	data := do.WeixinThirdAppConfig{}

	gconv.Struct(info, &data)

	data.Id = idgen.NextId()
	if data.ExtJson == "" {
		data.ExtJson = nil
	}
	affected, err := daoctl.InsertWithError(
		dao.WeixinThirdAppConfig.Ctx(ctx),
		data,
	)

	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "第三方应用配置信息创建失败", dao.WeixinThirdAppConfig.Table())
	}

	return s.GetThirdAppConfigById(ctx, gconv.Int64(data.Id))
}

// UpdateThirdAppConfig 更新第三方应用配置信息
func (s *sThirdAppConfig) UpdateThirdAppConfig(ctx context.Context, id int64, info *weixin_model.UpdateThirdAppConfig) (bool, error) {
	// 首先判断第三方应用配置信息是否存在
	consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinThirdAppConfig](dao.WeixinThirdAppConfig.Ctx(ctx), id)
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该第三方应用配置不存在", dao.WeixinThirdAppConfig.Table())
	}
	data := do.WeixinThirdAppConfig{}
	gconv.Struct(info, &data)

	model := dao.WeixinThirdAppConfig.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(model).OmitNilData().Where(do.WeixinThirdAppConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "第三方应用配置信息更新失败", dao.WeixinThirdAppConfig.Table())
	}

	return affected > 0, nil
}

// UpdateReleaseState 修改发布状态
func (s *sThirdAppConfig) UpdateReleaseState(ctx context.Context, id int64, releaseState int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinThirdAppConfig.Ctx(ctx).Data(do.WeixinThirdAppConfig{
		ReleaseState: releaseState,
	}).OmitNilData().Where(do.WeixinThirdAppConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "第三方应用配置状态修改失败", dao.WeixinThirdAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateState 修改应用状态
func (s *sThirdAppConfig) UpdateState(ctx context.Context, id int64, state int) (bool, error) {
	affected, err := daoctl.UpdateWithError(dao.WeixinThirdAppConfig.Ctx(ctx).Data(do.WeixinThirdAppConfig{
		State: state,
	}).OmitNilData().Where(do.WeixinThirdAppConfig{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "第三方应用配置状态修改失败", dao.WeixinThirdAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateAppAuthToken 更新Token  服务商应用授权token
func (s *sThirdAppConfig) UpdateAppAuthToken(ctx context.Context, info *weixin_model.UpdateAppAuthToken) (bool, error) {
	data := do.WeixinThirdAppConfig{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinThirdAppConfig.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinThirdAppConfig{AppId: info.AppId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "服务商应用Token修改失败", dao.WeixinThirdAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateAppConfig 修改服务商基础信息
func (s *sThirdAppConfig) UpdateAppConfig(ctx context.Context, info *weixin_model.UpdateThirdAppConfigReq) (bool, error) {
	data := do.WeixinThirdAppConfig{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinThirdAppConfig.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinThirdAppConfig{Id: info.Id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "服务商应用基础修改失败", dao.WeixinThirdAppConfig.Table())
	}
	return affected > 0, err
}

// UpdateAppConfigHttps 修改服务商应用Https配置
func (s *sThirdAppConfig) UpdateAppConfigHttps(ctx context.Context, info *weixin_model.UpdateThirdAppConfigHttpsReq) (bool, error) {
	data := do.WeixinThirdAppConfig{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinThirdAppConfig.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinThirdAppConfig{Id: info.Id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "服务商应用基础修改失败", dao.WeixinThirdAppConfig.Table())
	}
	return affected > 0, err
}

// WXCheckSignature 微信接入校验 设置Token需要验证
func (s *sGateway) WXCheckSignature(ctx context.Context, signature, timestamp, nonce, echostr string) string {
	// 与填写的服务器配置中的Token一致
	const Token = "comjditcokuaimk"
	fmt.Println(signature + "、" + timestamp + "、" + nonce + "、" + echostr)
	arr := []string{timestamp, nonce, Token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(Token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	sign := utility.Sha1(b.String())

	ok := utility.CheckSignature(sign, timestamp, nonce, Token)

	if !ok {
		log.Println("微信公众号接入校验失败!")
		return ""
	}

	log.Println("微信公众号接入校验成功!")

	g.RequestFromCtx(ctx).Response.Write(echostr)
	return echostr
}