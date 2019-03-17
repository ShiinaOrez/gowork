package models

type TagForCompanyModel struct {
	BaseModel
	Tagid uint64 `json:"tagid" gorm:"column:tagid"`
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
}