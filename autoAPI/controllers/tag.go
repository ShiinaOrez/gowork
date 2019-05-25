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
// @Param   Payload   body   models.CreateTagPayload   true   ""
// @Success 200   {object}   models.CreateTagReturn
// @Failure 400   {string}   bad request
// @Failure 401   {string}   authentication fail.
// @router / [post]
func (t *TagController) CreateTag() {

}

// @Title DeleteTag
// @Description Delete tag by tag id.
// @Param   TagID   path   int   "The tag id for query and "
// @Success 200   {string}   delete successful
// @Failure 401   {string}   authentication fail.
// @Failure 404   {string}   tag not found
// @router /:id  [delete]
func (t *TagController) DeleteTag() {

}

// @Title UpdateTag
// @Description Update a tag, but just change the tag name
// @Param   TagID   path   int   "The tag id to query"
// @Param   UpdateTagPayload   body   true   models.UpdateTagPayload
// @Success 200   {string}   OK
// @Failure 400   {string}   bad request
// @Failure 401   {string}   authentication fail
// @Failure 404   {string}   not found
// router /:id  [put]
func (t *TagController) UpdateTag() {

}

// @Title GetTagInformation
// @Description Get the tag information
// @Param   TagID   path   int   "The tag id to query."
// @Success 200   {object}   models.GetTagReturn
// @Failure 401   {string}   authentication fail
// @Failure 404   {string}   not found
// @router /:id  [get]
func (t *TagController) GetTag() {

}