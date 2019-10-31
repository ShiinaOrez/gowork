package models

type CompanyToUserInfoModel struct {
	BaseModel
	Userid    uint64 `json:"userid" gorm:"column:userid"`
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	Content   string `json:"content" gorm:"column:content"`
	Readed    bool   `json:"readed" gorm:"column:readed"`
}
