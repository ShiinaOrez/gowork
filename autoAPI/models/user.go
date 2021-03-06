package models

// User represents a registered user.
type UserModel struct {
	BaseModel
	Name         string `json:"name" gorm:"column:name"`
	AvatarUrl    string `json:"avatar_url" gorm:"column:avatar_url"`
	Sex          string `json:"sex" gorm:"column:sex"`
	BirthYear    uint64 `json:"birth_year" gorm:"column:birth_year"`
	BirthMonth   uint64 `json:"birth_month" gorm:"column:birth_month"`
	Location     string `json:"location" gorm:"column:location"`
	University   string `json:"university" gorm:"column:university"`
	EnrollYear   string `json:"enroll_year" gorm:"column:enroll_year"`
	GraduateYear string `json:"graduate_year" gorm:"column:graduate_year"`
	Education    string `json:"education" gorm:"column:education"`
	Major        string `json:"major" gorm:"column:major"`
	Status       string `json:"status" gorm:"column:status"`
	Intro        string `json:"intro" gorm:"column:intro"`
	Wechat       string `json:"wechat" gorm:"column:wechat"`
	Level        uint64 `json:"level" gorm:"column:level"`
	LastActivity string `json:"last_activity" gorm:"column:last_activity"`
	CanBeSearch  bool   `json:"canbesearch" gorm:"column:canbesearch"`
	CanSendMail  bool   `json:"cansendmail" gorm:"column:cansendmail"`
}

type TagForUserModelID int

type AuthInitTags struct {
	TagsList []TagForUserModelID `json:"tags_list"`
}

type AuthEmailSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResultOfUserLogin struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

type ResultOfCompanySignUp struct {
	CompanyID int `json:"company_id"`
}

type SignUpError struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
}

type PostConfirmInformationError struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
}

type CompanyInitTags struct {
	TagsList []TagForCompanyModel `json:"tags_list"`
}

type ResultOfReviewerSignUp struct {
	ReviewerID int `json:"reviewer_id"`
}

type ResultOfCompanyLogin struct {
	CompanyID int    `json:"company_id"`
	Token     string `json:"token"`
}

type CompanyLoginError struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
}

type ResultOfReviewerLogin struct {
	ReviewerID int    `json:"reviewer_id"`
	Token      string `json:"token"`
}

type ReviewerLoginError struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
}

type MissionInfo struct {
	MissionID   int    `json:"mission_id"`
	Status      bool   `json:"status"` // 该用户是否已经完成
	MissionName string `json:"mission_name"`
	Deadline    string `json:"deadline"`
	CompanyLogo string `json:"company_logo"`
	CompanyName string `json:"company_name"`
	CompanyID   int    `json:"company_id"`
}

type MissionEntityToShow struct {
	MissionID int  `json:"mission_id"`
	Statu     bool `json:"statu"`
}

type TagForGetUserInformation struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
	Value   int    `json:"value"`
}

type GetUserInformation struct {
	AvatarURL    string                     `json:"avatar_url"`
	Name         string                     `json:"name"`
	Status       string                     `json:"status"`
	Age          int                        `json:"age"`
	University   string                     `json:"university"`
	Attention    bool                       `json:"attention"`
	UserTagsList []TagForGetUserInformation `json:"user_tags_list"`
	Education    string                     `json:"education"`
	Tel          string                     `json:"tel"`
	EmailAddress string                     `json:"email_address"`
	StarLevel    int                        `json:"star_level"`
	Major        string                     `json:"major"`
	EduTimeRange string                     `json:"edu_time_range"`
	City         string                     `json:"city"`
	Intro        string                     `json:"intro"`
	MissionsList []MissionInfo              `json:"missions_list"`
}

type GetUserInformationInCompanyView struct {
	IsFollow     bool                       `json:"is_follow"`
	AvatarURL    string                     `json:"avatar_url"`
	Name         string                     `json:"name"`
	Status       string                     `json:"status"`
	Age          int                        `json:"age"`
	University   string                     `json:"university"`
	Attention    bool                       `json:"attention"`
	UserTagsList []TagForGetUserInformation `json:"user_tags_list"`
	Education    string                     `json:"education"`
	Tel          string                     `json:"tel"`
	EmailAddress string                     `json:"email_address"`
	StarLevel    int                        `json:"star_level"`
	Major        string                     `json:"major"`
	EduTimeRange string                     `json:"edu_time_range"`
	City         string                     `json:"city"`
	Intro        string                     `json:"intro"`
	MissionsList []MissionInfo              `json:"missions_list"`
}

