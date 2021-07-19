package gee

import (
	"log"
	"net/http"
)


// 将Engine作为最顶层的分组，使engine拥有RouterGroup所有的能力
type Engine struct{
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups   golang的嵌套类型，类似Java/python等语言的继承，这样Engine可以拥有RouterGroup的属性
}

type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc   // support middleware
	parent *RouterGroup  		// support nesting
	engine *Engine 				// all groups share a Engine instance 在group中保存一个指针，指向Engine，框架的资源由Engine统一协调，就可以通过Engine间接地访问各种接口
}


// HandlerFunc define the request handler used by gee
type HandlerFunc func(*Context)  // 换为了context对象包装


// 可以将和路由相关的函数都交给routergroup实现

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc)  {
	pattern := group.prefix + comp
	log.Printf("Route %4 - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)    // 该步骤实现了路由的映射，由于Engine某种意义上继承了 RouterGroup的所有属性和方法，因为(*Engine).engine是指向自己的
	// 这样实现，既可以像原来一样添加路由，也可以通过分组添加路由
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc)  {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}


// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error){  // 是ListenAndServe的包装
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){ // 解析请求的路径，查找路由映射表，查到就执行注册的处理方法；若查不到，返回404
	c := newContext(w, req)
	engine.router.handle(c)
}



