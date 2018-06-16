package main

import (
	"log"
	"net/rpc/jsonrpc"
)

// JSON-RPCでサーバーの計算処理を呼び出すクライアント

// Args : 引数
type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	var result int
	args := &Args{4, 5}
	err = client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("4 * 5 = %d\n", result)
}
