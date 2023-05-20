package weixin_pay

type pay struct {
	MerchantType merchantType
}

var Pay = pay{
	MerchantType: MerchantType,
}
