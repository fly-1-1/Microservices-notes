package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)
import "greeter/proto/greeter"

//定义远程调用结构体和方法,结构体需要实现接口 GreeterServer的接口

type Hello struct {
	greeter.UnimplementedGreeterServer
}

func (ts Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)

	return &greeter.HelloRes{
		Message: "你好" + req.Name,
	}, nil
}

func main() {

	//初始化grpc对象
	grpcServer := grpc.NewServer()
	//注册服务
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	//监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	//启动服务
	grpcServer.Serve(listener)

}
