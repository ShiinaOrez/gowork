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

type TagForUserModelID int

type AuthInitTags struct {
	TagsList []TagForUserModelID `json:"tags_list"`
}

type ResultOfUserLogin struct {
	UserID int     `json:"user_id"`
	Token  string  `json:"token"`
}

type ResultOfCompanySignUp struct {
	CompanyID int `json:"company_id"`
}

type SignUpError struct {
	ErrorCode int `json:"error_code"`
	Msg       string `json:"msg"`
}

type PostConfirmInformationError struct {
	ErrorCode int `json:"error_code"`
	Msg       string `json:"msg"`
}

type CompanyInitTags struct {
	TagsList []TagForCompanyModel `json:"tags_list"`
}

type ResultOfReviewerSignUp struct {
	ReviewerID int `json:"reviewer_id"`
}

type ResultOfCompanyLogin struct {
	CompanyID int     `json:"company_id"`
	Token     string  `json:"token"` 
}

type CompanyLoginError struct {
	ErrorCode int     `json:"error_code"`
	Msg       string  `json:"msg"`
}

type ResultOfReviewerLogin struct {
	CompanyID int     `json:"company_id"`
	Token     string  `json:"token"`
}

type ReviewerLoginError struct {
	ErrorCode int     `json:"error_code"`
	Msg       string  `json:"msg"`
}

type MissionEntityToShow struct {
	MissionID     int         `json:"mission_id"`
	Statu         bool        `json:"statu"`
}

type GetUserInformation struct {
	AvatarURL     string     `json:"avatar_url"`
	Name          string     `json:"name"`
	Status        string     `json:"status"`
	Age           int        `json:"age"`
	University    string     `json:"university"`
	Attention     bool       `json:"attention"`
	UserTagsList  []TagForUserModel `json:"user_tags_list"`
	Education     string     `json:"education"`
	Tel           string     `json:"tel"`
	EmailAddress  string     `json:"email_address"`
	StarLevel     int        `json:"star_level"`
	Major         string     `json:"major"`
	EduTimeRange  string     `json:"edu_time_range"`
	City          string     `json:"city"`
	Intro         string     `json:"intro"`
	MissionsList  []MissionEntityToShow `json:"missions_list"`
}

type GetUserBaseInformation struct {
	AvatarURL     string     `json:"avatar_url"`
	Name          string     `json:"name"`
	Sex           bool       `json:"sex"`
	EmailAddress  string     `json:"email_address"`
	Tel           string     `json:"tel"`
    BirthYear     int        `json:"birth_year"`
	BirthMonth    int        `json:"birth_month"`
	City          string     `json:"city"`
	University    string     `json:"university"`
	EnrollYear    int        `json:"enroll_year"`
	Education     string     `json:"education"`
	Major         string     `json:"major"`
	Status        string     `json:"status"`
	UserTagsList  []TagForUserModel `json:"user_tags_list"`
	Intro         string     `json:"intro"`
}

type Mission struct {
	ID            int        `json:"id"`
	MissionName   string     `json:"mission_name"`
	Status        bool       `json:"status"`
	CompanyName   string     `json:"company_name"`
	CompanyAvatar string     `json:"company_avatar"`
	StartTime     string     `json:"start_time"`
	EndTime       string     `json:"end_time"`
}

type UserMissionInformation struct {
	MissionList   []Mission  `json:"mission_list"`
	AverageStar   float32    `json:"average_star"`
}

type CompanySimpleInformation struct {
	CompanyID     int        `json:"company_id"`
	CompanyName   string     `json:"company_name"`
	CompanyAvatar string     `json:"company_avatar"`
}

type UserAttentionCompanyList struct {
	CompanyList   []CompanySimpleInformation  `json:"company_list"`
	Length        int                         `json:"length"`
}

type EditSettingBody struct {
	Name          string     `json:"name"`
	Sex           bool       `json:"sex"`
	BirthYear     int        `json:"birth_year"`
	BirthMonth    int        `json:"birth_month"`
	City          string     `json:"city"`
	University    string     `json:"university"`
	EnrollYear    int        `json:"enroll_year"`
	Education     string     `json:"education"`
	Major         string     `json:"major"`
	Status        string     `json:"status"`
	UserTagsList  []TagForUserModel `json:"user_tags_list"`
	Intro         string     `json:"intro"`
}

type UserPlatformSetting struct {

}

type CompanySuitability struct {
	CompanyID     int        `json:"company_id"`
	CompanyName   string     `json:"company_name"`
	CompanyAvatar string     `json:"company_avatar"`
	Suitability   float32    `json:"suitability"`
}

type GetCompanySuitabilityList struct {
	RespBody      []CompanySuitability  `json:"resp_body"`
	Length        int        `json:"length"`
}

type Tag struct {
	TagID         int        `json:"tag_id"`
	TagName       string     `json:"tag_name"`
}

type TagList struct {
	CategoryID    int        `json:"category_id"`
	CategoryName  string     `json:"category_name"`
	TagsList      []Tag      `json:"tags_list"`
}