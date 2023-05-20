package weixin_pay

import "github.com/kysion/base-library/utility/enum"

// 特约商户主体类型：1个体工商户、2企业、4事业单位、8社会组织、16政府机关

type MerchantUnionTypeEnum enum.IEnumCode[int]

//type merchantUnionType struct {
//	Individual         MerchantUnionTypeEnum
//	Enterprise         MerchantUnionTypeEnum
//	PublicInstitution  MerchantUnionTypeEnum
//	SocialOrganization MerchantUnionTypeEnum
//	GovernmentAgency   MerchantUnionTypeEnum
//}
//
//var MerchantUnionType = merchantUnionType{
//	Individual:   enum.New[MerchantUnionTypeEnum](1, "个体工商户"),
//	SubMerchant:  enum.New[MerchantUnionTypeEnum](2, "企业"),
//	ShopMerchant: enum.New[MerchantUnionTypeEnum](4, "事业单位"),
//
//	SubMerchant:  enum.New[MerchantUnionTypeEnum](2, "企业"),
//	ShopMerchant: enum.New[MerchantUnionTypeEnum](4, "事业单位"),
//}
//
//func (e merchantUnionType) New(code int) MerchantUnionTypeEnum {
//	if (code & MerchantUnionType.Individual.Code()) == MerchantUnionType.Individual.Code() {
//		return MerchantUnionType.Individual
//	}
//
//	if (code & MerchantUnionType.SubMerchant.Code()) == MerchantUnionType.SubMerchant.Code() {
//		return MerchantUnionType.SubMerchant
//	}
//
//	if (code & MerchantUnionType.ShopMerchant.Code()) == MerchantUnionType.ShopMerchant.Code() {
//		return MerchantUnionType.ShopMerchant
//	}
//
//	if (code & MerchantUnionType.ShopMerchant.Code()) == MerchantUnionType.ShopMerchant.Code() {
//		return MerchantUnionType.ShopMerchant
//	}
//
//	if (code & MerchantUnionType.ShopMerchant.Code()) == MerchantUnionType.ShopMerchant.Code() {
//		return MerchantUnionType.ShopMerchant
//	}
//
//	panic("MerchantUnionTypeEnum: error")
//}
