package admin

import (
	"github.com/gin-gonic/gin"
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

func (con UserController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
	//c.String(http.StatusOK, "%v", "string")
}

func (con UserController) DoEdit(c *gin.Context) {
	username := c.PostForm("username")
	fp1, err := c.FormFile("face1")
	if err != nil {
		panic(err)
	}
	dst1 := path.Join("./static/upload", fp1.Filename)
	c.SaveUploadedFile(fp1, dst1)

	fp2, err := c.FormFile("face2")
	if err != nil {
		panic(err)
	}
	dst2 := path.Join("./static/upload", fp2.Filename)
	c.SaveUploadedFile(fp2, dst2)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst1":     dst1,
		"dst2":     dst2,
	})
}

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	fp, err := c.FormFile("face")
	if err != nil {
		panic(err)
	}
	dst := path.Join("./static/upload", fp.Filename)
	c.SaveUploadedFile(fp, dst)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}
