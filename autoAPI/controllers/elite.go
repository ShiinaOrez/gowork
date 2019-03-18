package controllers

import "github.com/astaxie/beego"

type EliteController struct {
	beego.Controller
}

// @Title SearchElite
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /search [post]
func (u *MissionController) SearchElite() {
}

// @Title SearchEliteByIntelligentWay
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /industry/list [post]
func (u *MissionController) SearchEliteIntelligent() {
}