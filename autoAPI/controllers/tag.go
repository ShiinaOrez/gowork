package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Tags
type TagController struct {
	beego.Controller
}

// @Title CreateTag
// @Description Provide category id and tag information.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload   body   models.CreateTagPayload   true   "The payload to create tag."
// @Success 200   {object}   models.CreateTagReturn
// @Failure 400   {string}   bad request
// @Failure 401   {string}   authentication fail.
// @router / [post]
func (t *TagController) CreateTag() {
}

// @Title DeleteTagAPI
// @Description 删除一个Tag
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "TagID"
// @Success 200 {}
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /:id [delete]
func (t *TagController) DeleteTag() {
}

// @Title TagUpdateAPI
// @Description 更新Tag
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "TagID"
// @Param   UpdateTagPayload   body   models.UpdateTagPayload   true   "The payload to update tag."
// @Success 200 {}
// @Failure 400 {string}   bad request
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /:id [put]
func (t *TagController) UpdateTag() {
}

// @Title GetTagInformation
// @Description Get the tag information
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "TagID"
// @Success 200 {object} models.GetTagReturn
// @Failure 401 {string} auth failed
// @Failure 404 {string} not found
// @router /:id/ [get]
func (t *TagController) GetTag() {
}
