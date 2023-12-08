package models

type Nav struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	Status int    `json:"status"`
	Sort   int    `json:"sort"`
}

func (Nav) TableName() string {
	return "nav"
}
