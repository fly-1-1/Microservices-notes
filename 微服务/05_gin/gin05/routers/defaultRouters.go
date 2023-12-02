package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "default/index.html", gin.H{
				"msg": "我是msg",
			})
		})

		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}
}
