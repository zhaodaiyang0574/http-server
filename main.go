package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func ReadUserIP(r *http.Request) string {
	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {
	//当访问 {url}/healthz 时，应返回200
	w.Write([]byte("working"))
}

func analyseHeader(w http.ResponseWriter, r *http.Request) {
	//02读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)

	w.Write([]byte("Hello, this is server"))
	//01接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}

	fmt.Println(r.RemoteAddr)
	fmt.Printf("Ip is: %s\n", ReadUserIP(r))
	log.Printf("Success! Response code: %d", 200)
}

func main() {
	serverMux := http.NewServeMux()
	// 05. debug
	serverMux.HandleFunc("/debug/pprof/", pprof.Index)
	serverMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	serverMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	serverMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	serverMux.HandleFunc("/", analyseHeader)
	serverMux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", serverMux)
	if err == nil {
		log.Fatalf("Starting server failed on port 8080 failed")
	}
}
