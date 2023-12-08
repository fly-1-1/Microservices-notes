package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orm1/models"
	"os"
	"path"
	"time"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {

	//userlist := []models.User{}
	//
	//models.DB.Find(&userlist)
	//
	//c.JSON(200, gin.H{
	//	"result": userlist,
	//})

	//条件查询
	userlist := []models.User{}
	models.DB.Where("age>20").Find(&userlist)
	c.JSON(200, gin.H{
		"result": userlist,
	})

	c.String(200, "用户列表")
}

func (con UserController) Add(c *gin.Context) {

	user := models.User{
		Username: "qjy 2",
		Age:      22,
		Email:    "1111@qq.com",
		AddTime:  int(time.Now().Unix()),
	}
	models.DB.Create(&user)

	fmt.Println(user)
	c.String(200, "增加用户成功")
}

func (con UserController) Edit(c *gin.Context) {

	//user := models.User{Id: 6}
	//models.DB.Find(&user)
	//fmt.Println(user)
	//user.Username = "hahaha"
	//user.Email = "4411@qq.com"
	//models.DB.Save(&user)

	//user := models.User{}
	//models.DB.Model(&user).Where("id = ?", 6).Update("username", "dididi")
	//c.String(200, "修改用户成功")

	user := models.User{}
	models.DB.Where("id = ?", 6).Find(&user)

	user.Username = "哈哈哈"
	user.Email = "111111@qq.com"
	user.AddTime = int(time.Now().Unix())
	models.DB.Save(&user)
}

func (con UserController) Delete(c *gin.Context) {

	//user := models.User{Id: 6}
	//models.DB.Delete(&user)

	user := models.User{}
	models.DB.Where("username = ?", "gorm").Delete(&user)

	c.String(200, "删除用户成功")
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
