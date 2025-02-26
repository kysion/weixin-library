package gateway

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/pay-share-library/pay_model"
	"github.com/kysion/pay-share-library/pay_model/pay_enum"
	"github.com/kysion/pay-share-library/pay_service"
	"github.com/kysion/weixin-library/internal/logic/internal/weixin"
	"github.com/kysion/weixin-library/weixin_consts"
	"github.com/kysion/weixin-library/weixin_model"
	dao "github.com/kysion/weixin-library/weixin_model/weixin_dao"
	hook "github.com/kysion/weixin-library/weixin_model/weixin_hook"
	"github.com/kysion/weixin-library/weixin_service"
	"github.com/kysion/weixin-library/weixin_utility"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
	"time"
)

/*
	异步通知地址 (接收支付结果通知接口)
*/

type sMerchantNotify struct {
	// 异步通知Hook （针对于关注通知的业务需求）
	NotifyHook base_hook.BaseHook[hook.NotifyKey, hook.NotifyHookFunc]

	// 交易Hook 	（针对于关注支付结果的业务需求）
	TradeHook base_hook.BaseHook[hook.TradeHookKey, hook.TradeHookFunc]

	// 分账Hook (暂时没用到)
	// SubAccountHook base_hook.BaseHook[hook.SubAccountHookKey, hook.SubAccountHookFunc]
}

//func init() {
//	weixin_service.RegisterMerchantNotify(NewMerchantNotify())
//}

func NewMerchantNotify() weixin_service.IMerchantNotify {
	return &sMerchantNotify{}
}

/*
1. 收到通知报文

2. 通知验证签名，确保消息来自微信

3. 参数解密，得到json数据

4. 订阅Hook，处理账单、处理分账

5. 根据通知结果，返回通知应答
   1. 接收成功：HTTP应答状态码需返回200或204，无需返回应答报文。

   2. 接收失败：HTTP应答状态码需返回5XX或4XX，同时需返回应答报文，格式如下：

      ```
      {
          "code": "FAIL",
          "message": "失败"
      }
      ```
*/

// InstallNotifyHook 订阅异步通知Hook
func (s *sMerchantNotify) InstallNotifyHook(hookKey hook.NotifyKey, hookFunc hook.NotifyHookFunc) {
	_ = sys_service.SysLogs().InfoSimple(context.Background(), nil, "\n-------订阅订阅异步通知Hook： ------- ", "sPlatformUser")

	hookKey.HookCreatedAt = *gtime.Now()

	tradeHookExpireAt := weixin_consts.Global.TradeHookExpireAt
	if tradeHookExpireAt == 0 {
		tradeHookExpireAt = 7200
	}

	secondAt := gtime.New(tradeHookExpireAt * gconv.Int64(time.Second))
	hookKey.HookExpireAt = *gtime.New(hookKey.HookCreatedAt.Second() + secondAt.Second())

	s.NotifyHook.InstallHook(hookKey, hookFunc)
}

// InstallTradeHook 订阅支付Hook
func (s *sMerchantNotify) InstallTradeHook(hookKey hook.TradeHookKey, hookFunc hook.TradeHookFunc) {
	hookKey.HookCreatedAt = *gtime.Now()

	tradeHookExpireAt := weixin_consts.Global.TradeHookExpireAt
	if tradeHookExpireAt == 0 {
		tradeHookExpireAt = 7200
	}

	secondAt := gtime.New(tradeHookExpireAt * gconv.Int64(time.Second))

	hookKey.HookExpireAt = *gtime.New(hookKey.HookCreatedAt.Second() + secondAt.Second())

	s.TradeHook.InstallHook(hookKey, hookFunc)
}

