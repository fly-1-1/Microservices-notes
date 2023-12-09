package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orm1/models"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {

	//var articleList []models.Article
	//
	//models.DB.Find(&articleList)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"list": articleList,
	//})

	//c.String(200, "文章列表--")

	//var articleList []models.Article
	//
	//models.DB.Preload("ArticleCate").Find(&articleList)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"list": articleList,
	//})

	var articleCateList []models.ArticleCate

	models.DB.Preload("Article").Find(&articleCateList)

	c.JSON(http.StatusOK, gin.H{
		"list": articleCateList,
	})

}

func (con ArticleController) Add(c *gin.Context) {
	c.String(200, "文章列表add--")
}

func (con ArticleController) Edit(c *gin.Context) {
	c.String(200, "文章列表edit--")
}
