package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	// TODO 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	http.HandleFunc("/go", myHandler)

	// 4.当访问 localhost/healthz 时，应返回 200
	http.HandleFunc("/healthz", healthzHandler)

	http.ListenAndServe(":8888", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	for i, v := range r.Header {
		var vals string
		for _, val := range v {
			vals += val
		}
		w.Header().Set(i, vals)
	}

	// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Add("goroot", os.Getenv("GOROOT"))

	// TODO 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	remoteAddr := r.RemoteAddr
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	log.Printf("%s\n", remoteAddr)
}

// 4.当访问 localhost/healthz 时，应返回 200
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
