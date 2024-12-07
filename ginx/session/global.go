//@Date 2024/12/7 09:43
//@Desc

package session

import (
	"github.com/StarJoice/tools/ginx/gctx"
	"github.com/gin-gonic/gin"
)

const CtxSessionKey = "_session"

var defaultProvider Provider

func NewSession(ctx *gctx.Context, uid int64,
	jwtData map[string]string,
	sessData map[string]any) (Session, error) {
	return defaultProvider.NewSession(
		ctx,
		uid,
		jwtData,
		sessData)
}

// Get 参考 defaultProvider.Get 的说明
func Get(ctx *gctx.Context) (Session, error) {
	return defaultProvider.Get(ctx)
}

func SetDefaultProvider(sp Provider) {
	defaultProvider = sp
}

func DefaultProvider() Provider {
	return defaultProvider
}

func CheckLoginMiddleware() gin.HandlerFunc {
	return (&MiddlewareBuilder{sp: defaultProvider}).Build()
}

func RenewAccessToken(ctx *gctx.Context) error {
	return defaultProvider.RenewAccessToken(ctx)
}

func UpdateClaims(ctx *gctx.Context, claims Claims) error {
	return defaultProvider.UpdateClaims(ctx, claims)
}
