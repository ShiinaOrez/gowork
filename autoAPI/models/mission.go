package models

type MissionModel struct {
	BaseModel
	Name          string `json:"name" gorm:"column:name"`
	Checked       bool   `json:"checked" gorm:"column:checked"`
	Companyid     uint64 `json:"companyid" gorm:"column:companyid"`
	ReleaseDate   string `json:"release_date" gorm:"column:release_date"`
	SignUpEndDate string `json:"sign_up_end_date" gorm:"column:sign_up_end_date"`
	UploadEndDate string `json:"upload_end_date" gorm:"column:upload_end_date"`
	Detail        string `json:"detail" gorm:"column:detail"`
	Attachment    string `json:"attachment" gorm:"column:attachment"`
	PersonLimit   uint64 `json:"person_limit" gorm:"column:person_limit"`
	Online        bool   `json:"online" gorm:"column:online"`
	Location      string `json:"location" gorm:"column:location"`
	Level         uint64 `json:"level" gorm:"column:level"`
}
