package models

type FeedbackModel struct {
	BaseModel
	Userid uint64 `json:"userid" gorm:"column:userid"`
	Content string `json:"content" gorm:"column:content"`
}