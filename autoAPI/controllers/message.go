package controllers

import "github.com/astaxie/beego"

// Operations about Messages
type MessageController struct {
	beego.Controller
}

// @Title PlatformMessageList
// @Description Get platform message list.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Page        Query     int         true        "The page number"
// @Success 200 {object} models.PlatformMessageList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /platform/list [get]
func (u *MessageController) GetPlatformMessageList() {
}

// @Title FeedbackMessageList
// @Description Get feedback message list.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.FeedbackMessageList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /feedback/list [get]
func (u *MessageController) GetFeedbackMessageList() {
}

// @Title GetMessageInformationByID
// @Description Get the message information
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.MessageEntity
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | message not existed
// @router /:id [get]
func (u *MessageController) GetMessageInfo() {
}

// @Title ReadMessage
// @Description Read a message, but it might be done by backend.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {string} ok
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | message not existed
// @router /:id [put]
func (u *MessageController) ReadMessage() {
}

// @Title DeleteMessage
// @Description Delete a message.
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {string} delete successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | message not existed
// @router /:id [delete]
func (u *MessageController) DeleteMessage() {
}

// @Title CompanyMessageList
// @Description Get message not from platform, might from another company or a user
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanyMessageList
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router /company/list [get]
func (u *MessageController) GetCompanyMessageList() {
}