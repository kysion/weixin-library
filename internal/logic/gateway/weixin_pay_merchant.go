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

// 微信支付商户号
type sPayMerchant struct {
}

func NewPayMerchant() *sPayMerchant {
	return &sPayMerchant{}
}

// GetPayMerchantById 根据id查找商户号配置信息
func (s *sPayMerchant) GetPayMerchantById(ctx context.Context, id int64) (*weixin_model.PayMerchant, error) {
	return daoctl.GetByIdWithError[weixin_model.PayMerchant](dao.WeixinPayMerchant.Ctx(ctx), id)
}

// GetPayMerchantByMchid 根据Mchid查找商户号配置信息
func (s *sPayMerchant) GetPayMerchantByMchid(ctx context.Context, id int) (*weixin_model.PayMerchant, error) {
	data := weixin_model.PayMerchant{}

	err := dao.WeixinPayMerchant.Ctx(ctx).Where(do.WeixinPayMerchant{Mchid: id}).Scan(&data)

	return &data, err
}

// GetPayMerchantBySysUserId  根据商家id查询商户号配置信息
func (s *sPayMerchant) GetPayMerchantBySysUserId(ctx context.Context, sysUserId int64) (*weixin_model.PayMerchant, error) {
	result := weixin_model.PayMerchant{}
	model := dao.WeixinPayMerchant.Ctx(ctx)

	err := model.Where(dao.WeixinPayMerchant.Columns().SysUserId, sysUserId).Scan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreatePayMerchant  创建商户号配置信息
func (s *sPayMerchant) CreatePayMerchant(ctx context.Context, info *weixin_model.PayMerchant) (*weixin_model.PayMerchant, error) {
	data := do.WeixinPayMerchant{}

	gconv.Struct(info, &data)

	data.Id = idgen.NextId()
	if data.UnionAppid == "" {
		data.UnionAppid = nil
	}

	affected, err := daoctl.InsertWithError(
		dao.WeixinPayMerchant.Ctx(ctx),
		data,
	)

	if affected == 0 || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "商户号配置信息创建失败", dao.WeixinPayMerchant.Table())
	}

	return s.GetPayMerchantById(ctx, gconv.Int64(data.Id))
}

// UpdatePayMerchant 更新商户号配置信息
func (s *sPayMerchant) UpdatePayMerchant(ctx context.Context, id int64, info *weixin_model.UpdatePayMerchant) (bool, error) {
	// 首先判断商户号配置信息是否存在
	consumerInfo, err := daoctl.GetByIdWithError[entity.WeixinPayMerchant](dao.WeixinPayMerchant.Ctx(ctx), id)
	if err != nil || consumerInfo == nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "该商户号配置不存在", dao.WeixinPayMerchant.Table())
	}
	data := do.WeixinPayMerchant{}
	gconv.Struct(info, &data)

	model := dao.WeixinPayMerchant.Ctx(ctx)
	affected, err := daoctl.UpdateWithError(model.Data(data).OmitNilData().Where(do.WeixinPayMerchant{Id: id}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商户号配置信息更新失败", dao.WeixinPayMerchant.Table())
	}

	return affected > 0, nil
}

// SetCertAndKey  设置商户号证书及密钥文件
func (s *sPayMerchant) SetCertAndKey(ctx context.Context, mchId int64, info *weixin_model.SetCertAndKey) (bool, error) {
	data := do.WeixinPayMerchant{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinPayMerchant.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinPayMerchant{Mchid: mchId}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "设置商户号证书及密钥文件失败", dao.WeixinPayMerchant.Table())
	}
	return affected > 0, err
}

// SetAuthPath 设置商户号授权目录
func (s *sPayMerchant) SetAuthPath(ctx context.Context, info *weixin_model.SetAuthPath) (bool, error) {
	data := do.WeixinPayMerchant{}
	gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinPayMerchant.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinPayMerchant{Mchid: info.Mchid}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商户号授权目录修改失败", dao.WeixinPayMerchant.Table())
	}
	return affected > 0, err
}

// SetPayMerchantUnionId 设置商户号关联的AppId
func (s *sPayMerchant) SetPayMerchantUnionId(ctx context.Context, info *weixin_model.SetPayMerchantUnionId) (bool, error) {
	// TODO AppId需要进行排查，已经存在的不设置，不存在的进行设置，是一个数组结构 "["wx13r3453534","wx908989fsf7s9f9s"]"
	//data := do.WeixinPayMerchant{}
	//gconv.Struct(info, &data)
	//
	//affected, err := daoctl.UpdateWithError(dao.WeixinPayMerchant.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinPayMerchant{Id: info.Id}))
	//
	//if err != nil {
	//	return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商户号基础修改失败", dao.WeixinPayMerchant.Table())
	//}
	//return affected > 0, err
	return true, nil
}

// SetBankcardAccount 设置商户号银行卡号
func (s *sPayMerchant) SetBankcardAccount(ctx context.Context, info *weixin_model.SetBankcardAccount) (bool, error) {
	//data := do.WeixinPayMerchant{}
	//gconv.Struct(info, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinPayMerchant.Ctx(ctx).Data(info).OmitNilData().Where(do.WeixinPayMerchant{Mchid: info.Mchid}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "设置商户号银行卡号失败", dao.WeixinPayMerchant.Table())
	}
	return affected > 0, err
}