type GetUserBaseInformation struct {
	AvatarURL    string                     `json:"avatar_url"`
	Name         string                     `json:"name"`
	Sex          bool                       `json:"sex"`
	EmailAddress string                     `json:"email_address"`
	Tel          string                     `json:"tel"`
	BirthYear    int                        `json:"birth_year"`
	BirthMonth   int                        `json:"birth_month"`
	City         string                     `json:"city"`
	University   string                     `json:"university"`
	EnrollYear   int                        `json:"enroll_year"`
	GraduateYear int                        `json:"graduate_year"`
	Education    string                     `json:"education"`
	Major        string                     `json:"major"`
	Status       string                     `json:"status"`
	UserTagsList []TagForGetUserInformation `json:"user_tags_list"`
	Intro        string                     `json:"intro"`
}

type Mission struct {
	ID            int    `json:"id"`
	MissionName   string `json:"mission_name"`
	Status        bool   `json:"status"`
	CompanyName   string `json:"company_name"`
	CompanyAvatar string `json:"company_avatar"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
}

type UserMissionInformation struct {
	MissionList []Mission `json:"mission_list"`
	AverageStar float32   `json:"average_star"`
}

type CompanySimpleInformation struct {
	CompanyID     int    `json:"company_id"`
	CompanyName   string `json:"company_name"`
	CompanyAvatar string `json:"company_avatar"`
}

type UserAttentionCompanyList struct {
	List  []CompanySimpleInformation `json:"company_list"`
	Total int                        `json:"total"`
}

type EditSettingBody struct {
	Name         string `json:"name"`
	Sex          bool   `json:"sex"`
	BirthYear    int    `json:"birth_year"`
	BirthMonth   int    `json:"birth_month"`
	City         string `json:"city"`
	University   string `json:"university"`
	EnrollYear   int    `json:"enroll_year"`
	GraduateYear int    `json:"graduate_year"`
	Education    string `json:"education"`
	Major        string `json:"major"`
	Status       string `json:"status"`
	UserTagsList []int  `json:"user_tags_list"`
	Intro        string `json:"intro"`
}

type UserAccountSetting struct {
	Name         string `json:"name"`
	Sex          bool   `json:"sex"`
	BirthMonth   int    `json:"birth_month"`
	BirthYear    int    `json:"birth_year"`
	City         string `json:"city"`
	University   string `json:"university"`
	EnrollYear   int    `json:"enroll_year"`
	GraduateYear int    `json:"graduate_year"`
	Education    string `json:"education"`
	Major        string `json:"major"`
	Status       string `json:"status"`
	TagList      []Tag  `json:"tag_list"`
	Intro        string `json:"intro"`
}

type UserPersonalSetting struct {
	CanBeSearch bool `json:"can_be_search"`
	CanSendMail bool `json:"can_send_mail"`
}

type CompanySuitability struct {
	CompanyID     int     `json:"company_id"`
	CompanyName   string  `json:"company_name"`
	CompanyAvatar string  `json:"company_avatar"`
	Suitability   float32 `json:"suitability"`
}

type GetCompanySuitabilityList struct {
	RespBody []CompanySuitability `json:"resp_body"`
	Pagination
}

type Tag struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

type TagList struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	TagsList     []Tag  `json:"list"`
}

type Elite struct {
	EliteID         int     `json:"elite_id"`
	EliteAvatar     string  `json:"elite_avatar"`
	EliteName       string  `json:"elite_name"`
	EliteStatus     string  `json:"elite_status"`
	EliteAge        int     `json:"elite_age"`
	EliteUniversity string  `json:"elite_university"`
	EliteTags       []Tag   `json:"elite_tags"`
	EliteEdu        string  `json:"elite_edu"`
	EliteTel        string  `json:"elite_tel"`
	EliteEmail      string  `json:"elite_email"`
	EliteStarLevel  float32 `json:"elite_star_level"`
	IsAttention     bool    `json:"is_attention"`
}

type EliteList struct {
	List []*Elite `json:"list"`
}

type MissionInformation struct {
	MissionName      string `json:"mission_name"`
	MissionStartTime string `json:"mission_start_time"`
	MissionEndTime   string `json:"mission_end_time"`
	MissionOnline    bool   `json:"mission_online"`
	MissionMember    int    `json:"mission_member"`
	MissionTagList   []Tag  `json:"mission_tag_list"`
	MissionIntro     string `json:"mission_intro"`
	MissionDetail    string `json:"mission_detail"`
	MissionFile      string `json:"mission_file"`
}

type TagPayload struct {
	TagID      int `json:"tag_id"`
	CategoryID int `json:"category_id"`
}

type SearchPayload struct {
	Pattern string       `json:"pattern"`
	TagList []TagPayload `json:"tag_list"`
}

type MissionListSearchPayload struct {
	Pattern string       `json:"pattern"`
	TagList []TagPayload `json:"tag_list"`
}

type MissionEntity struct {
	MissionID   int    `json:"mission_id"`
	MissionName string `json:"mission_name"`
	IsOnline    bool   `json:"is_online"`
	TagList     []Tag  `json:"tag_list"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	JoinEndTime string `json:"join_end_time"`
	CompanyID   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
}

