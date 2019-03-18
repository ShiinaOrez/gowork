package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Tags
type TagController struct {
	beego.Controller
}

// @Title TagIndustryList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /industry/list [post]
func (u *TagController) TagIndustryList() {
}

// @Title TagPositionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /position/list [post]
func (u *TagController) TagPositionList() {
}

// @Title TagUniversityList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /university/list [post]
func (u *TagController) TagUniversityList() {
}

// @Title TagAbilityList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /ability/list [post]
func (u *TagController) TagAbilityList() {
}