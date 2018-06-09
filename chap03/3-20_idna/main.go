package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

// 国際化ドメインをアスキーに変換
func main() {
	src := "握力王"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", src, ascii) // => 握力王 -> xn--tfrv51b0yl
}
