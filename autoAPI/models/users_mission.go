package models


type UsersMissionModel struct {
	BaseModel
	Userid uint64 `json:"userid" gorm:"column:userid"`
	Missionid uint64 `json:"missionid" gorm:"column:missionid"`
	Status uint64 `json:"status" gorm:"column:status"`
	Intro string `json:"Intro" gorm:"column:Intro"`
	Attachment string `json:"attachment" gorm:"column:attachment"`
}