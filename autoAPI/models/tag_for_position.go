package models

type TagForPositionModel struct {
	BaseModel
	Tagid      uint64 `json:"tagid" gorm:"column:tagid"`
	Positionid uint64 `json:"positionid" gorm:"column:positionid"`
}
