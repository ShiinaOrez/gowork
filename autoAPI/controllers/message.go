package controllers

import "github.com/astaxie/beego"

// Operations about Messages
type MessageController struct {
	beego.Controller
}

// @Title PlatformMessageList
// @Description Get platform message list.
// @Success 200 {object} models.PlatformMessageList
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | company not existed
// @router /platform/list [get]
func (u *MessageController) GetPlatformMessageList() {
}

// @Title FeedbackMessageList
// @Description Get feedback message list.
// @Success 200 {object} models.FeedbackMessageList
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /feedback/list [get]
func (u *MessageController) GetFeedbackMessageList() {
}

// @Title GetMessageInformationByID
// @Description Get the message information
// @Success 200 {object} models.MessageEntity
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | message not existed
// @router /:MsgID [get]
func (u *MessageController) GetMessageInfo() {
}

// @Title ReadMessage
// @Description Read a message, but it might be done by backend.
// @Success 200 {string} ok
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | message not existed
// @router /:MsgID [put]
func (u *MessageController) ReadMessage() {
}

// @Title DeleteMessage
// @Description Delete a message.
// @Success 200 {string} delete successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | message not existed
// @router /:MsgID [delete]
func (u *MessageController) DeleteMessage() {
}

// @Title CompanyMessageList
// @Description Get message not from platform, might from another company or a user
// @Success 200 {object} models.CompanyMessageList
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @router /company/list [get]
func (u *MessageController) GetCompanyMessageList() {
}