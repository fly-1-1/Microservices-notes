package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {

	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8088")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer conn.Close()

	var reply AddGoodRes
	err2 := conn.Call("goods.AddGoods", AddGoodReq{
		Id:      10,
		Title:   "商品标题",
		Price:   23.6,
		Content: "商品详情",
	}, &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	//获取微服务返回的数据
	fmt.Println(reply)

	var goodsData GetGoodsRes
	err3 := conn.Call("goods.GetGoods", GetGoodsReq{
		Id: 12,
	}, &goodsData)
	if err3 != nil {
		fmt.Println(err3)
	}

	//获取微服务返回的数据
	fmt.Println(goodsData)

}
