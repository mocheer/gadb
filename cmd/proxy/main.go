package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

// 修改请求的函数
func modifyRequest(r *http.Request) {
	// 示例：修改请求的 URL
	// if strings.Contains(r.URL.String(), "example.com") {
	// 	r.URL, _ = url.Parse("https://newexample.com/newpath")
	// }

	// // 示例：修改请求头
	// r.Header.Set("User-Agent", "Custom-User-Agent")

	// 示例：修改请求体（如果需要）
	// 注意：修改请求体需要读取并重新设置 r.Body
}

// 反向代理处理函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 打印原始请求
	log.Printf("Received request: %s %s", r.Method, r.URL.String())

	// 修改请求
	modifyRequest(r)

	// 创建一个反向代理
	proxy := httputil.NewSingleHostReverseProxy(r.URL)

	// 转发请求
	proxy.ServeHTTP(w, r)
}

func main() {
	// 启动代理服务器
	http.HandleFunc("/", handleRequest)
	log.Println("Starting proxy server on :8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("Error starting proxy server: ", err)
	}
}
