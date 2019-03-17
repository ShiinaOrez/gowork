package models

type TagForUserModel struct {
	BaseModel
	Tagid uint64 `json:"tagid" gorm:"column:tagid"`
	Userid uint64 `json:"userid" gorm:"column:userid"`
	Value uint64 `json:"value" gorm:"column:value"`
}