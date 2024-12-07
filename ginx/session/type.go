//@Date 2024/12/6 21:45
//@Desc

package session

import (
	"context"
	"github.com/StarJoice/tools/common"
	"github.com/StarJoice/tools/ginx/errs"
	"github.com/StarJoice/tools/ginx/gctx"
)

type Session interface {
	// Set 将数据写入到 Session 里面
	Set(ctx context.Context, key string, val any) error
	// Get 从 Session 中获取数据，注意，这个方法不会从 JWT 里面获取数据
	Get(ctx context.Context, key string) common.AnyValue
	// Del 删除对应的数据
	Del(ctx context.Context, key string) error
	// Destroy 销毁整个 Session
	Destroy(ctx context.Context) error
	// Claims 编码进去了 JWT 里面的数据
	Claims() Claims
}

type Provider interface {
	// NewSession 将会初始化 Session
	// 其中 jwtData 将编码进去 jwt 中
	// sessData 将被放进去 Session 中
	NewSession(ctx *gctx.Context, uid int64, jwtData map[string]string,
		sessData map[string]any) (Session, error)
	// Get 尝试拿到 Session，如果没有，返回 error
	// Get 必须校验 Session 的合法性。
	// 也就是，用户可以预期拿到的 Session 永远是没有过期，直接可用的
	Get(ctx *gctx.Context) (Session, error)

	// UpdateClaims 修改 claims 的数据
	// 但是因为 jwt 本身是不可变的，所以实际上这里是重新生成了一个 jwt 的 token
	// 必须传入正确的 SSID
	UpdateClaims(ctx *gctx.Context, claims Claims) error

	// RenewAccessToken 刷新并且返回一个新的 access token
	// 这个过程会校验长 token 的合法性
	RenewAccessToken(ctx *gctx.Context) error
}

type Claims struct {
	Uid  int64
	SSID string
	Data map[string]string
}

func (c Claims) Get(key string) common.AnyValue {
	val, ok := c.Data[key]
	if !ok {
		return common.AnyValue{Err: errs.ErrSessionKeyNotFound}
	}
	return common.AnyValue{Val: val}
}
