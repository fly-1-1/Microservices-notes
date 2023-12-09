package models

type Lesson struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Student []Student `gorm:"many2many:lesson_student"`
}

func (Lesson) TableName() string {
	return "lesson"
}