type MissionList struct {
	List  []MissionEntity `json:"list"`
	Total int             `json:"total"`
}

type DoMissionPayload struct {
	Intro   string `json:"intro"`
	FileURL string `json:"file_url"`
}

type DoMissionResponse struct {
	UserID     int                      `json:"uid"`
	Username   string                   `json:"username"`
	UserAvatar string                   `json:"user_avatar"`
	UploadTime string                   `json:"upload_time"`
	IfFeedback bool                     `json:"if_feedback"`
	Intro      string                   `json:"intro"`
	FileURL    string                   `json:"file_url"`
	Feedback   CompanyDoMissionFeedback `json:"feedback"`
}

type DoMissionResponseList struct {
	List  []DoMissionResponse `json:"list"`
	Total int                 `json:"total"`
}

type CompanyDoMissionFeedback struct {
	Feedback  string  `json:"feedback"`
	StarLevel float32 `json:"star_level"`
}

type CreateMissionPayload struct {
	MissionName         string       `json:"mission_name"`
	TogetherCompanyList []int        `json:"together_company_list"`
	JoinEndTime         string       `json:"join_end_time"`
	EndTime             string       `json:"end_time"`
	IsOnline            bool         `json:"is_online"`
	MemberCount         int          `json:"member_count"`
	MissionAddress      string       `json:"address"`
	MissionDetail       string       `json:"mission_detail"`
	MissionLists        []int        `json:"mission_lists"`
	FileURL             string       `json:"file_url"`
	TagList             []TagPayload `json:"tag_list"`
}

type MissionStatistics struct {
	Member int `json:"member"`
}

type MissionInMissionList struct {
	MissionName        string `json:"mission_name"`
	MissionStatus      bool   `json:"mission_status"`
	MissionEndTime     string `json:"mission_end_time"`
	MissionMemberCount int    `json:"mission_member_count"`
	MissionMemberLimit int    `json:"mission_member_limit"`
	MissionStartTime   string `json:"mission_start_time"`
}

type MissionListInformation struct {
	MissionListName       string                 `json:"mission_list_name"`
	MissionListCreateTime string                 `json:"mission_list_create_time"`
	MissionListIntro      string                 `json:"mission_list_intro"`
	List                  []MissionInMissionList `json:"mission_list"`
	Total                 int                    `json:"total"`
}

type MissionListEntity struct {
	MissionListID         int    `json:"mission_list_id"`
	MissionListName       string `json:"mission_list_name"`
	MissionListCreateTime string `json:"mission_list_create_time"`
	CompanyID             int    `json:"company_id"`
	CompanyName           string `json:"company_name"`
	CompanyLogo           string `json:"company_logo"`
	MissionCount          int    `json:"mission_count"`
	FinishMissionCount    int    `json:"finish_mission_count"`
}

type MissionListList struct {
	List  []MissionListEntity `json:"list"`
	Total int                 `json:"total"`
}

type MissionListSearchEntity struct {
	MissionListID         int    `json:"mission_list_id"`
	MissionListName       string `json:"mission_list_name"`
	MissionListCreateTime string `json:"mission_list_create_time"`
	CompanyID             int    `json:"company_id"`
	CompanyName           string `json:"company_name"`
	CompanyLogo           string `json:"company_logo"`
}

