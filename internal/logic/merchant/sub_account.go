package merchant

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"io/ioutil"
	"log"
)

/*
	分账
*/
/*

单次分账 V2：
	- 单次分账请求按照传入的分账接收方账号和资金进行分账，同时会将订单剩余的待分账金额解冻给特约商户。
      故操作成功后，订单不能再进行分账，也不能进行分账完结。

多次分账 V2：
	- 微信订单支付成功后，服务商代子商户发起分账请求，将结算后的钱分到分账接收方。多次分账请求仅会按照传入的分账接收方进行分账，不会对剩余的金额进行任何操作。
      故操作成功后，在待分账金额不等于零时，订单依旧能够再次进行分账。
	- 多次分账，可以将本商户作为分账接收方直接传入，实现释放资金给本商户的功能

完结分账 V2：
	1、不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给特约商户
	2、调用多次分账接口后，需要解冻剩余资金时，调用本接口将剩余的分账金额全部解冻给特约商户
	3、已调用请求单次分账后，剩余待分账金额为零，不需要再调用此接口。

分账关系查询 V2：


注意：微信支付 V3 版本中取消了单次分账和多次分账的区分，合并为统一的分账接口，即商户可以在同一个订单中同时设置多个分账接收方和其对应的分账比例，

统一分账：
	即商户可以在同一个订单中同时设置多个分账接收方和其对应的分账比例，

添加分账接收方：
	- 相当于绑定了一个分账关系i，后续可通过发起分账请求将结算后的钱分到该分账接收方。

删除分账接收方：
	- 服务商代子商户发起删除分账接收方请求，删除后不支持将结算后的钱分到该分账接收方。

解冻剩余资金API：
	不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给特约商户。
	假如还存在分账关系，需要先进行分账结束后才能解冻

查询分账比例：
	服务商可以查询子商户设置的允许服务商分账的最大比例。

查询分账结果：
	- 发起分账请求后，可调用此接口查询分账结果；发起分账完结请求后，可调用此接口查询分账完结的执行结果。

分账动帐通知：
	 分账或分账回退成功后，微信会把相关变动结果发送给分账接收方（只支持商户）。

查询剩余分账金额：
	查询订单剩余待分金额

*/

type sSubAccount struct {
}

func init() {
	//weixin_service.RegisterSubAccount(NewSubAccount())
}

func NewSubAccount() weixin_service.ISubAccount {
	return &sSubAccount{}
}

func (s *sSubAccount) newClient(ctx context.Context, appId string) (client *core.Client, err error) {
	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}

	// 通过SpMchId拿到微信支付服务商商户号
	spMerchant, err := weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的商户号", "WeiXin-Pay")
	}

	//weixin.NewPayClient(ctx, spMerchantspMchId), spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)

	return weixin.NewPayClient(ctx, gconv.String(spMerchant.Mchid), spMerchant.PayPrivateKeyPem, spMerchant.CertSerialNumber, spMerchant.ApiV3Key)
}

// GetSubAccountMaxRatio 查询最大分账比例
func (s *sSubAccount) GetSubAccountMaxRatio(ctx context.Context, appId string) (*weixin_model.QueryMerchantRatioRes, error) {
	client, _ := s.newClient(ctx, appId)

	// 通过AppId拿到特约商户商户号
	subMerchant, err := weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, "该应用没有对应的特约商户", "WeiXin-Pay")
	}

	svc := profitsharing.MerchantsApiService{Client: client}
	resp, result, err := svc.QueryMerchantRatio(ctx,
		profitsharing.QueryMerchantRatioRequest{
			SubMchid: core.String(gconv.String(subMerchant.SubMchid)),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryMerchantRatio err:%s", err)
		return nil, err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}

	ret := kconv.Struct(resp, &weixin_model.QueryMerchantRatioRes{})

	return ret, err
}

// QuerySubAccountOrder 查询分账结果
func (s *sSubAccount) QuerySubAccountOrder(ctx context.Context, appId string, info *weixin_model.QueryOrderRequest) (*profitsharing.OrdersEntity, error) {
	client, _ := s.newClient(ctx, appId)

	//profitsharing.QueryOrderRequest{
	//	TransactionId: core.String("4208450740201411110007820472"),
	//	OutOrderNo:    core.String("P20150806125346"),
	//	SubMchid:      core.String("1900000109"),
	//},

	req := kconv.Struct(info, &profitsharing.QueryOrderRequest{})
	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.QueryOrder(ctx, *req)

	if err != nil {
		// 处理错误
		log.Printf("call QueryOrder err:%s", err)
		return nil, err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}

	return resp, nil
}

