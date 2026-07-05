package main

import (
	"log"
	"net/http"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() { // 负责启动 http 服务器、注册路由
	log.Println("Starting API Gateway")

	mux := http.NewServeMux() // 创建一个路由器

	mux.HandleFunc("POST /trip/preview", handleTripPreview) // 注册接口 POST /trip/preview

	server := &http.Server{ // 创建http server，默认监听8081端口
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
