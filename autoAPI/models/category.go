package models

type CategoryModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name"`
}
