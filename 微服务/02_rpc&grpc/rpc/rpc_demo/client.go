package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("网络连接失败")
	}

	var pd int

	err = cli.Call("Panda.GetInfo", 10000, &pd)
	if err != nil {
		fmt.Println("连接失败call")
	}

	fmt.Println("pd:", pd)

}
