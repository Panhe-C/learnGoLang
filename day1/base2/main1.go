package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all request
type Engine struct {} //定义了个空结构体


func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){  // 实现了方法ServeHttp
	switch req.URL.Path {
	case "/":
		fmt.Println(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header{
			fmt.Println(w, "Header[%q] = q%\n", k, v)
		}
	default:
		fmt.Println(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main(){
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine)) // 传入engine实例
}