type MissionListSearchList struct {
	List  []MissionListSearchEntity `json:"list"`
	Total int                       `json:"total"`
}

type CompanyInformation struct {
	CompanyAvatar  string `json:"company_avatar"`
	CompanyName    string `json:"company_name"`
	CompanyIntro   string `json:"company_intro"`
	CompanyEmail   string `json:"company_email"`
	CompanyTel     string `json:"company_tel"`
	CompanyAddress string `json:"company_address"`
	TagList        []Tag  `json:"tag_list"`
}

type CompanyInfoModify struct {
	CompanyName               string `json:"company_name"`
	CompanyIntro              string `json:"company_intro"`
	CompanyAddress            string `json:"company_address"`
	CompanyLegalBodyName      string `json:"company_legal_body_name"`
	CompanyContactName        string `json:"company_contact_name"`
	CompanyContactPhoneNumber string `json:"company_contact_phone_number"`
}

type CompanyLogoModify struct {
	Logo string `json:"logo"`
}

type CompanyPayAttentionBody struct {
	Uid uint64 `json:"uid"`
}

type CompanyCancelAttentionBody struct {
	Uid uint64 `json:"uid"`
}

type CompanyBaseInformation struct {
	CompanyLegalName string `json:"comapany_legal_name"`
	CompanyInformation
}

type CompanySearchList struct {
	Pagination
	CompanyList []CompanySimpleInformation `json:"company_list"`
}

type EliteSimpleInformation struct {
	EliteName   string `json:"elite_name"`
	EliteID     int    `json:"elite_id"`
	EliteAvatar string `json:"elite_avatar"`
	EliteAge    int    `json:"elite_age"`
}

type ElitesList struct {
	Total int                      `json:"total"`
	List  []EliteSimpleInformation `json:"list"`
}

type PositionSimple struct {
	PositionID   int    `json:"position_id"`
	PositionName string `json:"position_name"`
	PersonNumber string `json:"person_number"`
	Education    string `json:"education"`
}

type PositionsList struct {
	Pagination
	PositionList []PositionSimple `json:"position_list"`
}

type PositionPayload struct {
	PositionName         string `json:"position_name"`
	PositionDescription  string `json:"position_description"`
	PositionSalary       string `json:"position_salary"`
	PositionPersonNumber int    `json:"position_person_number"`
	PositionExperience   string `json:"position_experience"`
	PositionEducation    string `json:"position_education"`
	PositionTreatment    string `json:"position_treatment"`
	PositionInformation  string `json:"position_information"`
	BindMissionListID    int    `json:"bind_mission_list_id"`
}

type CompanyVerySimpleInfo struct {
	CompanyID   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
}

type CompanyVerySimpleInfoList struct {
	Pagination
	CompanyList []CompanyVerySimpleInfo `json:"company_list"`
}

type MissionListCreateInformation struct {
	MissionListName  string `json:"mission_list_name"`
	MissionIDList    []int  `json:"mission_id_list"`
	MissionListIntro string `json:"mission_list_intro"`
}

type PositionEntity struct {
	PositionID           int    `json:"position_id"`
	PositionName         string `json:"position_name"`
	CompanyID            int    `json:"company_id"`
	CompanyName          string `json:"company_name"`
	CompanyAvatar        string `json:"company_avatar"`
	PositionPersonNumber int    `json:"position_person_number"`
	PositionEducation    string `json:"position_education"`
}

type Positions struct {
	Pagination
	PositionList []PositionEntity `json:"position_list"`
}

type PlatformMessage struct {
	MessageID   int    `json:"message_id"`
	MessageText string `json:"message_text"`
	MessageTime string `json:"message_time"`
	//	IsRecommand          bool          `json:"is_recommand"`
	//	RecommandType        bool          `json:"recommand_type"`
	//	RecommandID          int           `json:"recommand_id"`
	Readed bool `json:"readed"`
}

type PlatformMessageList struct {
	Pagination
	MessageList []PlatformMessage `json:"message_list"`
}

