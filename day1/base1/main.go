package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil)) // 启动web服务  第一个参数是地址，:9999表示在 9999 端口监听。而第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。第二个参数，则是我们基于net/http标准库实现Web框架的入口。
}


//hander echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request){
	fmt.Println(w, "URL.Path = %q\n", req.URL.Path)
}


//handler echoes r.Url.Path
func helloHandler(w http.ResponseWriter, req *http.Request){
	for k, v := range req.Header{
		fmt.Println(w, "Header [%q] = %q\n", k ,v)
	}
}