// NotifyServices 异步通知地址  用于接收支付宝推送给商户的支付/退款成功的消息。
func (s *sMerchantNotify) NotifyServices(ctx context.Context) (string, error) {
	_ = sys_service.SysLogs().InfoSimple(ctx, nil, "\n----------微信支付异步通知", "WeiXin-Notify")

	appId := weixin_utility.GetAppIdFormContext(ctx)

	var err error
	var spMerchant *weixin_model.PayMerchant           // 服务商
	var subMerchant *weixin_model.WeixinPaySubMerchant // 特约商户

	subMerchant, err = weixin_service.PaySubMerchant().GetPaySubMerchantByAppId(ctx, appId)
	//if err != nil {
	//	return "", err
	//}

	if subMerchant != nil && subMerchant.SpMchid != 0 {
		spMerchant, err = weixin_service.PayMerchant().GetPayMerchantByMchid(ctx, subMerchant.SpMchid)
	} else {
		spMerchant, err = weixin_service.PayMerchant().GetPayMerchantByAppId(ctx, appId)
	}

	if err != nil {
		return "", err
	}

	err = dao.WeixinConsumerConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 1.验签并解密，得到通知json数据
		transaction, err := ParseNotifyRequestTransaction(ctx, spMerchant)
		if err != nil {
			return sys_service.SysLogs().ErrorSimple(ctx, err, "验签失败", "WeiXin-Notify")
		}

		{
			// a. 将交易元数据存储至 kmk_order
			transactionJson, _ := gjson.Encode(transaction)
			outTradeNo := gconv.Int64(transaction.OutTradeNo)  // 商户订单号，是我们自己指定的，OutTradeNo = orderId
			tradeId := gconv.String(transaction.TransactionId) // 微信交易凭证id。
			tradeJson := gconv.String(transactionJson)         // 交易元数据
			info := pay_model.UpdateOrderTradeInfo{
				PlatformOrderId: &tradeId,   // 微信交易凭证id。
				TradeSource:     &tradeJson, // 交易元数据
			}
			_, err = pay_service.Order().UpdateOrderTradeSource(ctx, outTradeNo, &info)
			if err != nil {
				return err
			}
		}

		{
			// 2.判断交易状态，然后修改对应的订单状态
			var orderState int
			switch *transaction.TradeState {
			case pay_enum.WeiXinTrade.TradeStatus.SUCCESS.Code():
				// 成功 --> 订单状态为已支付
				orderState = pay_enum.Order.StateType.Paymented.Code()
			case pay_enum.WeiXinTrade.TradeStatus.REFUND.Code():
				// 转入退款 --> 订单状态为退款中
				orderState = pay_enum.Order.StateType.Refunding.Code()

			case pay_enum.WeiXinTrade.TradeStatus.NOTPAY.Code():
				// 未支付 --> 订单状态为待支付
				orderState = pay_enum.Order.StateType.WaitPayment.Code()

			case pay_enum.WeiXinTrade.TradeStatus.CLOSED.Code():
				// 已关闭 --> 订单状态为已关闭
				orderState = pay_enum.Order.StateType.ClosedPayment.Code()

			case pay_enum.WeiXinTrade.TradeStatus.REVOKED.Code():
				// 已撤销 --> 订单状态为取消支付（仅付款码支付会返回）
				orderState = pay_enum.Order.StateType.CancelPayment.Code()

			case pay_enum.WeiXinTrade.TradeStatus.USERPAYING.Code():
				// 用户支付中 --> 订单状态为支付中（仅付款码支付会返回）
				orderState = pay_enum.Order.StateType.Paymenting.Code()

			case pay_enum.WeiXinTrade.TradeStatus.PAYERROR.Code():
				// 交易失败 --> 订单状态为已关闭（仅付款码支付会返回）
				orderState = pay_enum.Order.StateType.ClosedPayment.Code()

			}

			_, err := pay_service.Order().UpdateOrderState(ctx, gconv.Int64(transaction.OutTradeNo), orderState)
			if err != nil {
				return err
			}
		}

		orderInfo, err := pay_service.Order().GetOrderById(ctx, gconv.Int64(transaction.OutTradeNo))
		if err != nil {
			return err
		}

		// 3. 支付成功，添加账单account_bill  商家 消费者的账单  及广播业务层Hook
		if transaction.TradeState != nil && *transaction.TradeState == pay_enum.WeiXinTrade.TradeStatus.SUCCESS.Code() {
			// 4. 分账交易下单结算  需要支付状态为Success的订单

			// a.查询分账关系
			//relationBatch, _ := service.SubAccount().TradeRelationBatchQuery(ctx, gconv.String(bm["auth_app_id"]), gconv.String(bm["out_trade_no"]))
			//if relationBatch.AlipayTradeRoyaltyRelationBatchqueryResponse.ResultCode == enum.SubAccount.SubAccountBindRes.Fail.Code() {
			//	return nil
			//}

			// b.找到分账支出方账户  可选

			// c.组装分账明细信息 + 分账拓展参数

			// 2.次日分账，添加分账的定时任务

			// e.分账通知会发送到应用网关，然后我们判断分账结果，从而创建分账快照
			// alipay.trade.order.settle.notify(交易分账结果通知)  这是我们自己定义的接口吗，不，是应用网关

			// f.上面这些全部写到了具体业务端的定时任务器中操作，先查询所有未分账的账单，然后进行分账，然后请求未分账标记信息

			isClean := false

			// Trade发布广播
			s.TradeHook.Iterator(func(key hook.TradeHookKey, value hook.TradeHookFunc) {
				if key.WeiXinTradeStatus.Code() == pay_enum.WeiXinTrade.TradeStatus.SUCCESS.Code() {
					_ = sys_service.SysLogs().InfoSimple(ctx, nil, "\n-------微信支付异步通知TradeHook发布广播-------- ", "WeiXin-Notify")

					value(ctx, orderInfo)
				}

				s.TradeHook.UnInstallHook(key, func(filter hook.TradeHookKey, conditionKey hook.TradeHookKey) bool {
					// 如果超时了，那么就返回true，代表可以删除
					if key.HookExpireAt.Before(gtime.Now()) && key.TradeNo != "" {
						// 底层的filter和conditionKey是一样的
						return filter == conditionKey
					}
					// 没超时，但是业务层指定了isCLean为true，那么也执行删除
					return isClean && filter == conditionKey
				})
			})

			// TODO 4.远程设置设备通电  此异步任务需要写到实际的业务层，而不是这里 （当初用于给筷满客的设备开机）
			//go func() {
			//	url := "http://10.168.173.252:7771/device/setPowerState?serialNumber=" + orderInfo.ProductNumber + "&isPowerOn=true"
			//	g.Client().PostContent(ctx, url)
			//}()
		}
		return nil
	})

	// 根据支付状态，返回通知响应给微信支付
	if err != nil {
		return "success", err
	}

	//1. 接收成功：HTTP应答状态码需返回200或204，无需返回应答报文。
	//
	//2. 接收失败：HTTP应答状态码需返回5XX或4XX，同时需返回应答报文，格式如下：
	//
	//```
	//  {
	//      "code": "FAIL",
	//      "message": "失败"
	//  }

	respBody := g.Map{
		"code":    "SUCCESS",
		"message": "OK",
	}

	g.RequestFromCtx(ctx).Response.Write(respBody)

	return "success", nil
}

