// @Date 2024/12/6 17:33
// @Desc 定义自定义的 Context 类型，扩展了 Gin 的默认 Context，提供了对请求参数、查询字符串和 Cookie 的处理。

package contextx

import (
	"github.com/StarJoice/tools/common"
	"github.com/gin-gonic/gin"
)

// Context 结构体嵌入了 Gin 的 Context，
// 通过方法重写提供了对请求参数、查询参数和 Cookie 的处理。
type Context struct {
	*gin.Context // 嵌入 Gin 的 Context，直接调用 Gin 提供的所有方法
}

// Param 获取请求 URL 中的路径参数。
// 返回值是一个 `common.AnyValue` 类型，包含请求参数的值。
// 如果参数不存在，返回一个空值。
func (c *Context) Param(key string) common.AnyValue {
	return common.AnyValue{
		Val: c.Context.Param(key), // 从 Gin 的 Context 获取路径参数
	}
}

// Query 获取请求中的查询字符串参数。
// 返回值是一个 `common.AnyValue` 类型，包含查询参数的值。
// 如果查询参数不存在，返回空字符串。
func (c *Context) Query(key string) common.AnyValue {
	return common.AnyValue{
		Val: c.Context.Query(key), // 从 Gin 的 Context 获取查询参数
	}
}

// Cookie 获取请求中的 Cookie。
// 返回值是一个 `common.AnyValue` 类型，包含 Cookie 的值及可能发生的错误。
// 如果 Cookie 存在，返回 Cookie 的值；如果不存在或发生错误，返回错误信息。
func (c *Context) Cookie(key string) common.AnyValue {
	val, err := c.Context.Cookie(key) // 从 Gin 的 Context 获取 Cookie
	return common.AnyValue{
		Val: val, // 如果 Cookie 存在，则值为 Cookie 的内容
		Err: err, // 如果 Cookie 不存在或发生错误，则返回错误信息
	}
}
