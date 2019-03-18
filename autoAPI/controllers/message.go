package controllers

import "github.com/astaxie/beego"

type MessageController struct {
	beego.Controller
}

// @Title PlatformMessageList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /platform/list [post]
func (u *MissionController) GetPlatformMessageList() {
}

// @Title FeedbackMessageList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /feedback/list [post]
func (u *MissionController) GetFeedbackMessageList() {
}

// @Title GetMessageInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MsgID [post]
func (u *MissionController) GetMessageInfo() {
}

// @Title ReadMessage
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MsgID [put]
func (u *MissionController) ReadMessage() {
}

// @Title CompanyMessageList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /company/list [post]
func (u *MissionController) GetCompanyMessageList() {
}