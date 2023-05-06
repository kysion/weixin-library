package weixin_hook

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
)

// ServiceMsgHookFunc 回调消息 - 由某人产生  对应回调CallBack
type ServiceMsgHookFunc func(ctx context.Context, info g.Map) bool // 通常需要返回用户userId

type ServiceMsgHookInfo struct {
	Key   weixin_enum.CallbackMsgType
	Value ServiceMsgHookFunc
}
