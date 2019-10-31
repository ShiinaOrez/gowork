package models

type AccountModel struct {
	BaseModel
	Email        string `json:"email" gorm:"column:email" binding:"required" validate:"min=1,max=30"`
	Tel          string `json:"tel" gorm:"column:tel" binding:"required" validate:"min=1,max=20"`
	PasswordHash string `json:"-" gorm:"column:password_hash" binding:"required"`
	Role         string `json:"role" gorm:"column:role" binding:"required" validate:"min=1,max=10"`
	Uid          uint64 `json:"uid" gorm:"column:uid" binding:"required"`
}
