package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

// クライアント証明書
// サーバーがクライアントに証明書を要求し、正しく承認されたら通信を行う
func main() {
	server := &http.Server{
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert, // クライアント証明書を要求し、検証する
			MinVersion: tls.VersionTLS12,
		},
		Addr: ":18443",
	}
	http.HandleFunc("/", handler)
	log.Println("start http listening :18443")
	err := server.ListenAndServeTLS("server.crt", "server.key")
	log.Println(erR)
}
