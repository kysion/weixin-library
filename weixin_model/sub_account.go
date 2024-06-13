package weixin_model

import "github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"

// 查询分账比例--------------------------------------------------------------------------------------------------------

// QueryMerchantRatioRes 最大分账比例
type QueryMerchantRatioRes struct {
	// 子商户允许父商户分账的最大比例，单位万分比，比如2000表示20%
	MaxRatio *int64 `json:"max_ratio"`

	// 参考请求参数
	SubMchid *string `json:"sub_mchid"`
}

// 请求分账--------------------------------------------------------------------------------------------------------

type SubAccountReq struct {
	Appid string `json:"appid" dc:"服务商appId"`

	// 服务商系统内部的分账单号，在服务商系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo string `json:"out_order_no" dc:"分账单号，可自定义，保证在系统内部唯一"`

	// 分账接收方列表，可以设置出资商户作为分账接受方，最多可有50个分账接收方
	Receivers []CreateOrderReceiver `json:"receivers,omitempty"`

	// 微信分配的子商户公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要）
	SubAppid string `json:"sub_appid,omitempty" dc:"特约商户AppID"`

	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid string `json:"sub_mchid,omitempty" dc:"特约商户号"`

	// 微信支付订单号
	TransactionId string `json:"transaction_id" dc:"微信支付后返回的交易单号"`

	// 1、如果为true，该笔订单剩余未分账的金额会解冻回分账方商户； 2、如果为false，该笔订单剩余未分账的金额不会解冻回分账方商户，可以对该笔订单再次进行分账。
	UnfreezeUnsplit bool `json:"unfreeze_unsplit" dc:"是否资金解冻，true的话分账结束后资金解冻回分账方商户，false不解冻，可再次进行分账"`
}

// CreateOrderReceiver 分账接收方
type CreateOrderReceiver struct {
	// 1、类型是MERCHANT_ID时，是商户号 2、类型是PERSONAL_OPENID时，是个人openid  3、类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Account string `json:"account" dc:"账户"`

	// 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	Amount int64 `json:"amount" dc:"分账金额"`

	// 分账的原因描述，分账账单中需要体现
	Description string `json:"description" dc:"分账的原因描述，分账账单中需要体现"`

	// 可选项，在接收方类型为个人的时可选填，若有值，会检查与 name 是否实名匹配，不匹配会拒绝分账请求 1、分账接收方类型是PERSONAL_OPENID或PERSONAL_SUB_OPENID时，是个人姓名的密文（选传，传则校验） 此字段的加密的方式为：敏感信息加密说明 2、使用微信支付平台证书中的公钥 3、使用RSAES-OAEP算法进行加密 4、将请求中HTTP头部的Wechatpay-Serial设置为证书序列号
	Name string `json:"name,omitempty" encryption:"EM_APIV3" dc:"分账接收方姓名，可选值"`

	// 1、MERCHANT_ID：商户号 2、PERSONAL_OPENID：个人openid（由父商户APPID转换得到） 3、PERSONAL_SUB_OPENID: 个人sub_openid（由子商户APPID转换得到）
	Type string `json:"type" dc:"分账接收方类型"`
}

// 订单剩余可分金额--------------------------------------------------------------------------------------------------------

type QueryOrderAmountRequest struct {
	// 微信支付订单号
	TransactionId *string `json:"transaction_id" dc:"微信支付返回的交易订单id"`
}

type QueryOrderAmountRes profitsharing.QueryOrderAmountResponse

// 分账结果查询--------------------------------------------------------------------------------------------------------

type QueryOrderRequest struct {
	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid string `json:"sub_mchid,omitempty"`

	// 微信支付订单号
	TransactionId string `json:"transaction_id"`

	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@ 。 微信分账单号与商户分账单号二选一填写
	OutOrderNo string `json:"out_order_no"`
}

type OrdersEntityRes profitsharing.OrdersEntity

// 解冻剩余资金--------------------------------------------------------------------------------------------------------

type UnfreezeOrderRequest struct {
	// 分账的原因描述，分账账单中需要体现
	Description string `json:"description"`

	// 商户系统内部的分账单号，在商户系统内部唯一，同一分账单号多次请求等同一次。只能是数字、大小写字母_-|*@
	OutOrderNo string `json:"out_order_no"`

	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid string `json:"sub_mchid,omitempty"`

	// 微信支付订单号
	TransactionId string `json:"transaction_id"`
}

// 添加分账关系，分账关系绑定--------------------------------------------------------------------------------------------------------

