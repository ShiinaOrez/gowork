package models

type CompanyPlatformMessageModel struct {
	BaseModel
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	Content string `json:"content" gorm:"column:content"`
	Readed uint64 `json:"readed" gorm:"column:readed"`
}
