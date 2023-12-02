package routers

import (
	"gin14/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)

		apiRouters.GET("/userlist", api.ApiController{}.Userlist)

		apiRouters.GET("/plist", api.ApiController{}.Plist)

	}
}
