package weixin_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	v1 "github.com/kysion/kys-weixin-library/api/weixin_v1"
	"github.com/kysion/kys-weixin-library/weixin_service"
)

// WeiXin 网关
var WeiXin = cWeiXin{}

type cWeiXin struct{}

type StringRes string

// WeiXinServices 商家授权应用，等消息推送，消息通知，通过这个消息  针对B端
func (c *cWeiXin) WeiXinServices(ctx context.Context, req *v1.WeiXinServicesReq) (v1.StringRes, error) {
	result, err := weixin_service.WeiXin().WeiXinServices(ctx)
	return (v1.StringRes)(result), err
}

// WeiXinCallback C端业务小消息   消费者支付.....
func (c *cWeiXin) WeiXinCallback(ctx context.Context, req *v1.WeiXinCallbackReq) (api_v1.BoolRes, error) {
	result, err := weixin_service.WeiXin().WeiXinCallback(ctx)

	return result != "", err
}

func (c *cWeiXin) CheckSignature(ctx context.Context, req *v1.CheckSignatureReq) (v1.StringRes, error) {
	// 时间戳，单位秒
	// 时间戳，单位纳秒 UnixNano
	//unix := time.Now().Unix()

	return (v1.StringRes)(weixin_service.WeiXin().WXCheckSignature(ctx, req.Signature, req.Signature, req.Nonce, req.Echostr)), nil
}

// AlipayAuthUserInfo 用户登录信息

// 获取用户openId

// 获取用户unionID

// 用户网页授权

// 获取用户基本信息
