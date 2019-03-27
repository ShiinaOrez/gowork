package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Companies
type CompanyController struct {
	beego.Controller
}

// @Title CompanyInformationByID
// @Description Get company information by id. (information less)
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      int         true        "The company ID for get"
// @Success 200 {object} models.CompanyInformation
// @Failure 401 {string} auth failed
// @Failure 401 {string} user | company not existed
// @router /:Cid/info [get]
func (u *CompanyController) CompanyInfo() {
}

// @Title CompanyBaseInformationByID
// @Description Get company base information.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanyBaseInformation
// @Failure 401 {string} auth failed
// @Failure 401 {string} company not existed
// @router /info [get]
func (u *CompanyController) CompanyBaseInfo() {
}

// @Title CompanySearchByIntelligent
// @Description Search company by intelligent way.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanySearchList
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /search/intelligent [get]
func (u *CompanyController) CompanySearchIntelligent() {
}

// @Title GetCompanyAttentionList
// @Description Get company's attention elites list
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.ElitesList
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @router /attention/list [get]
func (u *CompanyController) CompanyAttentionList() {
}

// @Title CompanyPayAttentionToSomebody
// @Description Pay attention to a elite.
// @Param   Token       header    string      true        "The token to conform"
// @Param   EliteID     body      int         true        "The elite ID for attention"
// @Success 200 {string} pay attention successful!
// @Failure 401 {string} auth failed
// @Failure 402 {string} company | elite not existed
// @Failure 407 {string} already pay attention to it
// @router /attention [post]
func (u *CompanyController) CompanyPayAttention() {
}

// @Title CompanyCancelAttentionSomebody
// @Description Cancel to pay attention to a elite
// @Param   Token       header    string      true        "The token to conform"
// @Param   UiD         path      int         true        "The elite ID for attention"
// @Success 200 {string} cancel successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} company | user not existed
// @Failure 407 {string} never pay attention to it
// @router /attention/:Uid [put]
func (u *CompanyController) CompanyCancelAttention() {
}

// @Title GetCompanyPositionListByID
// @Description Get a company's position list
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      int         true        "The company ID for get position list"
// @Success 200 {object} models.PositionsList
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @router /:Cid/position/list [get]
func (u *CompanyController) GetCompanyPositionList() {
}

// @Title CreateAndPostPosition
// @Description Company create a position.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      object      models.PositionPayload
// @Success 200 {int} positionID
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @router /position [post]
func (u *CompanyController) CreateCompanyPosition() {
}

// @Title EditCompanyPositionByID
// @Description Edit the position information
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      object      models.PositionPayload
// @Success 200 {string} edit successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} company | position not existed
// @router /position/:Pid [put]
func (u *CompanyController) EditCompanyPosition() {
}

// @Title DeleteCompanyPositionByID
// @Description Delete position by ID
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {string} delete successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} company | position not existed
// @router /position/:Pid [delete]
func (u *CompanyController) DeleteCompanyPosition() {
}

// @Title GetCompanyMissionListByID
// @Description Get a company mission list by company ID.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      int         true        "The company ID for get mission list"
// @Success 200 {string} models.MissionList
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @router /:Cid/mission/list [get]
func (u *CompanyController) GetCompanyMissionList() {
}

// @Title UpdateCompanyLogo
// @Description Update company logo.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Logo        body      string      true        "The company new logo."
// @Success 200 {string} new logo URL
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @router /logo [put]
func (u *CompanyController) UpdateCompanyLogo() {
}

// @Title EditCompanyInformation
// @Description Edit company information
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {string} models.CompanyInformation
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @Failure 406 {string} image error
// @router /info [put]
func (u *CompanyController) EditCompanyInfo() {
}

// @Title UpdateCompanyTelephone
// @Description update the company telephone number.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Tel         body      string      true        "The telephone number to update"
// @Success 200 {string} update successful
// @Failure 401 {string} auth failed
// @Failure 401 {string} company not existed
// @router /tel [put]
func (u *CompanyController) UpdateCompanyTel() {
}

// @Title UpdateCompanyEmailAddress
// @Description update company email
// @Param   Token       header    string      true        "The token to conform"
// @Param   Email       body      string      true        "The email address to update"
// @Success 200 {string} update successful
// @Failure 401 {string} auth failed
// @Failure 401 {string} company not existed
// @router /email [put]
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