// UnfreezeOrder 解冻剩余资金API
func (s *sSubAccount) UnfreezeOrder(ctx context.Context, appId string, info *weixin_model.UnfreezeOrderRequest) (*profitsharing.OrdersEntity, error) {
	client, _ := s.newClient(ctx, appId)
	/*
		profitsharing.UnfreezeOrderRequest{
				Description:   core.String("解冻全部剩余资金"),
				OutOrderNo:    core.String("P20150806125346"),
				SubMchid:      core.String("1900000109"),
				TransactionId: core.String("4208450740201411110007820472"),
		},
	*/
	req := kconv.Struct(info, &profitsharing.UnfreezeOrderRequest{})

	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.UnfreezeOrder(ctx, *req)

	if err != nil {
		// 处理错误
		log.Printf("call UnfreezeOrder err:%s", err)
		return nil, err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
	return resp, nil
}

// SubAccountRequest 请求分账
func (s *sSubAccount) SubAccountRequest(ctx context.Context, appId string, info *weixin_model.SubAccountReq) (*profitsharing.OrdersEntity, error) {
	client, _ := s.newClient(ctx, appId)

	svc := profitsharing.OrdersApiService{Client: client}
	//req := profitsharing.CreateOrderRequest{
	//	Appid:      core.String(info.Appid),        // 服务商AppID
	//	OutOrderNo: core.String(info.OutOrderNo), // 订单号orderId = 商户订单号out_trade_no = 分账请求号OutOrderNo
	//
	//	Receivers: []profitsharing.CreateOrderReceiver{
	//		profitsharing.CreateOrderReceiver{
	//			Account:     core.String(info.Receivers[0].Account),
	//			Amount:      core.Int64(888),
	//			Description: core.String(info."分给商户A"),
	//			Name:        core.String(info."hu89ohu89ohu89o"),
	//			Type:        core.String(info."MERCHANT_ID"),
	//		}},
	//	SubAppid:        core.String(info."wx8888888888888889"),
	//	SubMchid:        core.String(info."1900000109"),
	//	TransactionId:   core.String(info."4208450740201411110007820472"),
	//	UnfreezeUnsplit: core.Bool(true),
	//}

	req := profitsharing.CreateOrderRequest{}

	_ = gconv.Struct(info, &req)

	// 分账下单
	resp, result, err := svc.CreateOrder(ctx, req)

	if err != nil {
		// 处理错误
		log.Printf("call CreateOrder err:%s", err)
		return nil, err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}

	return resp, nil
}

// QueryOrderAmount 查询剩余待分金额API
func (s *sSubAccount) QueryOrderAmount(ctx context.Context, appId string, info *weixin_model.QueryOrderAmountRequest) (*profitsharing.QueryOrderAmountResponse, error) {
	//appId := weixin_utility.GetAppIdFormContext(ctx) // 特约商户绑定的AppId
	client, _ := s.newClient(ctx, appId)

	req := profitsharing.QueryOrderAmountRequest{}

	_ = gconv.Struct(info, &req)

	svc := profitsharing.TransactionsApiService{Client: client}

	resp, result, err := svc.QueryOrderAmount(ctx, req)

	if err != nil {
		// 处理错误
		log.Printf("call QueryOrderAmount err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}

	return resp, nil
}

// AddReceiver 添加分账接收方（相当于绑定分账关系）
func (s *sSubAccount) AddReceiver(ctx context.Context, appId string, info *weixin_model.AddReceiverRequest) (*profitsharing.AddReceiverResponse, error) {
	client, _ := s.newClient(ctx, appId)

	svc := profitsharing.ReceiversApiService{Client: client}
	//resp, result, err := svc.AddReceiver(ctx,
	//	profitsharing.AddReceiverRequest{
	//		Account:        core.String("86693852"),
	//		Appid:          core.String("wx8888888888888888"),
	//		CustomRelation: core.String("代理商"),
	//		Name:           core.String("hu89ohu89ohu89o"),
	//		RelationType:   profitsharing.RECEIVERRELATIONTYPE_SERVICE_PROVIDER.Ptr(),
	//		SubAppid:       core.String("wx8888888888888889"),
	//		SubMchid:       core.String("1900000109"),
	//		Type:           profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
	//	},
	//)

	req := kconv.Struct(info, &profitsharing.AddReceiverRequest{})
	resp, result, err := svc.AddReceiver(ctx, *req)

	if err != nil {
		// 处理错误
		log.Printf("call AddReceiver err:%s", err)
		return nil, err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}

	return resp, nil

}

// AddProfitSharingReceivers 添加多个分账关系
func (s *sSubAccount) AddProfitSharingReceivers(ctx context.Context, appId string, info []weixin_model.AddReceiverRequest) (*profitsharing.AddReceiverResponse, error) {

	receivers := make([]weixin_model.AddReceiverRequest, 0)
	receivers = append(receivers, info...)

	reqBody := g.Map{
		"receivers": receivers,
	}

	client, _ := s.newClient(ctx, appId)

	result, err := client.Post(ctx, consts.WechatPayAPIServer+"/v3/profitsharing/receivers/add", reqBody)

	if err != nil {
		// 处理错误
		log.Printf("call AddReceiver err:%s", err)
		return nil, err
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, result)
	}

	// 处理成功响应结果方法1
	body, err := ioutil.ReadAll(result.Response.Body)

	res := profitsharing.AddReceiverResponse{}
	_ = gjson.DecodeTo(body, res)

	// 处理成功响应结果方法2
	//resp := new(profitsharing.AddReceiverResponse)
	//err = core.UnMarshalResponse(result.Response, resp)
	//if err != nil {
	//	return nil, err
	//}

	return &res, nil
}

// DeleteReceiver 删除分账接收方（相当于分账关系解绑）
func (s *sSubAccount) DeleteReceiver(ctx context.Context, appId string, info *weixin_model.DeleteReceiverRequest) (*profitsharing.DeleteReceiverResponse, error) {
	client, _ := s.newClient(ctx, appId)

	svc := profitsharing.ReceiversApiService{Client: client}
	//profitsharing.DeleteReceiverRequest{
	//	Account:  core.String("86693852"),
	//	Appid:    core.String("wx8888888888888888"),
	//	SubAppid: core.String("wx8888888888888889"),
	//	SubMchid: core.String("1900000109"),
	//	Type:     profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
	//},

	req := kconv.Struct(info, profitsharing.DeleteReceiverRequest{})

	resp, result, err := svc.DeleteReceiver(ctx, req)

	if err != nil {
		// 处理错误
		log.Printf("call DeleteReceiver err:%s", err)
		return nil, err

	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}

	return resp, nil
}

// 分账动帐结果通知 (1.调用分账下单的时候回返回结果 2.同步分账不会有通知，异步分账通知Post到回调url中 )
