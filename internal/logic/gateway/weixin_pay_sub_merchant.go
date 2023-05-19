package gateway

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	do "github.com/kysion/weixin-library/weixin_model/weixin_do"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// 微信支付特约商户 （也就是服务商模式下的子商户）
type sPaySubMerchant struct {
}

func NewPaySubMerchant() *sPaySubMerchant {
	return &sPaySubMerchant{}
}

// GetPaySubMerchantById 根据id查找特约商户配置信息
func (s *sPaySubMerchant) GetPaySubMerchantById(ctx context.Context, id int64) (*weixin_model.WeixinPaySubMerchant, error) {
	return daoctl.GetByIdWithError[weixin_model.WeixinPaySubMerchant](dao.WeixinPaySubMerchant.Ctx(ctx), id)
}

// GetPaySubMerchantByAppId 根据AppId查找特约商户配置信息
func (s *sPaySubMerchant) GetPaySubMerchantByAppId(ctx context.Context, appId string) (*weixin_model.WeixinPaySubMerchant, error) {
	data := weixin_model.WeixinPaySubMerchant{}

	err := dao.WeixinPaySubMerchant.Ctx(ctx).Where(do.WeixinPaySubMerchant{SubAppid: appId}).Scan(&data)

	return &data, err
}

// GetPaySubMerchantByMchid 根据Mchid查找特约商户配置信息
func (s *sPaySubMerchant) GetPaySubMerchantByMchid(ctx context.Context, id int) (*weixin_model.WeixinPaySubMerchant, error) {
	data := weixin_model.WeixinPaySubMerchant{}

	err := dao.WeixinPaySubMerchant.Ctx(ctx).Where(do.WeixinPaySubMerchant{SubMchid: id}).Scan(&data)

	return &data, err
}

// GetPaySubMerchantBySysUserId  根据用户id查询特约商户配置信息
func (s *sPaySubMerchant) GetPaySubMerchantBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.WeixinPaySubMerchant, error) {
	result := weixin_model.WeixinPaySubMerchant{}
	model := dao.WeixinPaySubMerchant.Ctx(ctx)

	err := model.Where(dao.WeixinPaySubMerchant.Columns().SysUserId, sysUserId).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreatePaySubMerchant  创建特约商户配置信息
func (s *sPaySubMerchant) CreatePaySubMerchant(ctx context.Context, info *weixin_model.WeixinPaySubMerchant) (*weixin_model.WeixinPaySubMerchant, error) {
	data := do.WeixinPaySubMerchant{}

	gconv.Struct(info, &data)

	data.Id = idgen.NextId()

	affected, err := daoctl.InsertWithError(
		dao.WeixinPaySubMerchant.Ctx(ctx),
		data,
	)

	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "特约商户配置信息创建失败", dao.WeixinPaySubMerchant.Table())
	}

	return s.GetPaySubMerchantById(ctx, gconv.Int64(data.Id))
}

// UpdatePaySubMerchant 更新特约商户配置信息
func (s *sPaySubMerchant) UpdatePaySubMerchant(ctx context.Context, id int64, info *weixin_model.UpdatePaySubMerchant) (bool, error) {
	// 首先判断特约商户配置信息是否存在
	consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinPaySubMerchant](dao.WeixinPaySubMerchant.Ctx(ctx), id)
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该特约商户配置不存在", dao.WeixinPaySubMerchant.Table())
	}
	data := do.WeixinPaySubMerchant{}
	gconv.Struct(info, &data)

	model := dao.WeixinPaySubMerchant.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(data).OmitNilData().Where(do.WeixinPaySubMerchant{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "特约商户配置信息更新失败", dao.WeixinPaySubMerchant.Table())
	}

	return affected > 0, nil
}

// SetAuthPath 设置特约商户授权目录
func (s *sPaySubMerchant) SetAuthPath(ctx context.Context, info *weixin_model.SetSubMerchantAuthPath) (bool, error) {
	data := do.WeixinPaySubMerchant{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinPaySubMerchant.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinPaySubMerchant{SubMchid: info.SubMchid}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "特约商户Token修改失败", dao.WeixinPaySubMerchant.Table())
	}
	return affected > 0, err
}
