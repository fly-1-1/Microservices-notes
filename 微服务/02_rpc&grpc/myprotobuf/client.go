package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "myprotobuf/protobuf"
)

func main() {

	//连接服务器
	conn, err := grpc.Dial("127.0.0.1:10086", grpc.WithInsecure())
	if err != nil {
		fmt.Println("网络异常", err)
	}
	defer conn.Close()

	//获得grpc句柄
	c := pb.NewHelloServerClient(conn)

	re, err := c.SayHello(context.Background(), &pb.HelloReq{Name: "熊猫"})
	if err != nil {
		fmt.Println("SayHello 服务调用失败", err)
	}

	fmt.Println("调用SayHello服务", re.Msg)

	rel, err := c.SayName(context.Background(), &pb.NameReq{Name: "托尼"})
	if err != nil {
		fmt.Println("SayName 服务调用失败", err)
	}

	fmt.Println("调用SayHello服务", rel.Msg)
}
