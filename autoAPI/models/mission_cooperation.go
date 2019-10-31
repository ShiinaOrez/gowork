package models

type MissionCooperationModel struct {
	BaseModel
	Missionid          uint64 `json:"missionid" gorm:"column:missionid"`
	MainCompanyid      uint64 `json:"main_companyid" gorm:"column:main_companyid"`
	FriendCompanyid    uint64 `json:"friend_companyid" gorm:"column:friend_companyid"`
	FeedbackPermission bool   `json:"feedback_permission" gorm:"column:feedback_permission"`
}
