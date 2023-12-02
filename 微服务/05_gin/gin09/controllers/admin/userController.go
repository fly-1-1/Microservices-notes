package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//c.String(200, "用户列表--")
	con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
	//c.String(http.StatusOK, "%v", "string")
}

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := path.Join("./static/upload", file.Filename)
		fmt.Println(file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}
