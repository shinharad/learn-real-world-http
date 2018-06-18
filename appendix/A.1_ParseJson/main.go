package main

import (
	"encoding/json"
	"fmt"
)

// Go言語の構造体タグを使ったJSONのパース

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`

	// omitempty: 出力時に空だったらそのキーごと省略する
	Price int `json:"price,omitempty"`
	// 入出力時に無視する。一時的なキャッシュデータを構造体内に持たせたい場合に指定する
	Comment string `json:"-"`

	// 省略されたか、ゼロ値かを判定する
	Price2 *int `json:"price2"`
}

var jsonString = []byte(`
[
	{"title": "The Art of Community", "author": "Jono Bacon", "price": 123, "comment": "xxx"},
	{"title": "Mithril", "author": "Yoshiki", "comment": "xxx", "price2": 345}
]`)

func main() {
	var books []Book
	err := json.Unmarshal(jsonString, &books)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		if book.Price2 != nil { // 値が入ってなかったらnilになる
			fmt.Println("price2: ", *book.Price2)
		}
		fmt.Println(book)
	}
}
