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
// @Failure 404 {string} user | company not existed
// @router /:Cid/info [get]
func (u *CompanyController) CompanyInfo() {
}

// TODO
// @Title CompanyBaseInformationByID
// @Description Get company base information.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanyBaseInformation
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /info [get]
func (u *CompanyController) CompanyBaseInfo() {
}

// TODO
// @Title CompanySearchByIntelligent
// @Description Search company by intelligent way.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanySearchList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /search/intelligent [get]
func (u *CompanyController) CompanySearchIntelligent() {
}

// TODO
// @Title GetCompanyAttentionList
// @Description Get company's attention elites list
// @Param   Token       header    string      true        "The token to conform"
// @Param 	PageId		path 	  int 		  true 		  "页码，从1开始"
// @Success 200 {object} models.ElitesList
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /attention/list/:PageId [get]
func (u *CompanyController) CompanyAttentionList() {
}

// TODO
// @Title CompanyPayAttentionToSomebody
// @Description Pay attention to a elite.
// @Param   Token       header    string      true        "The token to conform"
// @Param   EliteID     body      models.CompanyPayAttentionBody         true        "The elite ID for attention"
// @Success 200 {string} pay attention successful!
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | elite not existed
// @Failure 407 {string} already pay attention to it
// @router /attention [post]
func (u *CompanyController) CompanyPayAttention() {
}

// TODO
// @Title CompanyCancelAttentionSomebody
// @Description Cancel to pay attention to a elite
// @Param   Token       header    string      true        "The token to conform"
// @Param   Uid         path      int         true        "The elite ID for attention"
// @Success 200 {string} cancel successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | user not existed
// @Failure 407 {string} never pay attention to it
// @router /attention/:Uid [put]
func (u *CompanyController) CompanyCancelAttention() {
}

// TODO
// @Title GetCompanyPositionListByID
// @Description Get a company's position list
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      int         true        "The company ID for get position list"
// @Param   PageId      path      int  		  true 		  "页码"
// @Success 200 {object} models.PositionsList
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /:Cid/position/list/:PageId [get]
func (u *CompanyController) GetCompanyPositionList() {
}

// TODO
// @Title CreateAndPostPosition
// @Description Company create a position.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.PositionPayload  true   "PositionPayload "
// @Success 200 {int} positionID
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /position [post]
func (u *CompanyController) CreateCompanyPosition() {
}

// TODO
// @Title EditCompanyPositionByID
// @Description Edit the position information
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.PositionPayload   true   "PositionPayload "
// @Success 200 {string} edit successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | position not existed
// @router /position/:Pid [put]
func (u *CompanyController) EditCompanyPosition() {
}

// TODO
// @Title DeleteCompanyPositionByID
// @Description Delete position by ID
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {string} delete successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | position not existed
// @router /position/:Pid [delete]
func (u *CompanyController) DeleteCompanyPosition() {
}

// TODO
// @Title GetCompanyMissionListByID
// @Description Get a company mission list by company ID. 获取一个公司的任务列表（不是选材清单）
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      int         true        "The company ID for get mission list"
// @Param   PageId      path      int  		  true 		  "页码"
// @Success 200 {object} models.MissionList
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /:Cid/mission/list/:PageId [get]
func (u *CompanyController) GetCompanyMissionList() {
}

// @Title UpdateCompanyLogo
// @Description Update company logo.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Logo        body      models.CompanyLogoModify      true        "The company new logo."
// @Success 200 {string} ok
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /logo [put]
func (u *CompanyController) UpdateCompanyLogo() {
}

// @Title EditCompanyInformation
// @Description Edit company information
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.CompanyInfoModify   true   "edit company info body"
// @Success 200 {string} edit ok
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @Failure 406 {string} image error
// @router /info [put]
func (u *CompanyController) EditCompanyInfo() {
}

// TODO
// @Title UpdateCompanyTelephone
// @Description update the company telephone number.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Tel         body      string      true        "The telephone number to update"
// @Success 200 {string} update successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /tel [put]
func (u *CompanyController) UpdateCompanyTel() {
}

// TODO
// @Title UpdateCompanyEmailAddress
// @Description update company email
// @Param   Token       header    string      true        "The token to conform"
// @Param   Email       body      string      true        "The email address to update"
// @Success 200 {string} update successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /email [put]
func (u *CompanyController) UpdateCompanyEmail() {
}

// TODO
// @Title CompanySearchByFastWay
// @Description Search company by very fast way
// @Param   Token       header    string      true        "The token to conform"
// @Param   Pattern     body      string      true        "The pattern to make a fast search"
// @Success 200 {object} models.CompanyVerySimpleInfoList
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /search/fast [post]
func (u *CompanyController) CompanySearchFast() {
}

