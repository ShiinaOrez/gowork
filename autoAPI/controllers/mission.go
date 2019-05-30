package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Missions
type MissionController struct {
	beego.Controller
}

// @Title MissionInformationByID
// @Description Mission information.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to get"
// @Success 200 {object} models.MissionInformation
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | mission not existed
// @router /:id/info [get]
func (u *MissionController) MissionInfo() {
}

// @Title SearchForMission
// @Description Search mission.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.SearchPayload    true     "The payload"
// @Param   page      query      int  		  true 		  "页码 从0开始"
// @Param   limit 		query    int  			true  	  "limit, 每页的尺寸"
// @Success 200 {object} models.MissionList
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /search [post]
func (u *MissionController) MissionSearch() {
}

// @Title GainMission
// @Description Join a mission if permitted.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to Gain"
// @Success 200 {string} Join successful.
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | mission not existed
// @Failure 407 {string} gain failed
// @router /:id [put]
func (u *MissionController) GainMission() {
}

// @Title GetMissionStatusByID
// @Description Get the status between user and mission.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to Gain"
// @Success 200 {int} 1 for gained, 0 for not
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | mission not existed
// @router /:id/status [get]
func (u *MissionController) MissionStatus() {
}

// @Title SubmitMissionByID
// @Description Submit the works for mission
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to submit"
// @Param   Payload     body      models.DoMissionPayload    true    "The payload"
// @Success 200 {string} submit successfully!
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | mission not existed
// @Failure 406 {string} File error
// @Failure 407 {string} submit failed
// @router /:id [post]
func (u *MissionController) SubmitMission() {
}

// @Title UploadFileOnlyZip
// @Description Upload the file.
// @Param   Token       header    string      true        "The token to conform"
// @Param   File        body      string      true        "The file you want to upload"
// @Success 200 {string} File URL
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | mission not existed
// @Failure 406 {string} File error
// @router /file/upload [post]
func (u *MissionController) MissionUploadFile() {
}

// @Title GiveUpMission
// @Description If you want to give up mission...
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to give up"
// @Success 200 {string} give up successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | mission not existed
// @Failure 407 {string} never gain this mission
// @router /:id [delete]
func (u *MissionController) GiveUpMission() {
}

// @Title GetMissionFeedback
// @Description Get the mission feedback. (感觉好像没啥用?)
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to give up"
// @Success 200 {object} models.CompanyDoMissionFeedback
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | mission not existed
// @router /:id/feedback [get]
func (u *MissionController) SubmissionFeedback() {
}

// @Title CreateAndPostMission
// @Description Create a mission and post it.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      models.CreateMissionPayload    true     "The payload"
// @Success 200 {int} MissionID
// @Failure 401 {string} auth failed
// @Failure 404 {string} company not existed
// @router / [post]
func (u *MissionController) CreateMission() {
}

// @Title ModifyMissionByID
// @Description Company modify mission information.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to modify"
// @Param   Payload     body      models.CreateMissionPayload     true    "The payload"
// @Success 200 {string} modify successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} mission not existed
// @router /:id [put]
func (u *MissionController) ModifyMission() {
}

// @Title DeleteMissionByID
// @Description Delete a mission by ID.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to delete"
// @Success 200 {string} delete successful
// @Failure 401 {string} auth failed
// @Failure 404 {string} mission not existed
// @router /:id [delete]
func (u *MissionController) DeleteMission() {
}

// @Title GetStaticMissionTagList
// @Description Get the mission tags list.
// @Param   id         path      int         true        "The mission id to get"
// @Param   page      query      int  		  true 		  "页码 从0开始"
// @Param   limit 		query    int  			true  	  "limit, 每页的尺寸"
// @Success 200 {object} models.TagList
// @Failure 407 {string} something wrong
// @router /:id/tag/list [get]
func (u *MissionController) MissionStaticTagList() {
}

// @Title MissionStatisticsByID
// @Description New user must sign up.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The mission id to get"
// @Success 200 {object} models.MissionStatistics
// @Failure 401 {string} auth failed
// @Failure 404 {string} mission not existed
// @router /:id/statistics [get]
func (u *MissionController) MissionStatistics() {
}