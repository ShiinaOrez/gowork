package models

type MissionListMissionsModel struct {
	BaseModel
	MissionListid uint64 `json:"mission_listid" gorm:"column:mission_listid"`
	Missionid     uint64 `json:"missionid" gorm:"column:missionid"`
}