type AddReceiverRequest struct {
	// 类型是MERCHANT_ID时，是商户号 类型是PERSONAL_OPENID时，是个人openid 类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Account string `json:"account"`

	// 微信分配的公众账号ID
	Appid string `json:"appid"`

	// 子商户与接收方具体的关系，本字段最多10个字。  当字段relation_type的值为CUSTOM时，本字段必填 当字段relation_type的值不为CUSTOM时，本字段无需填写
	CustomRelation string `json:"custom_relation,omitempty"`

	// 分账接收方类型是MERCHANT_ID时，是商户全称（必传），当商户是小微商户或个体户时，是开户人姓名 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验） 分账接收方类型是PERSONAL_SUB_OPENID时，是个人姓名（选传，传则校验） 1、此字段需要加密，的加密方法详见：敏感信息加密说明 2、使用微信支付平台证书中的公钥 3、使用RSAES-OAEP算法进行加密 4、将请求中HTTP头部的Wechatpay-Serial设置为证书序列号
	Name string `json:"name,omitempty" encryption:"EM_APIV3"`

	// 子商户与接收方的关系。 本字段值为枚举： SERVICE_PROVIDER：服务商 STORE：门店  STAFF：员工 STORE_OWNER：店主 PARTNER：合作伙伴 HEADQUARTER：总部 BRAND：品牌方 DISTRIBUTOR：分销商 USER：用户 SUPPLIER：供应商 CUSTOM：自定义  * `SERVICE_PROVIDER` - 服务商，  * `STORE` - 门店，  * `STAFF` - 员工，  * `STORE_OWNER` - 店主，  * `PARTNER` - 合作伙伴，  * `HEADQUARTER` - 总部，  * `BRAND` - 品牌方，  * `DISTRIBUTOR` - 分销商，  * `USER` - 用户，  * `SUPPLIER` - 供应商，  * `CUSTOM` - 自定义，
	RelationType *ReceiverRelationType `json:"relation_type"`

	// 子商户的公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要）
	SubAppid string `json:"sub_appid,omitempty"`

	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid string `json:"sub_mchid,omitempty"`

	// 枚举值： MERCHANT_ID：商户ID PERSONAL_OPENID：个人openid（由父商户APPID转换得到） PERSONAL_SUB_OPENID：个人sub_openid（由子商户APPID转换得到）  * `MERCHANT_ID` - 商户号，  * `PERSONAL_OPENID` - 个人openid（由父商户APPID转换得到），  * `PERSONAL_SUB_OPENID` - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要），
	Type *ReceiverType `json:"type"`
}

// ReceiverRelationType   * `SERVICE_PROVIDER` - 服务商，  * `STORE` - 门店，  * `STAFF` - 员工，  * `STORE_OWNER` - 店主，  * `PARTNER` - 合作伙伴，  * `HEADQUARTER` - 总部，  * `BRAND` - 品牌方，  * `DISTRIBUTOR` - 分销商，  * `USER` - 用户，  * `SUPPLIER` - 供应商，  * `CUSTOM` - 自定义，
type ReceiverRelationType string

// ReceiverType   * `MERCHANT_ID` - 商户号，  * `PERSONAL_OPENID` - 个人openid（由父商户APPID转换得到），  * `PERSONAL_SUB_OPENID` - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要），
type ReceiverType string

type AddReceiverRes profitsharing.AddReceiverResponse

// 分账关系查询--------------------------------------------------------------------------------------------------------

// 分账关系解绑--------------------------------------------------------------------------------------------------------

type DeleteReceiverRequest struct {
	// 类型是MERCHANT_ID时，是商户号 类型是PERSONAL_OPENID时，是个人openid 类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Account string `json:"account"`

	// 微信分配的公众账号ID
	Appid string `json:"appid"`

	// 微信分配的子商户公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要）
	SubAppid string `json:"sub_appid,omitempty"`

	// 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要）
	SubMchid string `json:"sub_mchid,omitempty"`

	// 枚举值： MERCHANT_ID：商户ID PERSONAL_OPENID：个人openid（由父商户APPID转换得到） PERSONAL_SUB_OPENID：个人sub_openid（由子商户APPID转换得到）  * `MERCHANT_ID` - 商户号，  * `PERSONAL_OPENID` - 个人openid（由父商户APPID转换得到），  * `PERSONAL_SUB_OPENID` - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要），
	Type *ReceiverType `json:"type"`
}
type DeleteReceiverRes profitsharing.DeleteReceiverResponse
