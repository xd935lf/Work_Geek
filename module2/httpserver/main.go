// 接收客户端 request，并将 request 中带的
// header 写入 response header
// 读取当前系统的环境变量中的 VERSION 配置，
// 并写入 response header
// Server 端记录访问日志包括客户端 IP，HTTP 返回码，
// 输出到 server 端的标准输出
// 当访问 localhost/healthz 时，应返回 200
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main() {
	//http建立链接 端口号80
	fmt.Println("code running")
	http.HandleFunc("/", ret_Header)
	http.HandleFunc("/healthz", health)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal("error message:%s\n", err.Error())
		return
	}

}

// 获取Request Header写入Response header
// 读取version 写response header
func ret_Header(w http.ResponseWriter, r *http.Request) {
	//request拿到Header

	req_Map := make(map[string][]string)
	for k, v := range r.Header {
		req_Map[k] = v
	}
	//Header写到response里
	for res_k, res_v := range req_Map {
		w.Header().Set(res_k, fmt.Sprint(res_v))
	}
	env_Ver := os.Getenv("VERSION")
	if env_Ver == "" {
		os.Setenv("VERSION", runtime.Version())
	}
	w.Header().Set("VERSION", env_Ver)

	// status_Code := getStatusCode()
	logRecord(r, http.StatusOK)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status ok"))
	logRecord(r, http.StatusOK)
}

// 获取状态码 HTTP statuscode  问题  ：获取不到本地  报错！
// func getStatusCode() int {
// 	u, _ := url.Parse("http://localhost/")
// 	q := u.Query()
// 	u.RawQuery = q.Encode()
// 	res, err := http.Get(u.String())
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return 0
// 	}
// 	resCode := res.StatusCode
// 	res.Body.Close()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return 0
// 	}
// 	fmt.Printf("通过添加的函数拿到的状态码%d\r\n", resCode)
// 	return resCode
// }

// 日志信息
func logRecord(r *http.Request, statusCode int) {
	//截取request handle传来的IP地址
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		log.Println(ip, r.URL, statusCode)
	} else {
		log.Println("", r.URL, statusCode)
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	http.HandleFunc("/user", user)
// 	err := http.ListenAndServe("0.0.0.0:80", nil)
// 	if err != nil {
// 		//fmt.Println(err.Error())
// 		log.Fatal(err)

// 	}
// 	// 	listenner, err := net.Listen("tcp", "0.0.0.0:80")
// 	// 	if err != nil {
// 	// 		fmt.Printf("err: %v\n", err)
// 	// 	}
// 	// 	defer listenner.Close()
// }

// func user(w http.ResponseWriter, r *http.Request) {
// 	req_Map := make(map[string][]string)
// 	for k, v := range r.Header {
// 		req_Map[k] = v
// 	}
// 	env_Ver := os.Getenv("VERSION")

// 	w.Header().Set("VERSION", env_Ver)

// 	for req_k, req_v := range req_Map {
// 		w.Header().Set(req_k, fmt.Sprint(req_v))
// 	}
// 	fmt.Fprintf(w, "我的未来不是梦\n")
// 	fmt.Fprintf(w, "host:", r.Host+"<br>")
// 	fmt.Fprintf(w, "User_Agent:"+r.Header.Get("User-Agent")+"<br>")
// }
