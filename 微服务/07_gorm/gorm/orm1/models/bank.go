package models

type Bank struct {
	Id       int
	Username string
	Balance  float32
}

func (Bank) TableName() string {
	return "bank"
}
