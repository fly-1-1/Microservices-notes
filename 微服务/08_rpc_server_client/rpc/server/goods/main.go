package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Goods 创建远程调用函数 函数一般放在结构体里面
type Goods struct {
}

type AddGoodReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodRes struct {
	Success bool
	Message string
}

type GetGoodsReq struct {
	Id int
}

type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func (c Goods) AddGoods(req AddGoodReq, res *AddGoodRes) error {
	//执行增加 模拟
	fmt.Println(req)
	//返回增加的结果
	*res = AddGoodRes{
		Success: true,
		Message: "增加数据成功",
	}
	return nil
}

func (c Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	//执行增加 模拟
	fmt.Println(req)
	//返回增加的结果
	*res = GetGoodsRes{
		Id:      12,
		Title:   "服务器获取的数据",
		Price:   24.5,
		Content: "我是服务器数据库获取的内容",
	}
	return nil
}

func main() {
	err1 := rpc.RegisterName("goods", new(Goods))
	if err1 != nil {
		fmt.Println(err1)
	}

	listener, err2 := net.Listen("tcp", "127.0.0.1:8088")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer listener.Close()
	for {
		fmt.Println("准备建立连接")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		rpc.ServeConn(conn)
	}

}
