package models

type ComplaintModel struct {
	BaseModel
	Userid    uint64 `json:"userid" gorm:"column:userid"`
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	Content   string `json:"content" gorm:"column:content"`
	Material  string `json:"material" gorm:"column:material"`
}
