package wx

import (
	"fmt"
	"gin14/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WxContrller struct{}

func (con WxContrller) Index(c *gin.Context) {
	fmt.Println(models.UnixToTime(1234345345))
	//二级域名共享cookie
	c.SetCookie("username", "张三", 3600, "/", ".qjy.com", false, true)
	//c.SetCookie("hobby", "吃饭 睡觉", 5, "/", "localhost", false, true)

	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是msg",
		"t":   1234345345,
	})
}

func (con WxContrller) News(c *gin.Context) {
	username, _ := c.Cookie("username")
	//hobby, _ := c.Cookie("hobby")
	//c.String(200, "cookie:%v --- hobby:%v", username, hobby)
	c.String(200, "cookie:%v", username)
}

func (con WxContrller) Shop(c *gin.Context) {
	username, _ := c.Cookie("username")
	//hobby, _ := c.Cookie("hobby")
	//c.String(200, "cookie:%v --- hobby:%v", username, hobby)
	c.String(200, "cookie:%v", username)
}

func (con WxContrller) DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(200, "删除成功")
}
