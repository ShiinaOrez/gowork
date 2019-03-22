package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title GetUserInformationByID
// @Description Get user's information by ID
// @Param   Uid           path    int            true        "The user ID"
// @Param   Token         header  string         true        "The token to conform"
// @Success 200 {object} models.GetUserInformation
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /:Uid/info [Get]
func (u *UserController) UserInfo() {
}

// @Title GetUserBaseInformation
// @Description Get user base information
// @Param   Token         header  string         true        "The token to conform"
// @Success 200 {object} models.GetUserBaseInformation
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /info [get]
func (u *UserController) UserBaseInfo() {
}

// @Title UserEditPassword
// @Description User edit password.
// @Param   Token         header  string         true        "The token to conform"
// @Param   oldPassword   body    string         true        "The old password to change"
// @Param   newPassword   body    string         true        "The new password want to change"
// @Success 200 {object} models.GetUserBaseInformation
// @Failure 401 {string} auth failed || old password confirm failed
// @Failure 402 {string} user not existed
// @router /editPassword [post]
func (u *UserController) UserEditPassword() {
}

// @Title UserMissionInformation
// @Description Get user's mission information
// @Param   Token         header  string         true        "The token to conform"
// @Param   Type          query   string         true        "query mode: all, doing, finish"
// @Success 200 {object} models.UserMissionInformation
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /mission/info [get]
func (u *UserController) UserMissionInfo() {
}

// @Title UserPayAttentionToCompany
// @Description Pay attention to a company by Cid
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      string      true        "The company id"
// @Success 200 {string} pay attention successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | company not existed
// @Failure 405 {string} already attention
// @router /attention/:Cid [post]
func (u *UserController) UserAttentionCompany() {
}

// @Title UserCancelAttentionToCompany
// @Description User cancel pay attention to a company by id.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      string      true        "The company id"
// @Success 200 {string} cancel successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | company not existed | haven't attention
// @router /attention/:Cid [put]
func (u *UserController) UserCancelAttentionCompany() {
}

// @Title UserAttentionList
// @Description Get user's attention list
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.UserAttentionCompanyList
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /attention/list [get]
func (u *UserController) UserAttentionList() {
}

// @Title UserAttentionCompanyStatus
// @Description Query is a user-company attention relationship existed.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Cid         path      string      true        "The company id"
// @Success 200 {int}
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | company not existed
// @router /attention/:Cid/status [get]
func (u *UserController) AttentionCompanyStatu() {
}

// @Title UpdateAvatar
// @Description Update user's avatar
// @Param   Token       header    string      true        "The token to conform"
// @Param   newAvatar   body      string      true        "The avatar image"
// @Success 200 {string} newAvatarURL
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @Failure 406 {string} avatar upload failed
// @router /avatar [post]
func (u *UserController) UpdateAvatar() {
}

// @Title EditSetting
// @Description Edit personal information
// @Param   Token       header    string      true        "The token to conform"
// @Param   RequestBody {object}  models.EditSettingBody  true   "The request body for edit setting"
// @Success 200 {string} edit successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
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
// @Failure 402 {string} user not existed
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
// @Failure 402 {string} user not existed
// @router /tel/edit [post]
func (u *UserController) ConfirmTel() {
}

// @Title PlatformInformation
// @Description Get user's platform setting information
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.UserPlatformSetting
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /platform/setting [get]
func (u *UserController) PlatformSetting() {
}

// @Title EditPlatformSetting
// @Description New user must sign up.
// @Param   Token       header    string      true        "The token to conform"
// @Param   RequestBody body      object      models.UserPlatformSetting
// @Success 200 {string} edit successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /platform/setting [post]
func (u *UserController) EditPlatformSetting() {
}

// @Title CompanySuitabilityList
// @Description Get the user-company suitability list.
// @Success 200 {object} models.GetCompanySuitabilityList
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /company/suitability/list [get]
func (u *UserController) GetSuitabilityList() {
}