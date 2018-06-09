package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// multipart/form-data形式でファイルの送信
func main() {
	var buffer bytes.Buffer // io.Reader であると同時に io.Writer でもある
	writer := multipart.NewWriter(&buffer)

	writer.WriteField("name", "Michael Jacson")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	defer writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
