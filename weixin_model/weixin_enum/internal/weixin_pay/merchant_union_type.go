package weixin_pay

import "github.com/kysion/base-library/utility/enum"

// 特约商户主体类型：1个体工商户、2企业、4事业单位、8社会组织、16政府机关

type MerchantUnionTypeEnum enum.IEnumCode[int]

type merchantUnionType struct {
	Individual         MerchantUnionTypeEnum
	Enterprise         MerchantUnionTypeEnum
	PublicInstitution  MerchantUnionTypeEnum
	SocialOrganization MerchantUnionTypeEnum
	GovernmentAgency   MerchantUnionTypeEnum
}

var MerchantUnionType = merchantUnionType{
	Individual:         enum.New[MerchantUnionTypeEnum](1, "个体工商户"),
	Enterprise:         enum.New[MerchantUnionTypeEnum](2, "企业"),
	PublicInstitution:  enum.New[MerchantUnionTypeEnum](4, "事业单位"),
	SocialOrganization: enum.New[MerchantUnionTypeEnum](8, "社会组织"),
	GovernmentAgency:   enum.New[MerchantUnionTypeEnum](16, "政府机关"),
}

func (e merchantUnionType) New(code int) MerchantUnionTypeEnum {
	if (code & MerchantUnionType.Individual.Code()) == MerchantUnionType.Individual.Code() {
		return MerchantUnionType.Individual
	}

	if (code & MerchantUnionType.Enterprise.Code()) == MerchantUnionType.Enterprise.Code() {
		return MerchantUnionType.Enterprise
	}

	if (code & MerchantUnionType.PublicInstitution.Code()) == MerchantUnionType.PublicInstitution.Code() {
		return MerchantUnionType.PublicInstitution
	}

	if (code & MerchantUnionType.SocialOrganization.Code()) == MerchantUnionType.SocialOrganization.Code() {
		return MerchantUnionType.SocialOrganization
	}

	if (code & MerchantUnionType.GovernmentAgency.Code()) == MerchantUnionType.GovernmentAgency.Code() {
		return MerchantUnionType.GovernmentAgency
	}

	panic("MerchantUnionTypeEnum: error")
}
