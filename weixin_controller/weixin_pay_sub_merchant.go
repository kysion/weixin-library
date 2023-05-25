package weixin_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/kysion/weixin-library/api/weixin_v1/weixin_pay_sub_merchant_v1"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
)

var WeiXinPaySubMerchant = cWeiXinPaySubMerchant{}

type cWeiXinPaySubMerchant struct{}

// GetPaySubMerchantById 根据id查找特约商户配置信息
func (c *cWeiXinPaySubMerchant) GetPaySubMerchantById(ctx context.Context, req *weixin_pay_sub_merchant_v1.GetPaySubMerchantByIdReq) (*weixin_model.PaySubMerchantRes, error) {
	ret, err := weixin_service.PaySubMerchant().GetPaySubMerchantById(ctx, req.Id)

	return (*weixin_model.PaySubMerchantRes)(ret), err
}

// GetPaySubMerchantByMchid 根据Mchid查找特约商户配置信息
func (c *cWeiXinPaySubMerchant) GetPaySubMerchantByMchid(ctx context.Context, req *weixin_pay_sub_merchant_v1.GetPaySubMerchantByMchidReq) (*weixin_model.PaySubMerchantRes, error) {
	ret, err := weixin_service.PaySubMerchant().GetPaySubMerchantByMchid(ctx, req.MchId)

	return (*weixin_model.PaySubMerchantRes)(ret), err
}

// GetPaySubMerchantBySysUserId  根据用户id查询特约商户配置信息
func (c *cWeiXinPaySubMerchant) GetPaySubMerchantBySysUserId(ctx context.Context, req *weixin_pay_sub_merchant_v1.GetPaySubMerchantBySysUserIdReq) (*weixin_model.PaySubMerchantRes, error) {
	ret, err := weixin_service.PaySubMerchant().GetPaySubMerchantBySysUserId(ctx, req.SysUserId)

	return (*weixin_model.PaySubMerchantRes)(ret), err
}

// CreatePaySubMerchant  创建特约商户配置信息
func (c *cWeiXinPaySubMerchant) CreatePaySubMerchant(ctx context.Context, req *weixin_pay_sub_merchant_v1.CreatePaySubMerchantReq) (*weixin_model.PaySubMerchantRes, error) {
	ret, err := weixin_service.PaySubMerchant().CreatePaySubMerchant(ctx, &req.WeixinPaySubMerchant)

	return (*weixin_model.PaySubMerchantRes)(ret), err
}

// UpdatePaySubMerchant 更新特约商户配置信息
func (c *cWeiXinPaySubMerchant) UpdatePaySubMerchant(ctx context.Context, req *weixin_pay_sub_merchant_v1.UpdatePaySubMerchantReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.PaySubMerchant().UpdatePaySubMerchant(ctx, req.Id, &req.UpdatePaySubMerchant)

	return ret == true, err
}

// SetSubMerchantAuthPath 设置特约商户授权目录
func (c *cWeiXinPaySubMerchant) SetSubMerchantAuthPath(ctx context.Context, req *weixin_pay_sub_merchant_v1.SetSubMerchantAuthPathReq) (api_v1.BoolRes, error) {
	ret, err := weixin_service.PaySubMerchant().SetAuthPath(ctx, &req.SetSubMerchantAuthPath)

	return ret == true, err
}
