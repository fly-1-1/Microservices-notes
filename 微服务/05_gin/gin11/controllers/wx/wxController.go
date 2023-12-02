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
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是msg",
		"t":   1234345345,
	})
}

func (con WxContrller) News(c *gin.Context) {
	c.String(200, "新闻")
}
