package weixin_hook

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/weixin-library/weixin_model/weixin_enum"
)

// 异步通知Hook

type NotifyHookFunc func(ctx context.Context, info gmap.Map, hookInfo NotifyKey) bool

type NotifyHookInfo struct {
	Key   NotifyKey
	Value NotifyHookFunc
}

type NotifyKey struct {
	HookCreatedAt gtime.Time `json:"hook_created_at" dc:"Hook创建时间"`
	HookExpireAt  gtime.Time `json:"hook_expire_at" dc:"Hook有效期"`
	weixin_enum.NotifyType
	OrderId string `json:"order_id" dc:"交易id"`
}
