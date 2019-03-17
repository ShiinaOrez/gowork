package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title UserInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Uid/info [post]
func (u *UserController) UserInfo() {
}

// @Title UserBaseInformation
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /info [post]
func (u *UserController) UserBaseInfo() {
}

// @Title UserMissionInformation
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /mission/info [post]
func (u *UserController) UserMissionInfo() {
}

// @Title UserPayAttentionToCompany
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:Cid [post]
func (u *UserController) UserAttentionCompany() {
}

// @Title UserCancelAttentionToCompany
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:Cid [post]
func (u *UserController) UserCancelAttentionCompany() {
}

// @Title UserAttentionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/list [post]
func (u *UserController) UserAttentionList() {
}

// @Title UserAttentionCompanyStatus
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:Cid/status [post]
func (u *UserController) AttentionCompanyStatu() {
}

// @Title UpdateAvatar
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /avatar [post]
func (u *UserController) UpdateAvatar() {
}

// @Title EditSetting
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /info [post]
func (u *UserController) EditSetting() {
}

// @Title PlatformInformation
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /platform/setting [post]
func (u *UserController) PlatformSetting() {
}

// @Title EditPlatformSetting
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /platform/setting [post]
func (u *UserController) EditPlatformSetting() {
}

// @Title CompanySuitabilityList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /company/suitability/list [post]
func (u *UserController) GetSuitabilityList() {
}