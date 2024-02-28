package gateway

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/weixin-library/weixin_consts"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	do "github.com/kysion/weixin-library/weixin_model/weixin_do"
	entity "github.com/kysion/weixin-library/weixin_model/weixin_entity"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_utility/file"
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
	result, err := daoctl.GetByIdWithError[weixin_model.PayMerchant](dao.WeixinPayMerchant.Ctx(ctx), id)
	if err != nil {
		return nil, err
	}

	return result, err
}

// GetPayMerchantByAppId 根据AppId查找商户号配置信息
func (s *sPayMerchant) GetPayMerchantByAppId(ctx context.Context, appId string) (*weixin_model.PayMerchant, error) {
	data := weixin_model.PayMerchant{}

	err := dao.WeixinPayMerchant.Ctx(ctx).Where(do.WeixinPayMerchant{AppId: appId, MerchantType: weixin_enum.Pay.MerchantType.SubMerchant.Code()}).Scan(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// GetPayMerchantByMchid 根据Mchid查找商户号配置信息
func (s *sPayMerchant) GetPayMerchantByMchid(ctx context.Context, id int) (*weixin_model.PayMerchant, error) {
	data := weixin_model.PayMerchant{}

	err := dao.WeixinPayMerchant.Ctx(ctx).Where(do.WeixinPayMerchant{Mchid: id}).Scan(&data)
	if err != nil {
		return nil, err
	}
	
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
	if len(info.UnionAppid) == 0 {
		data.UnionAppid = nil
	} else {
		// 生产指定切片，并去重
		data.UnionAppid = garray.NewSortedStrArrayFrom(info.UnionAppid).Unique().Slice()
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

// 从配置文件中获取文件路径
func getFilePath() {
	weixin_consts.Global.PayCertP12Path = g.Cfg().MustGet(context.Background(), "service.payCertP12").String()
	weixin_consts.Global.PayPublicKeyPemPath = g.Cfg().MustGet(context.Background(), "service.payPublicKeyPem").String()
	weixin_consts.Global.PayPrivateKeyPemPath = g.Cfg().MustGet(context.Background(), "service.payPrivateKeyPem").String()
}

// 加载文件内容
func loadFileData(info *weixin_model.SetCertAndKey) *weixin_model.SetCertAndKey {

	//// 硬编码 加载商户号证书密钥文件
	//getFilePath()
	//PayCertP12Data, _ := file.GetFile(weixin_consts.Global.PayCertP12Path)
	//PayPublicKeyPemData, _ := file.GetFile(weixin_consts.Global.PayPublicKeyPemPath)
	//PayPrivateKeyPemData, _ := file.GetFile(weixin_consts.Global.PayPrivateKeyPemPath)

	res := weixin_model.SetCertAndKey{
		ApiV3Key:         info.ApiV3Key,
		ApiV2Key:         info.ApiV2Key,
		PayCertP12:       info.PayCertP12,
		PayPublicKeyPem:  info.PayPublicKeyPem,
		PayPrivateKeyPem: info.PayPrivateKeyPem,
		CertSerialNumber: info.CertSerialNumber,
	}

	if gfile.IsFile(info.PayCertP12) {
		//p12Data, _ := file.GetFile(info.PayCertP12)
		//res.PayCertP12 = string(p12Data)
	}

	if gfile.IsFile(info.PayPublicKeyPem) {
		pubData, _ := file.GetFile(info.PayPublicKeyPem)
		res.PayPublicKeyPem = string(pubData)
	}

	if gfile.IsFile(info.PayPrivateKeyPem) {
		priData, _ := file.GetFile(info.PayPrivateKeyPem)
		res.PayPrivateKeyPem = string(priData)
	}

	fmt.Println(info)

	return &res
}

// SetCertAndKey  设置商户号证书及密钥文件
func (s *sPayMerchant) SetCertAndKey(ctx context.Context, id int64, info *weixin_model.SetCertAndKey) (bool, error) {
	fileData := loadFileData(info)

	// 读取文件
	data := do.WeixinPayMerchant{}
	gconv.Struct(fileData, &data)

	affected, err := daoctl.UpdateWithError(dao.WeixinPayMerchant.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinPayMerchant{Id: id}))

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
	selectInfo, err := s.GetPayMerchantByMchid(ctx, info.Mchid)
	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商户号不存在，请检查", dao.WeixinPayMerchant.Table())
	}

	if len(info.UnionAppid) != 0 {
		if len(selectInfo.UnionAppid) > 0 {
			// 将指定切片复制，并去重
			info.UnionAppid = garray.NewSortedStrArrayFrom(append(info.UnionAppid, selectInfo.UnionAppid...)).Unique().Slice()
		}
	}

	data := do.WeixinPayMerchant{}
	gconv.Struct(info, &data)

	data.UnionAppid, _ = gjson.Encode(info.UnionAppid)

	affected, err := daoctl.UpdateWithError(dao.WeixinPayMerchant.Ctx(ctx).Data(data).OmitNilData().Where(do.WeixinPayMerchant{Mchid: info.Mchid}))

	if err != nil {
		return false, sys_service.SysLogs().ErrorSimple(ctx, err, "商户号基础修改失败", dao.WeixinPayMerchant.Table())
	}
	return affected > 0, err
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
