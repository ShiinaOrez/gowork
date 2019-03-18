package controllers

import "github.com/astaxie/beego"

// Operations about Messages
type MessageController struct {
	beego.Controller
}

// @Title PlatformMessageList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /platform/list [post]
func (u *MessageController) GetPlatformMessageList() {
}

// @Title FeedbackMessageList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /feedback/list [post]
func (u *MessageController) GetFeedbackMessageList() {
}

// @Title GetMessageInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MsgID [post]
func (u *MessageController) GetMessageInfo() {
}

// @Title ReadMessage
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MsgID [put]
func (u *MessageController) ReadMessage() {
}

// @Title CompanyMessageList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /company/list [post]
func (u *MessageController) GetCompanyMessageList() {
}