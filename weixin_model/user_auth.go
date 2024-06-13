package weixin_model

import "github.com/gogf/gf/v2/os/gtime"

// AccessTokenRes 微信AccessToken返回数据结构
type AccessTokenRes struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`

	AccessToken    string `json:"access_token"`
	ExpiresIn      int64  `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	Openid         string `json:"openid"`
	Scope          string `json:"scope"`
	Unionid        string `json:"unionid"`
	IsSnapshotuser int    `json:"is_snapshotuser"`
}

// UserInfoRes 微信用户信息返回数据结构
type UserInfoRes struct {
	OpenID     string `json:"openid" dc:"用户openId"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid" dc:"用户unionId"`

	AccessToken  string      `json:"accessToken"        description:"授权token"`
	RefreshToken string      `json:"refreshToken"       description:"微信用户授权刷新Token"`
	ExpiresIn    *gtime.Time `json:"expiresIn"          description:"令牌过期时间"`

	NickName   string `json:"nickname" dc:"昵称"`
	Sex        int    `json:"sex" dc:"性别"`
	Province   string `json:"province" dc:"省份"`
	City       string `json:"city" dc:"城市"`
	Country    string `json:"country" dc:"城镇"`
	HeadImgURL string `json:"headimgurl" dc:"头像"`
	ErrCode    int    `json:"errcode,omitempty" dc:""`
	ErrMsg     string `json:"errmsg,omitempty" dc:""`
}

type OpenIdAndSessionKeyReq struct {
	//componentAccessToken	string `json:"component_access_token" dc:"口调用凭证，该参数为 URL 参数，非 Body 参数。使用component_access_token"`
	Appid          string `json:"appid" dc:"小程序的 AppID"`
	GrantType      string `json:"grant_type" dc:"填authorization_code"`
	ComponentAppid string `json:"component_appid" dc:"第三方平台 appid"`
	JsCode         string `json:"js_code" dc:"wx.login 获取的 code"`
}

type OpenIdAndSessionKeyRes struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`

	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}