// TODO
// @Title CreateMissionList
// @Description Create and post a missionList.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.MissionListCreateInformation  true   "PositionPayload "
// @Success 200 {object} models.CreateMissionListResponse
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /missionList [post]
func (u *CompanyController) CreateCompanyMissionList() {
}

// TODO
// @Title AddMissionToMissionList
// @Description 向一个MissionList中增加一个Mission
// @Param Token       header    string      true        "公司用户Token"
// @Param Payload		body 	models.AddMissionToMissionListRequest true "Mission Id"
// @Param MlID			path    int         true 		"Mission List ID"
// @Success 200 {} create successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | category not existed
// @router /missionList/:MlID [post]
func (u *CompanyController) AddMissionToMissionList() {

}

// TOOD
// @Title RemoveMissionInMissionList
// @Description 删除一个MissionList中的Mission
// @Param Token       header    string      true        "公司用户Token"
// @Param Payload     body      models.RemoveMissionInListRequest true "Mission Id"
// @Param MlID			path    int         true 		"Mission List ID"
// @Success 200 {} delete successful
// @Failure 401 {} auth failed
// @Failure 404 {} company | category not existed
// @router /missionList/:MlID [delete]
func (u *CompanyController) RemoveMissionInMissionList() {

}

// TODO
// @Title CreateCompanyCategory
// @Description Create a category for company.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Name        body      models.CreateCompanyCategoryBody      true        "The new category name"
// @Success 200 {object} models.CategoryCreateResponse
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @Failure 407 {string} already exist a same category
// @router /category [post]
func (u *CompanyController) CreateCompanyCategory() {
}

// TODO
// @Title AddTagToCompanyCategory
// @Description Create a tag and join it to category
// @Param   Token       header    string      true        "The token to conform"
// @Param   CateID      path      int         true        "The category ID"
// @Param   TagName     body      models.CreateTagForCompanyBody      true        "tag name"
// @Success 200 {string} create successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | category not existed
// @router /category/:CateID/tags [post]
func (u *CompanyController) CompanyCategoryAddTag() {
}



// @Title SearchCompanyPositionByNormalWay
// @Description Search company positions.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Pattern     query     string      true        "The pattern string"
// @Success 200 {object} models.Positions
// @Failure 401 {string} sign up failed
// @Failure 404 {string} user | position not existed
// @router /position/:Pid/search [get]
func (u *CompanyController) CompanyPositionSearch() {
}


// TODO
// @Title GetAllMissionListForCompany
// @Description Get a company all mission list. 获取一个公司的全部选材清单
// @Param   Token       header    string      true        "The token to conform"
// @Param   PageId      path      int  		  true 		  "页码"
// @Success 200 {object} models.MissionListList
// @Failure 401 {string} sign up failed
// @Failure 404 {string} user | company not existed
// @router /:Cid/missionList/all/:PageId [get]
func (u *CompanyController) CompanyAllMissionList() {
}

// @Title CompanyMissionNumberOverview
// @Description 公司任务数目概览，包括结束，报名中，任务提交中
// @Param   Cid         path      int         true        "The company ID for get"
// @Param   Token       header    string      true        "The token to conform"
// @Param   PageId      path      int  		  true 		  "页码"
// @Success 200 {object} models.CompanyMissionNumberOverview
// @Failure 401 {string} sign up failed
// @Failure 404 {string} user | company not existed
// @router /:Cid/mission/overview [get]
func (u *CompanyController) CompanyMissionNumberOverview() {
}

// @Title ChangeMessageSetting
// @Description If company want to change message setting.
// @Param   Cid         path      int         true        "The company ID for change message setting"
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.CompanyMessagePayload true "The new message setting."
// @Success 200 {object} change successful!
// @Failure 401 {string} sign up failed
// @Failure 404 {string} company not existed
// @router /:Cid/message/setting [put]
func (u *CompanyController) CompanyMessageSetting() {
}

// @Title ComapnyChangePassword
// @Description If company want to change password.
// @Param   Token     header     string    true    "The token."
// @Param   Cid       path       int       true    "The user id to get user."
// @Param   Payload   body   models.ChangePasswordPayload  true  "The change password payload"
// @Success 200 {string} change successful
// @Failure 401 {string} cofirm code incorrect
// @Failure 404 {string} company not found
// @router /:Cid/password [put]
func (u *UserController) CompanyChangePassword() {
}