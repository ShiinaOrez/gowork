package controllers

import "github.com/astaxie/beego"

// Operations about Elites
type EliteController struct {
	beego.Controller
}

// @Title SearchElite
// @Description Search elite by normal way.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.SearchPayload    true    "The payload"
// @Success 200 {object} models.EliteList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /search [post]
func (u *EliteController) SearchElite() {
}

// @Title SearchEliteByIntelligentWay
// @Description Search elite by intelligent way.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.SearchPayload    true  "Search payload"
// @Success 200 {object} models.EliteList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /industry/list [post]
func (u *EliteController) SearchEliteIntelligent() {
}