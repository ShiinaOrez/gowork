package models

type ReviewerModel struct {
	BaseModel
	Level uint64 `json:"level" gorm:"column:level"`
	Name  string `json:"name" gorm:"column:name"`
	Intro string `json:"intro" gorm:"column:intro"`
}
