package weixin_controller

import (
	"context"
	"fmt"
	v1 "github.com/kysion/weixin-library/api/weixin_v1"
	"github.com/kysion/weixin-library/weixin_service"
)

// WeiXin 网关
var WeiXin = cWeiXin{}

type cWeiXin struct{}

type StringRes string

// WeiXinServices 商家授权应用，等消息推送，消息通知，通过这个消息  针对B端
func (c *cWeiXin) WeiXinServices(ctx context.Context, req *v1.ServicesReq) (v1.StringRes, error) {
	fmt.Println("推送消息：", req.MessageEncryptReq)

	/*
		{
			751453258
			j3J0XblMIC+GNHixG1NZmu8PoPUTugrrhwCO3cx+SRj40FrEq4QuxYb8hEt2XvU+mukvgHoXgYnoAibRdxVCG/fkXHD7S3OmoybKwmdUoH232XgLdjxRXJExpn4ScmKyRpQULLqLIjViY9IJFNCs5z1QWTke3W4A4BK6qOG1ppoSH+34wgw3LI/ILuQhcHEZKnic90C5iDY2VYePS7QG3V9G3Zl/z8j4Opwe54fThs0gt3PCQOQxChoDZosw7WeL/taFHJB//+LiHbvnthdwtDM4IiaMZ04rTqJlnFZXqn++3Gk0xas16iHu721bui6+M5R1S8k98nHm6brTzdiHXRv7rUF1UPodjr9Vm1EOD8fEFrMUykqz9b/P04y9mb8KogGfbaC1h5BSyLfM+isfoXilSQ87pOAiJMTXuFrXd1rWwZUGD6t7YnaFo9pRFcYRHUqDEs86fzmWIwWk0HA==
			976058fa2a5d14b4eb2c8f81ea7fd2b7eea60d68
			1677854810
		}
	*/

	/*
		wx534d1a08aa84c529
		tF7U9rjAzZQ5wJpBRmHjMndBHOyjOwu+70mty1IUStw5opir+5ShBdQJWi048GEwoEqbplaw+w7xS4a7xotTTJQJa29+0yiKsSb8HURhMT4HsFVkTIBC53xN10R5iE/uxnrJ57FCaN1en7VTAWjrwpjJ/p604Pmfcq7lV7bgd5jOsLyYLSUlPqL7m6VpY+RbNeg3VT22zSQJAeCvuyjvO9mgp9FBx59mB3mK9qD/ItAB0RxxbPBYmQNEQAwThmWEyhAeVRpGyEErEvA43vuLNrmC5MeDu+bko8/1GnY1B26OYT8JyD5DPBCawFf8ktn12HbYPL0lYde/p1iUYCln5Axod2Hwo91nIyFbINkOWXuFieF2J4wnOxAFIZ6v7h+nd5a2nvi+zxIkyKdKfYT9FQ6Ke6R/UXGZ/kC1oUP+oHh3U/h3QUwfQYNhPWwzqXXTfUGhhi2Oqt9jGBwL0Pw==
	*/
	_, err := weixin_service.Gateway().Services(ctx, &req.EventEncryptMsgReq, &req.MessageEncryptReq)

	return "success", err
}

// WeiXinCallback C端业务小消息   消费者支付.....
func (c *cWeiXin) WeiXinCallback(ctx context.Context, req *v1.CallbackReq) (v1.StringRes, error) {

	_, err := weixin_service.Gateway().Callback(ctx, &req.AuthorizationCodeRes, &req.EventEncryptMsgReq, &req.MessageEncryptReq)

	return "success", err
}

// WeiXinCallbackPost C端业务小消息   消费者支付.....
func (c *cWeiXin) WeiXinCallbackPost(ctx context.Context, req *v1.CallbackPostReq) (v1.StringRes, error) {

	_, err := weixin_service.Gateway().Callback(ctx, &req.AuthorizationCodeRes, &req.EventEncryptMsgReq, &req.MessageEncryptReq)

	return "success", err
}

func (c *cWeiXin) CheckSignature(ctx context.Context, req *v1.CheckSignatureReq) (v1.StringRes, error) {
	// 时间戳，单位秒
	// 时间戳，单位纳秒 UnixNano
	//unix := time.Now().Unix()

	return (v1.StringRes)(weixin_service.Gateway().WXCheckSignature(ctx, req.Signature, req.Signature, req.Nonce, req.Echostr)), nil
}
