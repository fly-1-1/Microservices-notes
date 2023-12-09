package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	conn, err1 := net.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer conn.Close()

	//建立基于json编码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err2 := client.Call("hello.SayHello", "我是客户端", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}

	//获取微服务返回的数据
	fmt.Println(reply)

}
