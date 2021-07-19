package gee

import (
	"net/http"
)

// 通过实现了serveHttp接口，接管所有的http请求

// HandlerFunc define the request handler used by gee
type HandlerFunc func(*Context)  // 换为了context对象包装

// Engine implement the interface of severHttp
type Engine struct {
	router *router   // 添加一张路由映射表router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}


func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc){ // 添加路由映射规则，key由 请求方法-静态路由地址  组成
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc){ // 调用 GET 方法时，会将路由和处理方法注册到映射表router中
	engine.addRoute("GET", pattern, handler)
}

// POST define the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error){  // 是ListenAndServe的包装
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){ // 解析请求的路径，查找路由映射表，查到就执行注册的处理方法；若查不到，返回404
	c := newContext(w, req)
	engine.router.handle(c)
}



