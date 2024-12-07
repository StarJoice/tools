//@Date 2024/12/7 10:43
//@Desc

package session

import (
	"github.com/StarJoice/tools/ginx/gctx"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// MiddlewareBuilder 登录校验
type MiddlewareBuilder struct {
	sp Provider
}

func (b *MiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess, err := b.sp.Get(&gctx.Context{Context: ctx})
		if err != nil {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set(CtxSessionKey, sess)
	}
}
