package routers

import (
	"gin08/controllers/wx"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", wx.WxContrller{}.Index)

		defaultRouters.GET("/news", wx.WxContrller{}.News)
	}
}
