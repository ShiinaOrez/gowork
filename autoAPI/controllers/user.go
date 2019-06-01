package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title GetUserInformationByID
// @Description Get user's information by ID. MissionList中，Status指是否已经完成。deadline指结束时间，格式2018-01-01
// @Param   id           path    int            true        "The user ID"
// @Param   Token         header  string         true        "The token to conform"
// @Success 200 {object} models.GetUserInformation
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /info/:ID [Get]
func (u *UserController) UserInfo() {
}


// @Title GetUserBaseInformation
// @Description Get user base information
// @Param   Token         header  string         true        "The token to conform"
// @Success 200 {object} models.GetUserBaseInformation
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /info [get]
func (u *UserController) UserBaseInfo() {
}

// @Title UserMissionInformation
// @Description Get user's mission information
// @Param   Token         header  string         true        "The token to conform"
// @Param   Type          query   string         true        "query mode: all, doing, finish"
// @Success 200 {object} models.UserMissionInformation
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /mission/info [get]
func (u *UserController) UserMissionInfo() {
}

// @Title UserPayAttentionToCompany
// @Description Pay attention to a company by id
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      string      true        "The company id"
// @Success 200 {string} pay attention successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @Failure 405 {string} already attention
// @router /attention/:id [post]
func (u *UserController) UserAttentionCompany() {
}

// @Title UserCancelAttentionToCompany
// @Description User cancel pay attention to a company by id.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      string      true        "The company id"
// @Success 200 {string} cancel successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed | haven't attention
// @router /attention/:id [put]
func (u *UserController) UserCancelAttentionCompany() {
}

// @Title UserAttentionList
// @Description Get user's attention list
// @Param   Token       header    string      true        "The token to conform"
// @Param   page      query      int  		  true 		  "页码 从0开始"
// @Param   limit 		query    int  			true  	  "limit, 每页的尺寸"
// @Success 200 {object} models.UserAttentionCompanyList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /attentions [get]
func (u *UserController) UserAttentionList() {
}

// @Title UserAttentionCompanyStatus
// @Description Query is a user-company attention relationship existed.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      string      true        "The company id"
// @Success 200 {int}
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /attention/:id/status [get]
func (u *UserController) AttentionCompanyStatu() {
}

// @Title UpdateAvatar
// @Description Update user's avatar
// @Param   Token       header    string      true        "The token to conform"
// @Param   newAvatar   body      string      true        "The avatar image"
// @Success 200 {string} newAvatarURL
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @Failure 406 {string} avatar upload failed
// @router /avatar [post]
func (u *UserController) UpdateAvatar() {
}

// @Title EditSetting
// @Description Edit personal information
// @Param   Token       header    string      true        "The token to conform"
// @Param   RequestBody body      models.EditSettingBody  true   "The request body for edit setting"
// @Success 200 {string} edit successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /info [post]
func (u *UserController) EditSetting() {
}

// @Title ConfirmEmailAddress
// @Description Confirm email address by code and change it
// @Param   Token       header    string      true        "The token to conform"
// @Param   Email       body      string      true        "The email address you want to change"
// @Param   Code        body      string      true        "The confirm code sent by server"
// @Success 200 {string} confirm successful
// @Failure 401 {string} auth | confirm failed
// @Failure 404 {string} user not existed
// @router /email/edit [post]
func (u *UserController) ConfirmEmail() {
}

// @Title ConfirmTelephoneNumber
// @Description Confirm telephone number by code and change it
// @Param   Token       header    string      true        "The token to conform"
// @Param   Tel         body      string      true        "The telephone number you want to change"
// @Param   Code        body      string      true        "The confirm code sent by server"
// @Success 200 {string} confirm successful
// @Failure 401 {string} auth | confirm failed
// @Failure 404 {string} user not existed
// @router /tel/edit [post]
func (u *UserController) ConfirmTel() {
}

// @Title PersonalSettingInformation
// @Description Get user's personal setting information
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.UserPersonalSetting
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /personal/setting [get]
func (u *UserController) PersonalSetting() {
}

// @Title EditPersonalSetting
// @Description Edit the personal setting.
// @Param   Token       header    string      true        "The token to conform"
// @Param   RequestBody body      models.UserPersonalSetting true "test"
// @Success 200 {string} edit successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /personal/setting [post]
func (u *UserController) EditPersonalSetting() {
}

// @Title CompanySuitabilityList
// @Description Get the user-company suitability list.
// @Success 200 {object} models.GetCompanySuitabilityList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /company/suitability/list [get]
func (u *UserController) GetSuitabilityList() {
}

// @Title UserResetPassword
// @Description If user forget password. Get email and confirm code.
// @Param   Payload   body   models.ResetPasswordPayload  true  "The reset password payload."
// @Success 200 {string} change successful
// @Failure 401 {string} cofirm code incorrect
// @Failure 404 {string} user not found
// @router /password [put]
func (u *UserController) ResetPassword() {
}