package weixin_hook

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
)

// ServiceNotifyHookFunc 应用通知 - 平台主动发的 对应ServiceNotify
type ServiceNotifyHookFunc func(ctx context.Context, info g.Map) bool

type ServiceNotifyHookInfo struct {
	Key   weixin_enum.ServiceNotifyType
	Value ServiceNotifyHookFunc
}
