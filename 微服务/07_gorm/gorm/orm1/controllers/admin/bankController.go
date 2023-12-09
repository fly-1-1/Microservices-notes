package admin

import (
	"github.com/gin-gonic/gin"
	"orm1/models"
)

type BankController struct {
	BaseController
}

func (con BankController) Index(c *gin.Context) {

	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			con.error(c)
		}
	}()
	//panic("异常")
	//张三的账户减去100元
	u1 := models.Bank{Id: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance - 100
	err := tx.Save(&u1).Error
	if err != nil {
		tx.Rollback()
		con.error(c)
		return
	}

	//李四的账户增加100元
	u2 := models.Bank{Id: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance + 100
	err = tx.Save(&u2).Error
	if err != nil {
		tx.Rollback()
		con.error(c)
		return
	}

	tx.Commit()
	con.success(c)
}
