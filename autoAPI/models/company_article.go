package models

type CompanyArticleModel struct {
	BaseModel
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	Title string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
	Date string `json:"date" gorm:"column:date"`
}