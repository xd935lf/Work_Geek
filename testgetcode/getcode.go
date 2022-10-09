package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost/")
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	resCode := res.StatusCode
	res.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf("通过添加的函数拿到的状态码%d\r\n", resCode)
}
