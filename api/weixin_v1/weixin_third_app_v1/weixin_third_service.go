package weixin_third_app_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model"
)

// 微信第三方应用相关服务

type GetAuthorizerListReq struct {
	g.Meta `path:"/getAuthorizerList" method:"post" summary:"获取已授权账号列表" tags:"WeiXin服务商服务"`
	weixin_model.GetAuthorizerList
}

// 获取公众号关联的小程序
