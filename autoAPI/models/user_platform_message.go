package models

type UserPlatformMessageModel struct {
	BaseModel
	Userid  uint64 `json:"userid" gorm:"column:userid"`
	Content string `json:"content" gorm:"column:content"`
	Readed  bool   `json:"readed" gorm:"column:readed"`
}
