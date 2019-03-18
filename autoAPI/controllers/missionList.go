package controllers

import "github.com/astaxie/beego"

type MissionListController struct {
	beego.Controller
}

// @Title GetMissionListInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MlID [post]
func (u *MissionController) GetMissionListInformation() {
}

// @Title DeleteMissionListByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MlID [post]
func (u *MissionController) DelMissionList() {
}

// @Title GetUserToMissionListStatusByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:MlID/status [post]
func (u *MissionController) GetMissionListStatus() {
}

// @Title PayAttentionToMissionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:MlID [post]
func (u *MissionController) AttentionMissionList() {
}

// @Title CancelMissionListAttention
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/:MlID [put]
func (u *MissionController) CancelAttentionMissionList() {
}

// @Title GetMissionListAttentionList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /attention/list [post]
func (u *MissionController) GetMissionListAttentionList() {
}

// @Title MissionListGetElite
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:MlID/getElite [post]
func (u *MissionController) GetEliteByMissionList() {
}
