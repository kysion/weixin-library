package weixin_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_pay_merchant_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

var WeiXinPayMerchant = cWeiXinPayMerchant{}

type cWeiXinPayMerchant struct{}

// GetPayMerchantById 根据id查找商户号配置信息
func (c *cWeiXinPayMerchant) GetPayMerchantById(ctx context.Context, req *weixin_pay_merchant_v1.GetPayMerchantByIdReq) (*weixin_model.PayMerchantRes, error) {
	ret, err := weixin_service.PayMerchant().GetPayMerchantById(ctx, req.Id)

	return (*weixin_model.PayMerchantRes)(ret), err
}

// GetPayMerchantByMchid 根据Mchid查找商户号配置信息
func (c *cWeiXinPayMerchant) GetPayMerchantByMchid(ctx context.Context, req *weixin_pay_merchant_v1.GetPayMerchantByMchidReq) (*weixin_model.PayMerchantRes, error) {
	ret, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, req.MchId)

	return (*weixin_model.PayMerchantRes)(ret), err
}

// GetPayMerchantBySysUserId  根据商家id查询商户号配置信息
func (c *cWeiXinPayMerchant) GetPayMerchantBySysUserId(ctx context.Context, req *weixin_pay_merchant_v1.GetPayMerchantBySysUserIdReq) (*weixin_model.PayMerchantRes, error) {
	ret, err := weixin_service.PayMerchant().GetPayMerchantBySysUserId(ctx, req.SysUserId)

	return (*weixin_model.PayMerchantRes)(ret), err
}

// CreatePayMerchant  创建商户号配置信息
func (c *cWeiXinPayMerchant) CreatePayMerchant(ctx context.Context, req *weixin_pay_merchant_v1.CreatePayMerchantReq) (*weixin_model.PayMerchantRes, error) {
	ret, err := weixin_service.PayMerchant().CreatePayMerchant(ctx, &req.PayMerchant)

	return (*weixin_model.PayMerchantRes)(ret), err
}

// UpdatePayMerchant 更新商户号配置信息
func (c *cWeiXinPayMerchant) UpdatePayMerchant(ctx context.Context, req *weixin_pay_merchant_v1.UpdatePayMerchantReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.PayMerchant().UpdatePayMerchant(ctx, req.Id, &req.UpdatePayMerchant)

	return ret == true, err
}

// SetCertAndKey  设置商户号证书及密钥文件
func (c *cWeiXinPayMerchant) SetCertAndKey(ctx context.Context, req *weixin_pay_merchant_v1.SetCertAndKeyReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.PayMerchant().SetCertAndKey(ctx, req.Id, &req.SetCertAndKey)

	return ret == true, err
}

// SetAuthPath 设置商户号授权目录
func (c *cWeiXinPayMerchant) SetAuthPath(ctx context.Context, req *weixin_pay_merchant_v1.SetAuthPathReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.PayMerchant().SetAuthPath(ctx, &req.SetAuthPath)

	return ret == true, err
}

// SetPayMerchantUnionId 设置商户号关联的AppId
func (c *cWeiXinPayMerchant) SetPayMerchantUnionId(ctx context.Context, req *weixin_pay_merchant_v1.SetPayMerchantUnionIdReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.PayMerchant().SetPayMerchantUnionId(ctx, &req.SetPayMerchantUnionId)

	return ret == true, err
}

// SetBankcardAccount 设置商户号银行卡号
func (c *cWeiXinPayMerchant) SetBankcardAccount(ctx context.Context, req *weixin_pay_merchant_v1.SetBankcardAccountReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.PayMerchant().SetBankcardAccount(ctx, &req.SetBankcardAccount)

	return ret == true, err
}
