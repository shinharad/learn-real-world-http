package main

import (
	"fmt"
	"net/http"
)

// Googleと通信してプロトコルバージョンを確認する
//
// $ go run main.go
// Protocol Version: HTTP/2.0
// $ GODEBUG=http2client=0 go run main.go
// Protocol Version: HTTP/1.1
// $ GODEBUG=http2debug=1 go run main.go
//
// https://deeeet.com/writing/2015/11/19/go-http2/
func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Protocol Version: %s\n", resp.Proto)
}
