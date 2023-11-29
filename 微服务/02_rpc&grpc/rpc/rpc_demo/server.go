package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

/*
- 方法是导出的
- 方法有两个参数，都是导出类型或内建类型
- 方法的第二个参数是指针
- 方法只有一个error接口类型的返回值
事实上，方法必须看起来像这样：
func (t *T) MethodName(argType T1, replyType *T2) error
*/

type Panda int

func (c *Panda) GetInfo(argType int, replyType *int) error {

	fmt.Println("打印对端发送过来的内容为:", argType)

	*replyType = argType + 123

	return nil
}

func main() {
	http.HandleFunc("/panda", pandatext)

	//注册对象
	pd := new(Panda)
	rpc.Register(pd)
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误")
	}

	http.Serve(ln, nil)
}
