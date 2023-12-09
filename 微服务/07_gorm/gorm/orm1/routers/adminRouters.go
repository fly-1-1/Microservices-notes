package routers

import (
	"github.com/gin-gonic/gin"
	"orm1/controllers/admin"
	"orm1/middlewares"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.Init) //在这里可以加载中间件
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)
		adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)
		adminRouters.POST("/user/doEdit", admin.UserController{}.DoEdit)
		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)

		adminRouters.GET("/nav", admin.NavController{}.Index)

		adminRouters.GET("/student", admin.StudentController{}.Index)

		adminRouters.GET("/bank", admin.BankController{}.Index)
	}
}
