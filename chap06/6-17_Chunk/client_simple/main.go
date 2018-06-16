package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
)

// 改行文字単位でチャンクごとに受信
func main() {
	resp, err := http.Get("http://localhost:18888/chunked")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}
