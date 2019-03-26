package controllers

import "github.com/astaxie/beego"

// Operations about MissionLists
type MissionListController struct {
	beego.Controller
}

// @Title GetMissionListInformationByID
// @Description Get missionList information by ID.
// @Param   Token       header    string      true        "The token to conform"
// @Param   MlID        path      int         true        "The missionList ID to get information"
// @Success 200 {object} models.MissionListInformation
// @Failure 401 {string} auth failed
// @Failure 402 {string} missionList not existed
// @router /:MlID [get]
func (u *MissionListController) GetMissionListInformation() {
}

// @Title DeleteMissionListByID
// @Description Delete a missionList by ID.
// @Param   Token       header    string      true        "The token to conform"
// @Param   MlID        path      int         true        "The missionList ID to delete it"
// @Success 200 {string} delete successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} missionList not existed
// @Failure 407 {string} delete failed
// @router /:MlID [delete]
func (u *MissionListController) DelMissionList() {
}

// @Title GetUserToMissionListStatusByID
// @Description Get a user pay attention to a missionList or not.
// @Param   Token       header    string      true        "The token to conform"
// @Param   MlID        path      int         true        "The missionList ID to get"
// @Success 200 {int} the status
// @Failure 401 {string} auth failed
// @Failure 402 {string} missionList not existed
// @router /attention/:MlID/status [get]
func (u *MissionListController) GetMissionListStatus() {
}

// @Title PayAttentionToMissionList
// @Description User pay attention to a mission list.
// @Param   Token       header    string      true        "The token to conform"
// @Param   MlID        path      int         true        "The missionList ID to get"
// @Success 200 {string} successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} missionList not existed
// @Failure 407 {string} already pay attention
// @router /attention/:MlID [post]
func (u *MissionListController) AttentionMissionList() {
}

// @Title CancelMissionListAttention
// @Description Cancel attention.
// @Param   Token       header    string      true        "The token to conform"
// @Param   MlID        path      int         true        "The missionList ID to get"
// @Success 200 {string} sign up successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} missionList not existed
// @Failure 407 {string} haven't pay attention
// @router /attention/:MlID [put]
func (u *MissionListController) CancelAttentionMissionList() {
}

// @Title GetMissionListAttentionList
// @Description Get a user's attention missionList.
// @Param   Token       header    string      true        "The token to conform"
// @Param   MlID        path      int         true        "The missionList ID to get"
// @Success 200 {object} models.MissionListList
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