package main

import (
	"fmt"
	"gin14/models"
	"gin14/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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

	//配置redis session中间件
	//创建基于cookie的存储引擎 参数为秘钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("redissession", store))

	//配置全局中间件 可以配置多个
	//r.Use(initMiddleware)

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run(":80")

}
