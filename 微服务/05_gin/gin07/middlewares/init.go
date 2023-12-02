package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Init(c *gin.Context) {
	//判断用户是否登陆

	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Set("username", "张三")

	cCp := c.Copy()

	//定义一个go程统计日志
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done int path:" + cCp.Request.URL.Path)
	}()
}
