package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"orm1/models"
)

type StudentController struct {
	BaseController
}

func (con StudentController) Index(c *gin.Context) {

	//var studentList []models.Student
	//
	//models.DB.Find(&studentList)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": studentList,
	//})

	//var lessonList []models.Lesson
	//
	//models.DB.Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": lessonList,
	//})

	//var studentList []models.Student
	//models.DB.Preload("Lesson").Find(&studentList)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": studentList,
	//})

	//var student models.Student
	//models.DB.Preload("Lesson").Where("name = ?", "李四").Find(&student)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": student,
	//})

	//课程被哪些学生选修了
	//var lessonList []models.Lesson
	//models.DB.Preload("Student").Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": lessonList,
	//})

	//计算机网络被哪些学生选修了
	//var lesson models.Lesson
	//models.DB.Preload("Student").Where("name = ?", "计算机网络").Find(&lesson)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": lesson,
	//})

	//指定条件
	//var lessonList []models.Lesson
	//models.DB.Preload("Student").Offset(2).Limit(2).Order("id desc").Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": lessonList,
	//})

	//var lessonList []models.Lesson
	////models.DB.Preload("Student", "name != ?", "张三").Find(&lessonList)
	//models.DB.Preload("Student", "id not in (1,2,3,4) ").Find(&lessonList)
	//c.JSON(http.StatusOK, gin.H{
	//	"list": lessonList,
	//})

	//自定义预加载sql
	var lessonList []models.Lesson
	models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Where("id >3").Order("student.id desc")
	}).Find(&lessonList)
	c.JSON(http.StatusOK, gin.H{
		"list": lessonList,
	})

}
