package main

import (
	"fmt"
	"gin10/models"
	"gin10/routers"
	"github.com/gin-gonic/gin"
	"html/template"
)

type UserInfo struct {
	Username string `json:"username"  form:"username"`
	Password string `json:"password"  form:"password" `
}

type Article struct {
	Title   string `json:"title"  xml:"title"`
	Content string `json:"content"  xml:"content" `
}

func initMiddleware(c *gin.Context) {
	fmt.Println("中间件1-1")
	c.Next() //执行请求 再回来执行中间件
	//c.Abort() //不执行请求而是继续运行中间件
	fmt.Println("中间件1-2")
	//fmt.Println(start - end)
}

func main() {

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	r.LoadHTMLGlob("templates\\**\\*")

	//配置静态web服务
	//1 路由 2 映射目录
	r.Static("/static", "./static")

	//配置全局中间件 可以配置多个
	//r.Use(initMiddleware)

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run(":8080")

}
