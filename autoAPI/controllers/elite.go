package controllers

import "github.com/astaxie/beego"

// Operations about Elites
type EliteController struct {
	beego.Controller
}

// @Title SearchElite
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /search [post]
func (u *EliteController) SearchElite() {
}

// @Title SearchEliteByIntelligentWay
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /industry/list [post]
func (u *EliteController) SearchEliteIntelligent() {
}