func new2NotifyHandler(ctx context.Context, spMerchant *weixin_model.PayMerchant) (*core.Client, *notify.Handler) {
	pri, _ := weixin.LoadPrivateKey(spMerchant.PayPrivateKeyPem)
	mchId := gconv.String(spMerchant.Mchid)

	// 1. 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchId, spMerchant.CertSerialNumber, pri, spMerchant.ApiV3Key),
	}

	client, _ := core.NewClient(ctx, opts...)

	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mchId)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(spMerchant.ApiV3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	// 4. 使用client进行接口调用
	// ...
	return client, handler
}

// ParseNotifyRequestTransaction 解析异步通知内容到结构体里面
func ParseNotifyRequestTransaction(ctx context.Context, spMerchant *weixin_model.PayMerchant) (*partnerpayments.Transaction, error) {
	// 初始化 NotifyHandler
	handler := weixin.NewNotifyHandler(ctx, spMerchant)
	request := g.RequestFromCtx(ctx).Request
	//var handler notify.Handler

	content := new(partnerpayments.Transaction)
	// 验签并解密报文
	notifyReq, err := handler.ParseNotifyRequest(ctx, request, content) //unsupported Wechatpay-Signature-Type: WECHATPAY2-SHA256-RSA2048
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 处理通知内容
	fmt.Println(notifyReq.Summary)

	_ = sys_service.SysLogs().InfoSimple(ctx, nil, fmt.Sprintf("支付的通知内容为：%s", content), "WeiXin-Notify")

	return content, nil
}
