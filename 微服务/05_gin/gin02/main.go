package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserList struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("gin02/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "%v", "首页")
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好1",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好2",
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		a := UserList{
			Title:   "china",
			Desc:    "good",
			Content: "2024",
		}
		c.JSON(200, a)
	})

	//http://localhost:8080/jsonp?callback=xxx  跨域
	r.GET("/jsonp", func(c *gin.Context) {
		a := UserList{
			Title:   "china1",
			Desc:    "good2",
			Content: "2024_1",
		}
		c.JSONP(200, a)
	})

	r.GET("/showxml", func(c *gin.Context) {

		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "xml hello",
		})
	})

	r.GET("/news", func(c *gin.Context) {

		c.HTML(http.StatusOK, "user.html", gin.H{
			"title": "后台数据",
		})
	})

	r.GET("/goods", func(c *gin.Context) {

		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "商品数据",
			"price": 20,
		})
	})

	r.Run()
}
