// @Date 2024/12/6 17:52
// @Desc 包含了 Gin 路由处理接口和结构体的定义，用于扩展和管理公共路由和私有路由。

package ginx

import (
	"github.com/StarJoice/tools/ginx/gctx"
	"github.com/gin-gonic/gin"
)

// Handler 定义了处理路由的接口，包含两个方法：
// - PrivateRoutes: 用于注册私有路由
// - PublicRoutes: 用于注册公共路由
// 实现该接口的结构体必须提供这两个方法。
type Handler interface {
	PrivateRoutes(server *gin.Engine)
	PublicRoutes(server *gin.Engine)
}

// Result 是一个通用的 API 响应结构体，用于标准化响应格式。
// - Code: 响应的状态码，通常用于表示请求是否成功。(一般为自定义的系统内的错误码)
// - Msg: 响应的消息描述，通常用于提示成功或错误信息。
// - Data: 响应的数据部分，类型为 `any`，可以是任何类型，表示返回的具体数据。
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Context 为 `gctx.Context` 类型的别名，
// 这样可以在其他地方使用 `Context` 来代替 `gctx.Context`，
// 方便管理自定义的请求上下文。
type Context = gctx.Context
