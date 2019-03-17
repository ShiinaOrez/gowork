package models

type MissionFeedbackModel struct {
	BaseModel
	Missionid uint64 `json:"missionid" gorm:"column:missionid"`
	UsersMissionid uint64 `json:"users_missionid" gorm:"column:users_missionid"`
	Userid uint64 `json:"userid" gorm:"column:userid"`
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	Content string `json:"content" gorm:"column:content"`
	Score uint64 `json:"score" gorm:"column:score"`  // 0~100分
}