type FeedbackMessage struct {
	MessageID     int    `json: message_id`
	CompanyID     int    `json:"company_id"`
	CompanyName   string `json:"company_name"`
	CompanyAvatar string `json:"company_avatar"`
	MessageText   string `json:"message_text"`
	Time          string `json:"time"`
	Readed        bool   `json:"readed"`
}

type FeedbackMessageList struct {
	Pagination
	MessageList []FeedbackMessage `json:"message_list"`
}

type MessageEntity struct {
	MessageText   int    `json:"message_text"`
	IsPlatform    bool   `json:"is_platform"`
	IsUser        bool   `json:"is_user"`
	FromID        int    `json:"from_id"`
	FromName      string `json:"from_name"`
	FromAvatar    string `json:"from_avatar"`
	MessageReaded bool   `json:"message_readed"`
	MessageTime   string `json:"message_time"`
}

type CompanyMessage struct {
	MessageID   int    `json:"message_id"`
	IsCompany   bool   `json:"is_company"`
	FromName    string `json:"from_name"`
	FromID      string `json:"from_id"`
	FromAvatar  string `json:"from_avatar"`
	MessageText string `json:"message_text"`
	Readed      bool   `json:"readed"`
	Time        string `json:"time"`
}

type CompanyMissionNumberOverview struct {
	MissionNumber    int `json:"mission_number"`    // 任务总数
	FinishedNumber   int `json:"finished_number"`   // 已结束
	SubmittingNumber int `json:"submitting_number"` // 提交中
	ApplyingNumber   int `json:"applying_number"`   // 申请中
}

type CompanyMessageList struct {
	Total int              `json:"total"`
	List  []CompanyMessage `json:"list"`
}

type Pagination struct {
	PageNow  int  `json:"page_now"`
	HasNext  bool `json:"has_next"`
	PageSize int  `json:"page_size"`
	RowsNum  int  `json:"rows_num"`
}

type ResetPasswordPayload struct {
	EmailAddress string `json:"email_address"`
	Code         string `json:"code`
	NewPassword  string `json:"new_password"`
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type CompanyAllowSendEmailPayload struct {
	IfAllow bool `json:"if_allow"`
}

type CompanyMessagePayload struct {
	Message string `json:"message"`
}

type CreateCompanyCategoryBody struct {
	Name string `json:"name"`
}

type CategoryCreateResponse struct {
	Id uint64 `json:"id"`
}

type CreateTagForCompanyBody struct {
	TagName string `json:"tag_name"`
}

type CreateMissionListResponse struct {
	Id uint64 `json:"id"`
}

type AddMissionToMissionListRequest struct {
	MissionId     uint64 `json:"mission_id"`
	MissionListId uint64 `json:"mission_list_id"`
}

type RemoveMissionInListRequest struct {
	MissionId     uint64 `json:"mission_id"`
	MissionListId uint64 `json:"mission_list_id"`
}

type UpdateTelStep1Payload struct {
	Tel string `json:"tel"`
}

type UpdateTelStep2Payload struct {
	Tel  string `json:"tel"`
	Code string `json:"code"`
}

type UpdateEmailStep1Payload struct {
	Email string `json:"email"`
}

type UpdateEmailStep2Payload struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type CreateTagPayload struct {
	CategoryID uint64 `json:"category_id"`
	TagName    string `json:"tag_name`
}

type CreateTagReturn struct {
	Msg   string `json:"msg"`
	TagID uint64 `json:"tag_id`
}

type UpdateTagPayload struct {
	Name string `json:"name"`
}

type GetTagReturn struct {
	Name       string `json:"name"`
	CategoryID uint64 `json:"category_id"`
}

type CreateCategoryPayload struct {
	Name       string `json:"name"`
	ForElite   bool   `json:"for_elite"`
	ForMission bool   `json:"for_mission"`
	ForCompany bool   `json:"for_company"`
}

type CreateCategoryReturn struct {
	CategoryID uint64 `json:"category_id"`
}

type AddTagsToCategoryPayload struct {
	TagsList []uint64 `json:"tags_list"`
}

type RemoveTagFromCategory struct {
	TagID uint64 `json:"tag_id"`
}

type CategoryAndTags struct {
	CategoryName string `json:"category_name"`
	CategoryID   uint64 `json:"category_id"`
	CategoryTags []Tag  `json:"category_tags"`
}

type CategoryForSearch struct {
	CategoryList []CategoryAndTags `json:"category_list"`
}
