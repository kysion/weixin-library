package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/pay-share-library/pay_model/pay_enum"
	"github.com/kysion/pay-share-library/pay_service"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
	"strconv"
	"time"
)

type sWeiXinPay struct {
}

//func init() {
//	weixin_service.RegisterWeiXinPay(NewWeiXinPay())
//}

func NewWeiXinPay() weixin_service.IWeiXinPay {

	result := &sWeiXinPay{}

	//result.injectHook()
	return result
}

// PayTradeCreate  1、创建交易订单   （AppId的H5是没有的，需要写死，小程序有的 ）
func (s *sWeiXinPay) PayTradeCreate(ctx context.Context, info *weixin_model.TradeOrder, openId string) (*weixin_model.PayParamsRes, error) {
	_ = sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------WeiXin创建交易订单 ------- ", "WeiXin-Pay")
	appId := weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId

	// 商家AppId解析，获取商家应用，创建微信支付客户端
	merchantApp, err := weixin_service.MerchantAppConfig().GetMerchantAppConfigByAppId(ctx, appId)
	if err != nil {
		return nil, err
	}

	sysUser, err := sys_service.SysUser().GetSysUserById(ctx, merchantApp.SysUserId)
	if err != nil {
		return nil, err
	}

	info.Order.TradeSourceType = pay_enum.Order.TradeSourceType.Weixin.Code() // 交易源类型
	info.Order.UnionMainId = merchantApp.UnionMainId
	info.Order.UnionMainType = sysUser.Type

	// 支付前创建交易订单，支付后修改交易订单元数据
	orderInfo, err := pay_service.Order().CreateOrder(ctx, &info.Order) // CreatedOrder不能修改订单id
	if err != nil || orderInfo == nil {
		return nil, err
	}

	// 判断是否是第三方待开发
	isParent := false
	if merchantApp.ThirdAppId != "" {
		isParent = true
	}

	var prepayId string

	switch merchantApp.AppType {
	case weixin_enum.AppManage.AppType.PublicAccount.Code(): // 公众号
		if isParent {
			//  公众号  -- 服务商模式
			prepayId, err = s.JsapiCreateOrder(ctx, &weixin_model.TradeOrder{
				ReturnUrl: info.ReturnUrl, // 支付成功后的返回地址
				Order:     info.Order,
			}, openId)

		} else {
			//  公众号  -- 直连模式
			prepayId, err = s.JsapiCreateOrderByDirect(ctx, &weixin_model.TradeOrder{
				ReturnUrl: info.ReturnUrl, // 支付成功后的返回地址
				Order:     info.Order,
			}, openId)
		}

	case weixin_enum.AppManage.AppType.TinyApp.Code(): // 小程序
		// 小程序  JsApi支付产品
		prepayId, err = s.JsapiCreateOrder(ctx, &weixin_model.TradeOrder{
			ReturnUrl: info.ReturnUrl, // 支付成功后的返回地址
			Order:     info.Order,
		}, openId)

	case weixin_enum.AppManage.AppType.H5.Code(): // H5

	case weixin_enum.AppManage.AppType.App.Code(): // APP

	}

	// 支付订单创建成功后，需要拼接好支付参数，然后返回给前端
	return s.makePayParams(ctx, gconv.String(orderInfo.Id), appId, prepayId)
}

// 生成支付所需参数
func (s *sWeiXinPay) makePayParams(ctx context.Context, orderId, appId, prepay_id string) (*weixin_model.PayParamsRes, error) {
	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	//if err != nil {
	//	return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	//}
	// 通过AppId拿到特约商户商户号
	//var spMerchant *weixin_model.PayMerchant
	var payPrivateKeyPem string

	spMerchant, _ := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if spMerchant != nil && spMerchant.PayPrivateKeyPem != "" {
		payPrivateKeyPem = spMerchant.PayPrivateKeyPem
	}

	if payPrivateKeyPem == "" {
		merchant, _ := weixin_service.PayMerchant().GetPayMerchantByAppId(ctx, appId)
		if merchant != nil && merchant.PayPrivateKeyPem != "" {
			payPrivateKeyPem = merchant.PayPrivateKeyPem
		}
	}

	ret := &weixin_model.PayParamsRes{
		AppId:     appId,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  weixin_utility.Md5Hash(orderId),
		Package:   "prepay_id=" + prepay_id,
		SignType:  "RSA",
		PaySign:   "",
	}

	// 1.构建签名串
	/*
		应用ID
		时间戳
		随机字符串
		订单详情扩展字符串
	*/
	var content = ret.AppId + "\n" +
		ret.TimeStamp + "\n" +
		ret.NonceStr + "\n" +
		ret.Package + "\n"

	// 2.计算签名值paySign = appId、timeStamp、nonceStr、package ==》 通过私钥进行SHA256 with RSA签名 ==》 对签名结果进行Base64编码得到签名值
	privateKey, err := weixin.LoadPrivateKey(payPrivateKeyPem)

	var sign, _ = weixin.SignSHA256WithRSA(content, privateKey) // qDLKva8l1HPQ0GDjQA9cHMqIg8cI4JWv0/toKBoA+8dSgIKKySQniAv8AKapAj3DHX1Td6xS9Tgm2LPUewdP4KkZ6aYOdbtiDLaoCiuLNud4S0mTsek7Re9oOaA5OCIqsz2E5AYOWJkGxebrIOhWAWChKiT/+JKZXWdBozuYIN0tqtirfK3xuhaPszlx0sJwD0V7Gn2tYK9VVVVYfpFNdXZeQaehdpDVfj5xkVXaH8yQwweoljoy1qWC+UFmZ+/8TIu5w3OslMnbrWIlMOckJdfnv5bXyvkChzETfO4R46eiOdkXi1dP6759S9FZn7JVFglu22aJdTVk3g7e8BmtHA==
	ret.PaySign = sign

	// 3.将支付参数返回至前端

	return ret, err
	/*
		wx6cc2c80416074df3
		1684393307
		f9aa40a057cae16c37a2b97db23a86ed
		prepay_id=wx18150015642076d683b4336866f9370000
	*/
}
