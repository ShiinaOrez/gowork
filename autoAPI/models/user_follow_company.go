package models

type UserFollowCompanyModel struct {
	BaseModel
	Userid uint64 `json:"userid" gorm:"column:userid"`
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
}