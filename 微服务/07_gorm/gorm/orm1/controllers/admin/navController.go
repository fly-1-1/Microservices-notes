package admin

import (
	"github.com/gin-gonic/gin"
)

type NavController struct {
}

func (con NavController) Index(c *gin.Context) {

	//查询全部数据
	//navlist := []models.Nav{}
	//models.DB.Find(&navlist)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navlist,
	//})

	//nav := models.Nav{Id: 21}
	//models.DB.Find(&nav)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": nav,
	//})

	//var navList []models.Nav
	//models.DB.Where("id > 3").Find(&navList)
	//var id []int
	//for _, n := range navList {
	//	id = append(id, n.Id)
	//}
	//fmt.Println(id)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Where("id > 3 and id <9").Find(&navList)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	////models.DB.Where("id in (3,4,6)").Find(&navList)
	//models.DB.Where("id in (?)", []int{3, 5, 6}).Find(&navList)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Where("title like ? ", "%会%").Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Where("id between ? and ? ", 3, 9).Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Where("id = ? or id > ? ", 3, 9).Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Where("id = ?", 2).Or("id = ?", 3).Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Select("id,title").Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Order("id desc").Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Order("id desc").Limit(2).Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//models.DB.Order("id desc").Offset(2).Limit(2).Find(&navList)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": navList,
	//})

	//var navList []models.Nav
	//var count int64
	//models.DB.Find(&navList).Count(&count)
	//c.JSON(http.StatusOK, gin.H{
	//	"result": count,
	//})

	/***************************************************************************************/
	//sql 原生

	//models.DB.Exec("delete from user where id = ?", 5)

	//Scan 赋值

}
