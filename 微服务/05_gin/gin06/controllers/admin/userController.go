package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//c.String(200, "用户列表--")
	con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	c.String(200, "用户列表add--")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表edit--")
}
