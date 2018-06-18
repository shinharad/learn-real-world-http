package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// JSON読み込み時に型変換をして読み込む日付型

type DueDate struct {
	time.Time
}

// UnmarshalJSONメソッドを持つ構造体は json.Unmarshaler インターフェイスを満たす
func (d *DueDate) UnmarshalJSON(raw []byte) error {
	epoch, err := strconv.Atoi(string(raw))
	if err != nil {
		return err
	}
	d.Time = time.Unix(int64(epoch), 0)
	return nil
}

type ToDo struct {
	Task string  `json:"task"`
	Time DueDate `json:"Due"`
	Done bool
}

var jsonString = []byte(`[
	{"task": "幼稚園投稿", "due": 1486600200},
	{"task": "エリクソン研究会に行く", "due": 1486634400}
]`)

// 日付のシリアライズ
func (d *DueDate) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(d.Unix()))), nil
}

type ToDoList []ToDo

// リストをフィルタリングしてからシリアライズ
func (l ToDoList) MarshalJSON() ([]byte, error) {
	tmpList := make([]ToDo, 0, len(l))
	for _, todo := range l {
		if !todo.Done {
			tmpList = append(tmpList, todo)
		}
	}
	return json.Marshal(tmpList)
}

func main() {
	var todos []ToDo
	err := json.Unmarshal(jsonString, &todos)
	if err != nil {
		panic(err)
	}
	for _, todo := range todos {
		fmt.Printf("%s: %v\n", todo.Task, todo.Time)
	}
}
