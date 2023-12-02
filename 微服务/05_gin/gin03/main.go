package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("C:\\Users\\qjy\\Desktop\\gin\\gin03\\templates\\**\\*")

	//配置静态web服务
	//1 路由 2 映射目录
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin03")
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title":  "后台首页",
			"scores": 60,
			"hobby":  []string{"11", "22", "33"},
			"newsList": []interface{}{
				&Article{
					Title:   "china1",
					Content: "good",
				},
				&Article{
					Title:   "china2",
					Content: "good",
				},
			},
			"news": &Article{
				Title:   "china3",
				Content: "good",
			},
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {

		a := Article{
			Title:   "china",
			Content: "good",
		}
		c.HTML(http.StatusOK, "admin/user.html", gin.H{
			"title": "新闻页面",
			"news":  a,
		})
	})

	r.GET("/default", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "前台首页",
		})
	})

	r.GET("/default/news", func(c *gin.Context) {

		a := Article{
			Title:   "china",
			Content: "good",
		}
		c.HTML(http.StatusOK, "default/user.html", gin.H{
			"title": "新闻页面",
			"news":  a,
		})
	})

	r.Run(":8080")

}
