package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"greeter_c/proto/greeter"
)

func main() {
	//建立连接  WithTransportCredentials 配置连接级别安全凭证
	grpcClient, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	//注册客户端
	client := greeter.NewGreeterClient(grpcClient)

	res, err := client.SayHello(context.Background(), &greeter.HelloReq{Name: "张三"})
	fmt.Printf("%#v\n", res)
	fmt.Println(res.Message)
}
