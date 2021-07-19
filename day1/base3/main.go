package main

// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]
// curl http://localhost:9999/world
// 404 NOT FOUND: /world

import (
	"fmt"
	"net/http"

	"gee"
)

func main()  {
	r := gee.New() //使用new创建gee实例
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {  // 使用GET(*)方法添加路由
		fmt.Println(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Println(w, "Header[%q]\n", k, v)
		}
	})

	r.Run(":9999")  // 使用Run启动web服务
}

