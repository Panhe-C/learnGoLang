package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc  // 使用map键值对映射
}

func newRouter() *router{
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc)  {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handler参数变为context  提供了动态路由功能
func (r *router) handle(c *Context)  {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok{
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
