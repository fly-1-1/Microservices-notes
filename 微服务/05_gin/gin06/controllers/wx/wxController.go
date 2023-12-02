package wx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WxContrller struct{}

func (con WxContrller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是msg",
	})
}

func (con WxContrller) News(c *gin.Context) {
	c.String(200, "新闻")
}
