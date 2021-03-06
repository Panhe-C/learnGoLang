package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{} // 起别名 gee.H,构建Json数据是显得更简洁

type Context struct {
	// origin objects
	Writer http.ResponseWriter  // 只包含resp，req，提供method和path两个常用属性的直接访问
	Req *http.Request
	// Request info
	Path string
	Method string
	Params map[string]string
	// Response info
	StatusCode int
	// middleware -> 添加了两个参数
	handlers []HandlerFunc
	index int  // 记录当前执行到第几个中间件
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path: req.URL.Path,
		Method: req.Method,
		Req: req,
		Writer: w,
		index: -1,
	}
}

// index记录当前执行到第几个中间件，当在中间件调用Next方法时，控制权交给下一个中间件，直到调用下一个中间件，然后再从后往前，调用每个中间件在Next方法之后定义的部分。
func (c *Context)  Next(){
	c.index++
	s := len(c.handlers)
	for ; c.index < s;c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message":err})
}

func (c *Context) Param(key string) string  {
	value, _ := c.Params[key] // // 解析后的参数存到params中，通过c.param("lang")的方式获取对应的值
	return value
}

//func newContext(w http.ResponseWriter, req *http.Request) *Context {
//	return &Context{
//		Writer: w,
//		Req: req,
//		Path: req.URL.Path,
//		Method: req.Method,
//	}
//}

// 提供访问query和postform参数的方法
func (c *Context) PostForm(key string) string  {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string{
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string)  {
	c.Writer.Header().Set(key, value)
}


// 提供快速构造 string/data/json/html相应的方法
func (c *Context) String(code int, format string, values ...interface{})  {
	c.SetHeader("Context-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{})  {
	c.SetHeader("Context-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil{
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte)  {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string)  {
	c.SetHeader("Content-type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}