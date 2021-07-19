package gee

import (
	"log"
	"time"
)

// 中间件定义与路由映射的handler一致，处理的输入是Context对象。
// 插入点是框架接受到请求初始化context对象之后，允许用户使用自己定义的中间件做一些额外的处理，例如记录日志等，以及对context进行二次加工
// 另外通过调用 (*Context).Next()函数，中间件可等待用户自己定义的Handler处理结束后，再做一些额外的操作。例如计算本次处理所用时间
//

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()  // 表示等待执行其他的中间件
		// Calcaulate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
