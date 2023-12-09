package models

type LessonStudent struct {
	LessonId  int `json:"lesson_id"`
	StudentId int `json:"student_id"`
}

func (LessonStudent) TableName() string {
	return "lesson_student"
}
