package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

// クッキーの送受信
// http.Client構造体を使用して、クッキーの送受信を行う
func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Jar: jar,
	}
	for i := 0; i < 2; i++ {
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}
