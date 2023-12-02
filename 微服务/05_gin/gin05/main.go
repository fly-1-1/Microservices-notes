package main

import (
	"gin05/routers"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"  form:"username"`
	Password string `json:"password"  form:"password" `
}

type Article struct {
	Title   string `json:"title"  xml:"title"`
	Content string `json:"content"  xml:"content" `
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates\\**\\*")

	//配置静态web服务
	//1 路由 2 映射目录
	r.Static("/static", "./static")

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run(":8080")

}
