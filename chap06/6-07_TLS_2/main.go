package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

// OSに依存せずに証明書を読み込んでHTTPS通信を行うクライアントコード
func main() {

	cert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	// x509: IOSで定められた証明書の形式
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	tlsConfig := &tls.Config{
		RootCAs: certPool,
		// 証明書を確認しない設定
		// InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))

}
