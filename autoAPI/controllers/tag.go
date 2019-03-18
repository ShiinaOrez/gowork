package controllers

import "github.com/astaxie/beego"

type TagController struct {
	beego.Controller
}

// @Title TagIndustryList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /industry/list [post]
func (u *MissionController) TagIndustryList() {
}

// @Title TagPositionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /position/list [post]
func (u *MissionController) TagPositionList() {
}

// @Title TagUniversityList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /university/list [post]
func (u *MissionController) TagUniversityList() {
}

// @Title TagAbilityList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /ability/list [post]
func (u *MissionController) TagAbilityList() {
}