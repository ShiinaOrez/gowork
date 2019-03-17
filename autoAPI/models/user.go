package models

// User represents a registered user.
type UserModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url"`
	Sex string `json:"sex" gorm:"column:sex"`
	BirthYear uint64 `json:"birth_year" gorm:"column:birth_year"`
	BirthMonth uint64 `json:"birth_month" gorm:"column:birth_month"`
	Location string `json:"location" gorm:"column:location"`
	University string `json:"university" gorm:"column:university"`
	EnrollYear string `json:"enroll_year" gorm:"column:enroll_year"`
	GraduateYear string `json:"graduate_year" gorm:"column:graduate_year"`
	Education string `json:"education" gorm:"column:education"`
	Major string `json:"major" gorm:"column:major"`
	Status string `json:"status" gorm:"column:status"`
	Intro string `json:"intro" gorm:"column:intro"`
	Wechat string `json:"wechat" gorm:"column:wechat"`
	Level uint64 `json:"level" gorm:"column:level"`
	LastActivity string `json:"last_activity" gorm:"column:last_activity"`
	CanBeSearch bool `json:"canbesearch" gorm:"column:canbesearch"`
	CanSendMail bool `json:"cansendmail" gorm:"column:cansendmail"`
}