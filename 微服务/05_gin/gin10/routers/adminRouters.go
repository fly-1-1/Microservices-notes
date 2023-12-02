package routers

import (
	"gin10/controllers/admin"
	"gin10/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.Init) //在这里可以加载中间件
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.POST("/user/doEdit", admin.UserController{}.DoEdit)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
	}
}
