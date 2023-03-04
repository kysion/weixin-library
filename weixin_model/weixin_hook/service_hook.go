package weixin_hook

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
)

type ServiceMsgHookFunc func(ctx context.Context, info g.Map) bool

type ServiceMsgHookInfo struct {
	Key   weixin_enum.InfoType
	Value ServiceMsgHookFunc
}
