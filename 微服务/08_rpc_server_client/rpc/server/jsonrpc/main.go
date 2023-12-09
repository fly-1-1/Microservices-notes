package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Hello 定义一个远程调用的函数
type Hello struct {
}

// SayHello
/*
方法只能有两个可以序列化的参数
req表示获取客户端传过来的数据
res表示给客户端返回数据
参数不可为管道.函数
*/
func (c Hello) SayHello(req string, res *string) error {
	*res = "你不好" + req
	return nil
}

func main() {
	//注册rpc服务
	err1 := rpc.RegisterName("hello", new(Hello))
	if err1 != nil {
		fmt.Println(err1)
	}
	//监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer listener.Close()

	for {
		fmt.Println("开始建立连接")
		//建立连接
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println("建立连接失败")
		}
		//rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
