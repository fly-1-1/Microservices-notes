package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `json:"username"  form:"username"`
	Password string `json:"password"  form:"password" `
}

type Article struct {
	Title   string `json:"title"  xml:"title"`
	Content string `json:"content"  xml:"content" `
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("C:\\Users\\qjy\\Desktop\\gin\\gin04\\templates\\**\\*")

	//配置静态web服务
	//1 路由 2 映射目录
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {

		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	r.GET("/article", func(c *gin.Context) {

		id := c.DefaultQuery("id", "1")

		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"msg": "新闻详情",
		})

	})

	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/user.html", gin.H{})
	})

	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	//获取 GET POST 传递的数据绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {

		user := &UserInfo{}

		err := c.ShouldBind(&user)

		fmt.Println("pass:", user.Password)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}

		c.JSON(http.StatusOK, user)
	})

	r.POST("/doAddUser2", func(c *gin.Context) {
		user := &UserInfo{}
		err := c.ShouldBind(&user)
		fmt.Println("pass:", user.Password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
		c.JSON(http.StatusOK, user)
	})

	//获取POST xml数据
	r.POST("/xml", func(c *gin.Context) {
		xmlData, _ := c.GetRawData()
		article := &Article{}
		err := xml.Unmarshal(xmlData, &article)

		fmt.Println(xmlData)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
		c.JSON(http.StatusOK, article)

	})

	//动态路由传值  list/123
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(200, "%v", cid)
	})

	r.Run(":8080")

}
