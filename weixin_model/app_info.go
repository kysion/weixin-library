package weixin_model

// 所有类目------------------------------------------------------------------------------------------------------------

type AppCategoryInfoRes struct {
	ErrorCommon
	CategoryList []CategoryList `json:"category_list" dc:"类目信息列表"`
}

type CategoryList struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	FirstId     int    `json:"first_id"`
	SecondId    int    `json:"second_id"`
}

// 小程序基本信息------------------------------------------------------------------------------------------------------------

type AccountVBasicInfoRes struct {
	ErrorCommon
	Appid             string        `json:"appid"`
	AccountType       int           `json:"account_type"`
	PrincipalType     int           `json:"principal_type"`
	PrincipalName     string        `json:"principal_name"`
	RealnameStatus    int           `json:"realname_status"`
	WxVerifyInfo      WxVerifyInfo  `json:"wx_verify_info"`
	SignatureInfo     SignatureInfo `json:"signature_info"`
	HeadImageInfo     HeadImageInfo `json:"head_image_info"`
	Nickname          string        `json:"nickname"`
	RegisteredCountry int           `json:"registered_country"`
	NicknameInfo      NicknameInfo  `json:"nickname_info"`
	Credential        string        `json:"credential"`
	CustomerType      int           `json:"customer_type"`
}

type WxVerifyInfo struct {
	QualificationVerify bool `json:"qualification_verify"`
	NamingVerify        bool `json:"naming_verify"`
}

type SignatureInfo struct {
	Signature       string `json:"signature"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

type HeadImageInfo struct {
	HeadImageUrl    string `json:"head_image_url"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}
type NicknameInfo struct {
	Nickname        string `json:"nickname"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}
