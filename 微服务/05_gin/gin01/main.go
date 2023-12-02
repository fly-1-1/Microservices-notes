package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//配置路由
	r.GET("/", func(context *gin.Context) {
		context.String(200, "值:%v", "你好gin")
	})

	r.GET("/news", func(context *gin.Context) {
		context.String(http.StatusOK, "值:%v", "新闻")
	})

	r.POST("/add", func(context *gin.Context) {
		context.String(200, "post返回")
	})

	r.PUT("/edit", func(context *gin.Context) {
		context.String(200, "跟新")
	})

	r.DELETE("/del", func(context *gin.Context) {
		context.String(200, "删除")
	})

	r.Run(":8000")
}
