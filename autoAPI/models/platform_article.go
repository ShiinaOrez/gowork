package models

type PlatformArticleModel struct {
	BaseModel
	Title string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	Date string `json:"date" gorm:"column:date"`
}
