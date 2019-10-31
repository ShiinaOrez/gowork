package models

type TagForMissionModel struct {
	BaseModel
	Tagid     uint64 `json:"tagid" gorm:"column:tagid"`
	Missionid uint64 `json:"missionid" gorm:"column:missionid"`
}
