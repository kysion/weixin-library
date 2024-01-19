package weixin_model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

type WeixinConsumerConfig struct {
	Id                 int64       `json:"id"                 description:"id"`
	OpenId             string      `json:"openId"             description:"微信用户openId，不同应用下的用户具备不同的openId"`
	SysUserId          int64       `json:"sysUserId"          description:"用户id"`
	Avatar             string      `json:"avatar"             description:"头像"`
	Province           string      `json:"province"           description:"省份"`
	City               string      `json:"city"               description:"城市"`
	NickName           string      `json:"nickName"           description:"昵称"`
	IsStudentCertified int         `json:"isStudentCertified" description:"是否学生认证"`
	UserType           int         `json:"userType"           description:"用户账号类型，和sysUserType保持一致"`
	UserState          int         `json:"userState"          description:"状态：0未激活、1正常、-1封号、-2异常、-3已注销"`
	IsCertified        int         `json:"isCertified"        description:"是否实名认证"`
	Sex                int         `json:"sex"                description:"性别：0女 1男"`
	AccessToken        string      `json:"accessToken"        description:"授权token"`
	ExtJson            string      `json:"extJson"            description:"拓展字段"`
	CreatedAt          *gtime.Time `json:"createdAt"          description:""`
	UpdatedAt          *gtime.Time `json:"updatedAt"          description:""`
	DeletedAt          *gtime.Time `json:"deletedAt"          description:""`
	UnionId            string      `json:"unionId"            description:"微信用户union_id，同一个公众号下的用户只有一个unionId"`
	SessionKey         string      `json:"sessionKey"         description:"微信用户会话key"`
	RefreshToken       string      `json:"refreshToken"       description:"微信用户授权刷新Token"`
	ExpiresIn          *gtime.Time `json:"expiresIn"          description:"令牌过期时间"`
	AuthState          int         `json:"authState"          description:"微信用户授权状态：1已授权、2未授权"`
	AppType            int         `json:"appType"            description:"应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店"`
	IsFollowPublic     int         `json:"isFollowPublic"     description:"是否关注公众号：1关注、2未关注"`
	AppId              string      `json:"appId"              description:"商家应用Id"`
}

type UpdateConsumerReq struct {
	Id                 int64  `json:"id"                 dc:"id"`
	Avatar             string `json:"avatar"             description:"头像"`
	Province           string `json:"province"           description:"省份"`
	City               string `json:"city"               description:"城市"`
	NickName           string `json:"nickName"           description:"昵称"`
	IsStudentCertified int    `json:"isStudentCertified" description:"是否学生认证"`
	OpenId             string `json:"openId"             description:"微信用户openId，不同应用下的用户具备不同的openId"`
	//AppType            int    `json:"appType"            description:"应用类型：1公众号 2小程序 4网站应用H5  8移动应用  16视频小店"`
	//IsFollowPublic int    `json:"isFollowPublic"     description:"是否关注公众号：1关注、2未关注"`
	//AppId          string `json:"appId"              description:"商家应用Id"`
}

type UpdateConsumerTokenReq struct {
	// Id     int64  `json:"id"                 dc:"id"`
	// OpenId string `json:"openId"             description:"微信用户openId，不同应用下的用户具备不同的openId"`
	// UnionId     string `json:"unionId"            description:"微信用户union_id，同一个公众号下的用户只有一个unionId"`
	AccessToken  *string     `json:"accessToken"        description:"授权token"`
	RefreshToken *string     `json:"refreshToken"       description:"微信用户授权刷新Token"`
	ExpiresIn    *gtime.Time `json:"expiresIn"          description:"令牌过期时间"`

	SessionKey *string `json:"sessionKey"         description:"微信用户会话key"`
}

type WeixinConsumerConfigListRes base_model.CollectRes[WeixinConsumerConfig]
