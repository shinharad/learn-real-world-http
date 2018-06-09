package main

import (
	"log"
	"net/http"
	"os"
)

// 任意のボディをPOST送信
func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
}
