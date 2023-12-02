package admin

import (
	"gin14/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

	//判断后缀名是否合法
	extName := path.Ext(fp.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
		".gif":  true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//创建图片保存的目录
	day := models.GetDay()
	dir := "./static/upload/" + day
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		panic(err)
		c.String(200, "MkdirAll 失败")
		return
	}
	//生成文件名称和文件保存的目录
	fileName := models.GetTimeStamp() + extName
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(fp, dst)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}
