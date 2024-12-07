//@Date 2024/12/6 17:53
//@Desc

package ginx

import (
	"errors"
	"github.com/StarJoice/tools/ginx/errs"
	"github.com/StarJoice/tools/ginx/gctx"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// Context 为 `gctx.Context` 类型的别名，
// 这样可以在其他地方使用 `Context` 来代替 `gctx.Context`，
// 方便管理自定义的请求上下文。
type Context = gctx.Context

// WithResult 将一个接受上下文的业务逻辑函数封装成 Gin 的 Handler。
// 它自动处理业务逻辑中的错误情况，并返回对应的 HTTP 状态码和响应数据。
func WithResult(fn func(ctx *Context) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 执行业务逻辑
		res, err := fn(&Context{Context: ctx})
		// 如果是 "不需要响应" 错误，直接返回，不做响应
		if errors.Is(err, errs.ErrNoResponse) {
			slog.Debug("不需要响应", slog.Any("err", err))
			return
		}
		// 如果是未授权错误，返回 401 状态码并停止后续处理
		if errors.Is(err, errs.ErrUnauthorized) {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 处理其他错误，返回 500 错误
		if err != nil {
			slog.Error("执行业务逻辑失败", slog.Any("err", err))
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		// 成功时返回 200 状态码和结果
		ctx.JSON(http.StatusOK, res)
	}
}

// WithRequest 将一个接受请求参数的业务逻辑函数封装成 Gin 的 Handler。
// 它会自动绑定请求参数并处理业务逻辑，同时处理错误情况。
func WithRequest[Req any](fn func(ctx *Context, req Req) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		// (因为不确定绑定参数是什么类型，所以使用shouldBind来尝试自动绑定请求数据)
		if err := ctx.ShouldBind(&req); err != nil {
			slog.Debug("绑定参数失败", slog.Any("err", err))
			ctx.JSON(http.StatusBadRequest, Result{Code: http.StatusBadRequest, Msg: "请求参数绑定失败"})
			return
		}
		// 执行业务逻辑
		res, err := fn(&Context{Context: ctx}, req)
		// 错误处理
		if errors.Is(err, errs.ErrNoResponse) {
			slog.Debug("不需要响应", slog.Any("err", err))
			return
		}
		if errors.Is(err, errs.ErrUnauthorized) {
			slog.Debug("未授权", slog.Any("err", err))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err != nil {
			slog.Error("执行业务逻辑失败", slog.Any("err", err))
			ctx.JSON(http.StatusInternalServerError, res)
			return
		}
		// 成功时返回 200 状态码和结果
		ctx.JSON(http.StatusOK, res)
	}
}
