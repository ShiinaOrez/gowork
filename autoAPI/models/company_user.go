package models

type CompanyUserModel struct {
	BaseModel
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	Name string `json:"name" gorm:"column:name"`
}
