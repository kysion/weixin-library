package weixin_model

// AuthorizationCodeRes 授权结束后返回的回调数据
type AuthorizationCodeRes struct {
	AuthorizationCode string `json:"authorization_code" dc:"授权码"`
	ExpiresIn         int    `json:"expires_in" dc:"有效期，单位：秒"`
}

// AuthorizerAccessTokenReq 获取商家应用Token
type AuthorizerAccessTokenReq struct {
	//ComponentAccessToken   string `json:"component_access_token" dc:"第三方平台component_access_token，不是authorizer_access_token"`
	ComponentAppid         string `json:"component_appid" dc:"第三方平台 appid"`
	AuthorizerAppid        string `json:"authorizer_appid" dc:"授权方 appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token" dc:"刷新令牌，获取授权信息时得到"`
}

// AuthorizerAccessTokenRes 接口调用凭据
type AuthorizerAccessTokenRes struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token" dc:"接口调用凭据 authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in" dc:"有效期，单位：秒"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token" dc:"刷新token"`
}

// AuthorizationCodeReq 授权通知结构体 和上面的EventMessageBody一样
type AuthorizationCodeReq struct {
	AppId                        string `json:"app_id" dc:" 第三方平台 appid"`
	CreateTime                   int    `json:"create_time" dc:"时间戳"`
	InfoType                     string `json:"info_type" dc:"通知类型，详见InfoType 说明"`
	AuthorizerAppid              string `json:"authorizer_appid" dc:"公众号或小程序的 appid"`
	AuthorizationCode            string `json:"authorization_code" dc:"授权码，可用于获取授权信息"`
	AuthorizationCodeExpiredTime int    `json:"authorization_code_expired_time" dc:"授权码过期时间 单位秒"`
	PreAuthCode                  string `json:"pre_auth_code" dc:"预授权码"`
}

// QueryAuthReq 使用授权码获取授权信息Req
type QueryAuthReq struct {
	ComponentAccessToken string `json:"component_access_token" dc:"第三方平台component_access_token，不是authorizer_access_token"`
	ComponentAppid       string `json:"component_appid" dc:"第三方平台 appid"`
	AuthorizationCode    string `json:"authorization_code" dc:"授权码, 会在授权成功时返回给第三方平台，详见第三方平台授权流程说明"`
}

// AuthorizationInfoRes 使用授权码获取授权信息 Res  包括商家的authorizer_access_token
type AuthorizationInfoRes struct {
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}

type ConfirmInfo struct {
	NeedConfirm    int `json:"need_confirm"`
	AlreadyConfirm int `json:"already_confirm"`
	CanConfirm     int `json:"can_confirm"`
}

type AuthorizationInfo struct {
	AuthorizerAppid       string `json:"authorizer_appid" dc:"授权方 appid"`
	ExpiresIn             int    `json:"expires_in" dc:"authorizer_access_token 的有效期（在授权的公众号/小程序具备API权限时，才有此返回值），单位：秒"`
	AuthorizerAccessToken string `json:"authorizer_access_token" dc:"接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值）"`
	// 刷新令牌（在授权的公众号具备API权限时，才有此返回值），刷新令牌主要用于第三方平台获取和刷新已授权用户的 authorizer_access_token。一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌。用户重新授权后，之前的刷新令牌会失效
	AuthorizerRefreshToken string            `json:"authorizer_refresh_token" dc:"刷新令牌"`
	FuncInfo               FuncscopeCategory `json:"func_info" dc:"授权给开发者的权限集列表"`
}

type FuncInfo struct {
	FuncscopeCategory
}
type FuncscopeCategory struct {
	Id int64 `json:"id" dc:"权限id"`
}

// POST https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=COMPONENT_ACCESS_TOKEN

// AuthorizerInfoReq 获取授权信息Req
type AuthorizerInfoReq struct {
	ComponentAppid    string `json:"component_appid" dc:"第三方平台 appid" `
	AuthorizerAppid   string `json:"authorizer_appid" dc:"授权商家应用的 appid"`
	AuthorizationCode string `json:"authorization_code" dc:"授权码, 会在授权成功时返回给第三方平台"`
}
