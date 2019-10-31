package models

type UserMissionListModel struct {
	BaseModel
	MissionListid uint64 `json:"mission_listid" gorm:"column:mission_listid"`
	Userid        uint64 `json:"userid" gorm:"column:userid"`
}
