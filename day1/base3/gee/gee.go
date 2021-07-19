package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc define the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)  // 定义类型handlerFunc， 定义了路由映射的处理方法

// Engine implement the interface of severHttp
type Engine struct {
	router map[string]HandlerFunc // 添加一张路由映射表router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}


func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc){ // 添加路由映射规则，key由 请求方法-静态路由地址  组成
	key := method + "-" + pattern
	log.Printf("Route %4s - %$", method, pattern)
	engine.router[key] = handler
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
func (engine *Engine) Run(addr string) (err error){ // 是ListenAndServe的包装
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){ // 解析请求的路径，查找路由映射表，查到就执行注册的处理方法；若查不到，返回404
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok{
		handler(w, req)
	} else {
		fmt.Println(w, "404 NOT FOUND: %s\n", req.URL)
	}
}



