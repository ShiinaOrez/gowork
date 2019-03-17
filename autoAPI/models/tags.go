package models

type TagsModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name"`
	Categoryid uint64 `json:"categoryid" gorm:"column:categoryid"`
	Checked bool `json:"checked" gorm:"column:checked"`
}