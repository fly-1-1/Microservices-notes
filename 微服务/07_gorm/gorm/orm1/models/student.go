package models

type Student struct {
	Id       int
	Number   string
	Password string
	ClassId  int
	Name     string
	Lesson   []Lesson `gorm:"many2many:lesson_student;"`
}

func (Student) TableName() string {
	return "student"
}
