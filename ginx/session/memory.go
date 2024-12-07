//@Date 2024/12/7 10:40
//@Desc

package session

import (
	"context"
	"github.com/StarJoice/tools/common"
	"github.com/StarJoice/tools/ginx/errs"
	"github.com/StarJoice/tools/ginx/gctx"
)

var _ Session = &MemorySession{}

// MemorySession 一般用于测试
type MemorySession struct {
	data   map[string]any
	claims Claims
}

func (m *MemorySession) Destroy(ctx context.Context) error {
	return nil
}

func (m *MemorySession) UpdateClaims(ctx *gctx.Context, claims Claims) error {
	return nil
}

func (m *MemorySession) Del(ctx context.Context, key string) error {
	delete(m.data, key)
	return nil
}

func NewMemorySession(cl Claims) *MemorySession {
	return &MemorySession{
		data:   map[string]any{},
		claims: cl,
	}
}

func (m *MemorySession) Set(ctx context.Context, key string, val any) error {
	m.data[key] = val
	return nil
}

func (m *MemorySession) Get(ctx context.Context, key string) common.AnyValue {
	val, ok := m.data[key]
	if !ok {
		return common.AnyValue{Err: errs.ErrSessionKeyNotFound}
	}
	return common.AnyValue{Val: val}
}

func (m *MemorySession) Claims() Claims {
	return m.claims
}
