package weixin_hook

import (
	"context"
	"github.com/kysion/weixin-library/weixin_model"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
)

type ConsumerHookFunc func(ctx context.Context, info interface{}) int64 // 别人订阅我，通常需要返回sys_user_id

type ConsumerHookInfo struct {
	Key   ConsumerKey
	Value ConsumerHookFunc
}
type ConsumerKey struct {
	weixin_enum.Category `json:"category" dc:"业务类别"`
	weixin_enum.ConsumerAction
}
type UserInfo struct {
	SysUserId int64 `json:"sys_user_id" dc:"筷满客平台用户id"`
	weixin_model.UserInfoRes
}
