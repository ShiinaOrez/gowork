package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type CompanyController struct {
	beego.Controller
}

// @Title CompanyInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Cid/info [post]
func (u *CompanyController) CompanyInfo() {
}

// @Title CompanyBaseInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /info [post]
func (u *CompanyController) CompanyBaseInfo() {
}

// @Title CompanySearchByIntelligent
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /search/intelligent [post]
func (u *CompanyController) CompanySearchIntelligent() {
}

// @Title GetCompanyAttentionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/list [post]
func (u *CompanyController) CompanyAttentionList() {
}

// @Title CompanyPayAttentionToSomebody
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention [post]
func (u *CompanyController) CompanyPayAttention() {
}

// @Title CompanyCancelAttentionSomebody
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention [delete]
func (u *CompanyController) CompanyCancelAttention() {
}

// @Title GetCompanyPositionListByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Cid/position/list [post]
func (u *CompanyController) GetCompanyPositionList() {
}

// @Title CreateAndPostPosition
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /position [post]
func (u *CompanyController) CreateCompanyPosition() {
}

// @Title EditCompanyPositionByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /position/:Pid [put]
func (u *CompanyController) EditCompanyPosition() {
}

// @Title DeleteCompanyPositionByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /position/:Pid [delete]
func (u *CompanyController) DeleteCompanyPosition() {
}

// @Title GetCompanyMissionListByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /mission/list [get]
func (u *CompanyController) GetCompanyMissionList() {
}

// @Title UpdateCompanyLogo
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /logo [put]
func (u *CompanyController) UpdateCompanyLogo() {
}

// @Title EditCompanyInformation
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /info [put]
func (u *CompanyController) EditCompanyInfo() {
}

// @Title UpdateCompanyTelephone
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /tel [put]
func (u *CompanyController) UpdateCompanyTel() {
}

// @Title UpdateCompanyEmailAddress
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /email [post]
func (u *CompanyController) UpdateCompanyEmail() {
}

// @Title ConfirmCompanyEmailAddress
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /email/confirm [post]
func (u *CompanyController) ConfirmCompanyEmail() {
}

// @Title CompanySearchByFastWay
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /search/fast [post]
func (u *CompanyController) CompanySearchFast() {
}

// @Title CreateAndPostCompanyOffer
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /offer [post]
func (u *CompanyController) CreateCompanyOffer() {
}

// @Title GetCompanyOfferInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /offer/:Oid [get]
func (u *CompanyController) GetCompanyOfferInfo() {
}

// @Title DeleteCompanyOfferByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /offer/:Oid [delete]
func (u *CompanyController) DeleteCompanyOffer() {
}

// @Title CreateAndPostMissionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /missionList [post]
func (u *CompanyController) CreateCompanyMissionList() {
}

// @Title CreateAndPostCompanyCategory
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /category [post]
func (u *CompanyController) CreateCompanyCategory() {
}

// @Title AddTagToCompanyCategory
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /category/:CateID/tags [post]
func (u *CompanyController) CompanyCategoryAddTag() {
}

// @Title ResumeToCompanyPosition
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /position/:Pid/resume [post]
func (u *CompanyController) ResumeToCompanyPosition() {
}

// @Title SearchCompanyPositionByNormalWay
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /position/:Pid/search [post]
func (u *CompanyController) CompanyPositionSearch() {
}