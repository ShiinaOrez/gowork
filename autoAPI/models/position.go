package models

type PositionModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name"`
	Companyid uint64 `json:"companyid" gorm:"column:companyid"`
	Description string `json:"description" gorm:"column:description"`
	Salary string `json:"salary" gorm:"column:salary"`
	PersonNumber uint64 `json:"person_number" gorm:"column:person_number"`
	Experience string `json:"experience" gorm:"column:experience"`
	Education string `json:"education" gorm:"column:education"`
	Treatment string `json:"treatment" gorm:"column:treatment"`
	PositionInfo string `json:"position_info" gorm:"column:position_info"`
}