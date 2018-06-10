package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

// DELETEメソッドのリクエストを送信
func main() {

	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))

}
