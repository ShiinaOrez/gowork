package models

type CompanyModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name"`
	Tel string `json:"tel" gorm:"column:tel"`
	Mail string `json:"mail" gorm:"column:mail"`
	Wechat string `json:"wechat" gorm:"column:wechat"`
	Logo string `json:"logo" gorm:"column:logo"`
	Nature string `json:"nature" gorm:"column:nature"`
	Intro string `json:"intro" gorm:"column:intro"`
	BusinessLicenseUrl string `json:"business_license_url" gorm:"column:business_license_url"`
	BusinessLicenseLocation string `json:"business_license_location" gorm:"column:business_license_location"`
	OrganizationCodeUrl string `json:"organization_code_url" gorm:"column:organization_code_url"`
	LegalBodyName string `json:"legal_body_name" gorm:"column:legal_body_name"`
	LegalBodyIdNumber string `json:"legal_body_id_number" gorm:"column:legal_body_id_number"`
	ContactName string `json:"contact_name" gorm:"column:contact_name"`
	ContactPhoneNumber string `json:"contact_phone_number" gorm:"column:contact_phone_number"`
	Level uint64 `json:"level" gorm:"column:level"`
	AllowMissionReuse bool `json:"allow_mission_reuse" gorm:"column:allow_mission_reuse"`
	AllowSendMail bool `json:"allow_send_mail" gorm:"column:allow_send_mail"`
	AddressProvince string `json:"address_province" gorm:"column:address_province"`
	AddressCity string `json:"address_city" gorm:"column:address_city"`
	AddressCounty string `json:"address_county" gorm:"column:address_county"`
	AddressDetail string `json:"address_detail" gorm:"column:address_detail"`
	Checked bool `json:"checked" gorm:"column:checked"`
}