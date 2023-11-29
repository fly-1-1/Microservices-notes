package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "myprotobuf/protobuf"
	"net"
)

type server struct {
	pb.UnimplementedHelloServerServer
}

func (c *server) SayHello(ctx context.Context, in *pb.HelloReq) (out *pb.HelloRsp, err error) {

	return &pb.HelloRsp{
		Msg: "hello" + in.Name,
	}, nil
}
func (c *server) SayName(ctx context.Context, in *pb.NameReq) (out *pb.NameRsp, err error) {
	return &pb.NameRsp{
		Msg: in.Name + "晚上好",
	}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误", err)
	}

	srv := grpc.NewServer()
	pb.RegisterHelloServerServer(srv, &server{})

	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络错误", err)
	}

}
