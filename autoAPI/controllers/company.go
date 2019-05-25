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
// @Param   id         path      int         true        "The company ID for get"
// @Success 200 {object} models.CompanyInformation
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /info/:id/ [get]
func (u *CompanyController) CompanyInfo() {
}


// @Title CompanySearchByIntelligent
// @Description Search company by intelligent way.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanySearchList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /search/intelligent [get]
func (u *CompanyController) CompanySearchIntelligent() {
}

// @Title GetCompanyAttentionList
// @Description Get company's attention elites list
// @Param   Token       header    string      true        "The token to conform"
// @Param 	page		query 	  int 		  true 		  "页码，从0开始"
// @Param 	limit		query 	  int 		  true 		  "limit 可以理解为每页有多少"
// @Success 200 {object} models.ElitesList
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /attention/list [get]
func (u *CompanyController) CompanyAttentionList() {
}

// @Title CompanyPayAttentionToSomebody
// @Description Pay attention to a elite.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.CompanyPayAttentionBody         true        "The elite ID for attention"
// @Success 200 {string} pay attention successful!
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | elite not existed
// @Failure 407 {string} already pay attention to it
// @router /attention [post]
func (u *CompanyController) CompanyPayAttention() {
}

// @Title CompanyCancelAttentionSomebody
// @Description Cancel to pay attention to a elite
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload       body      models.CompanyCancelAttentionBody        true        "被关注的精英的ID"
// @Success 200 {string} cancel successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | user not existed
// @Failure 407 {string} never pay attention to it
// @router /attention [put]
func (u *CompanyController) CompanyCancelAttention() {
}


// @Title GetCompanyMissionListByID
// @Description Get a company mission list by company ID. 获取一个公司的任务列表（不是选材清单）
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The company ID for get mission list"
// @Param   page       query      int  		  true 		  "页码，从0开始"
// @Param 	limit 	   query	 int 		  true 		  "limit 可以视为每一页有多少"
// @Success 200 {object} models.MissionList
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /mission/:id/list [get]
func (u *CompanyController) GetCompanyMissionList() {
}

// @Title UpdateCompanyLogo
// @Description Update company logo.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Logo        body      models.CompanyLogoModify      true        "The company new logo."
// @Success 200 {string} ok
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /info/logo [put]
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
// @Title UpdateCompanyTelephoneStep1
// @Description 获取手机号，给相应手机号发验证码（长6位的数字）
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload         body      models.UpdateTelStep1Payload      true        "手机号"
// @Success 200 {}
// @Failure 400 {} 验证码不正确或失效
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /info/tel/step1 [put]
func (u *CompanyController) UpdateCompanyTelStep1() {
}

// TODO
// @Title UpdateCompanyTelephoneStep2
// @Description 获取验证码(长6位的数字)，如果验证码正确，则更改手机号
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload         body     models.UpdateTelStep2Payload   true        "The telephone number to update"
// @Success 200 {} update successful
// @Failure 400 {} 验证码不正确或失效
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /info/tel/step2 [put]
func (u *CompanyController) UpdateCompanyTelStep2() {
}

// TODO
// @Title UpdateCompanyEmailAddressStep1
// @Description 前端发送一个邮箱地址，给相应邮箱发送一个验证码（6位长数字）。
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload       body      models.UpdateEmailStep1Payload      true        "The email address to update"
// @Success 200 {}
// @Failure 400 {} 验证码不正确或失效
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /info/email/step1 [put]
func (u *CompanyController) UpdateCompanyEmailStep1() {
}

// @Title UpdateCompanyEmailAddressStep2
// @Description 填写6位长验证码，若正确则更改
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload       body      models.UpdateEmailStep2Payload      true        "The email address to update"
// @Failure 400 {} 验证码不正确或失效
// @Success 200  {} update successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /info/email/step2 [put]
func (u *CompanyController) UpdateCompanyEmailStep2() {
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

// @Title CreateMissionList
// @Description Create and post a missionList.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.MissionListCreateInformation  true   "mission list Payload "
// @Success 200 {object} models.CreateMissionListResponse
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /missionList [post]
func (u *CompanyController) CreateCompanyMissionList() {
}

// @Title DeleteMissionList
// @Description 删除一个选材清单
// @Param 	Token 		header    string      true        "The token to conform"
// @Param 	mlid 		query	  int         true 		"Mission List ID"
// @Success 200 {} create successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | category not existed
// @router /missionList/ [delete]
func (u *CompanyController) DeleteCompanyMissionList() {
}

// @Title AddMissionToMissionList
// @Description 向一个MissionList中增加一个Mission
// @Param Token       header    string      true        "公司用户Token"
// @Param Payload		body 	models.AddMissionToMissionListRequest true "Mission Id and missionList ID"
// @Success 200 {} create successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | category not existed
// @router /missionList/mission [post]
func (u *CompanyController) AddMissionToMissionList() {

}

// TOOD
// @Title RemoveMissionInMissionList
// @Description 删除一个MissionList中的Mission
// @Param Token       header    string      true        "公司用户Token"
// @Param Payload     body      models.RemoveMissionInListRequest true "Mission Id and missionList ID"
// @Success 200 {} delete successful
// @Failure 401 {} auth failed
// @Failure 404 {} company | category not existed
// @router /missionList/mission [delete]
func (u *CompanyController) RemoveMissionInMissionList() {

}

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

// @Title AddTagToCompanyCategory
// @Description Create a tag and join it to category
// @Param   Token       header    string      true        "The token to conform"
// @Param   id      path      int         true        "The category ID"
// @Param   TagName     body      models.CreateTagForCompanyBody      true        "tag name"
// @Success 200 {string} create successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} company | category not existed
// @router /category/:id/tags [post]
func (u *CompanyController) CompanyCategoryAddTag() {
}



// @Title GetAllMissionListOfCompany
// @Description Get a company all mission list. 获取一个公司的全部选材清单
// @Param   Token       header    string      true        "The token to conform"
// @Param 	id 			path 	 	int 		true 		"想查看的公司的ID"
// @Param   page      query      int  		  true 		  "页码 从0开始"
// @Param   limit 		query    int  			true  	  "limit, 每页的尺寸"
// @Success 200 {object} models.MissionListList
// @Failure 401 {string} sign up failed
// @Failure 404 {string} user | company not existed
// @router /missionList/:id/list/ [get]
func (u *CompanyController) CompanyAllMissionList() {
}

// @Title CompanyMissionNumberOverview
// @Description 公司任务数目概览，包括结束，报名中，任务提交中.
// @Param   id         path      int         true        "The company ID for get"
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanyMissionNumberOverview
// @Failure 401 {string} sign up failed
// @Failure 404 {string} user | company not existed
// @router /mission/:id/overview [get]
func (u *CompanyController) CompanyMissionNumberOverview() {
}

// @Title ChangeAllowSendEmail
// @Description If company want to change message setting.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.CompanyAllowSendEmailPayload true "是否允许发送邮件"
// @Success 200 {object} change successful!
// @Failure 401 {string} sign up failed
// @Failure 404 {string} company not existed
// @router /info/email/allow [put]
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
// @router /password/:id [put]
func (u *UserController) CompanyChangePassword() {
}