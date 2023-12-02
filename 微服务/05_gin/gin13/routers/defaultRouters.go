package routers

import (
	"gin14/controllers/wx"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", wx.WxContrller{}.Index)

		defaultRouters.GET("/news", wx.WxContrller{}.News)

		defaultRouters.GET("/shop", wx.WxContrller{}.Shop)

		defaultRouters.GET("/deleteCookie", wx.WxContrller{}.DeleteCookie)
	}
}
