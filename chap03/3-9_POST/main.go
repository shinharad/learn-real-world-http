package main

import (
	"log"
	"net/http"
	"strings"
)

// 任意の文字列をPOST送信
func main() {
	reader := strings.NewReader("テキスト")
	resp, err := http.Post("http://localhost:18888", "test/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
}
