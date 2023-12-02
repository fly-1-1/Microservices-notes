package routers

import (
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "后台首页")
		})
		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(200, "用户列表add")
		})
		adminRouters.GET("/user/add", func(c *gin.Context) {
			c.String(200, "用户列表add")
		})
		adminRouters.GET("/user/edit", func(c *gin.Context) {
			c.String(200, "用户列表edit")
		})

		adminRouters.GET("/article", func(c *gin.Context) {
			c.String(200, "新闻列表")
		})

		adminRouters.GET("/article/add", func(c *gin.Context) {
			c.String(200, "新闻列表add")
		})

		adminRouters.GET("/article/edit", func(c *gin.Context) {
			c.String(200, "新闻列表edit")
		})
	}
}
