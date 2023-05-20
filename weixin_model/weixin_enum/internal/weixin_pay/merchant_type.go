package weixin_pay

import "github.com/kysion/base-library/utility/enum"

// MerchantTypeEnum 商户号类型：1服务商、2商户、4门店商家
type MerchantTypeEnum enum.IEnumCode[int]

type merchantType struct {
	SpMerchant   MerchantTypeEnum
	SubMerchant  MerchantTypeEnum
	ShopMerchant MerchantTypeEnum
}

var MerchantType = merchantType{
	SpMerchant:   enum.New[MerchantTypeEnum](1, "服务商"),
	SubMerchant:  enum.New[MerchantTypeEnum](2, "商户"),
	ShopMerchant: enum.New[MerchantTypeEnum](4, "门店商家"),
}

func (e merchantType) New(code int) MerchantTypeEnum {
	if (code & MerchantType.SpMerchant.Code()) == MerchantType.SpMerchant.Code() {
		return MerchantType.SpMerchant
	}

	if (code & MerchantType.SubMerchant.Code()) == MerchantType.SubMerchant.Code() {
		return MerchantType.SubMerchant
	}

	if (code & MerchantType.ShopMerchant.Code()) == MerchantType.ShopMerchant.Code() {
		return MerchantType.ShopMerchant
	}

	panic("MerchantTypeEnum: error")
}
