package models

type MissionListModel struct {
	BaseModel
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	MissionListName string `json:"mission_list_name" gorm:"column:mission_list_name"`
}
