package controllers

import "github.com/astaxie/beego"

// Operations about MissionLists
type MissionListController struct {
	beego.Controller
}

// @Title GetMissionListInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MlID [post]
func (u *MissionListController) GetMissionListInformation() {
}

// @Title DeleteMissionListByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MlID [post]
func (u *MissionListController) DelMissionList() {
}

// @Title GetUserToMissionListStatusByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:MlID/status [post]
func (u *MissionListController) GetMissionListStatus() {
}

// @Title PayAttentionToMissionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:MlID [post]
func (u *MissionListController) AttentionMissionList() {
}

// @Title CancelMissionListAttention
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:MlID [put]
func (u *MissionListController) CancelAttentionMissionList() {
}

// @Title GetMissionListAttentionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/list [post]
func (u *MissionListController) GetMissionListAttentionList() {
}

// @Title MissionListGetElite
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MlID/getElite [post]
func (u *MissionListController) GetEliteByMissionList() {
}