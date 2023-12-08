package models

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// TableName 配置数据库表名称
func (User) TableName() string {
	return "user"
}
