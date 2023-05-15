package weixin_model

type GetAuthorizerList struct {
	ComponentAppid string `json:"component_appid"          description:"服务商第三方应用APPID"`
	Offset         int    `json:"offset" dc:"偏移位置/起始位置"`
	Count          int    `json:"count" dc:"拉取数量，最大为 500"`
}

type GetAuthorizerListRes struct {
	TotalCount int    `json:"total_count" dc:"授权的帐号总数"`
	List       []List `json:"list"`
}
type List struct {
	AuthorizerAppid string `json:"authorizer_appid" dc:"已授权帐号的 appid"`
	RefreshToken    string `json:"refresh_token" dc:"刷新令牌authorizer_refresh_token"`
	AuthTime        int    `json:"auth_time" dc:"授权的时间"`
}

type GetOpenAccountRes struct {
	OpenAppid string `json:"open_appid" dc:"应用绑定的开放平台账号AppID"`
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
}
