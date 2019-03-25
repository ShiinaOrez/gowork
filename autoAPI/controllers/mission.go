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
// @Param   Mid         path      int         true        "The mission id to get"
// @Success 200 {object} models.MissionInformation
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | mission not existed
// @router /:Mid/info [get]
func (u *MissionController) MissionInfo() {
}

// @Title SearchForMission
// @Description Search mission.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      object      models.SearchPayload
// @Success 200 {object} models.MissionList
// @Failure 401 {string} auth failed
// @Failure 402 {string} user not existed
// @router /search [post]
func (u *MissionController) MissionSearch() {
}

// @Title GainMission
// @Description Join a mission if permitted.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to Gain"
// @Success 200 {string} Join successful.
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | mission not existed
// @Failure 407 {string} gain failed
// @router /:Mid [put]
func (u *MissionController) GainMission() {
}

// @Title GetMissionStatusByID
// @Description Get the status between user and mission.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to Gain"
// @Success 200 {int} 1 for gained, 0 for not
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | mission not existed
// @router /:Mid/status [get]
func (u *MissionController) MissionStatus() {
}

// @Title SubmitMissionByID
// @Description Submit the works for mission
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to submit"
// @Param   Payload     body      object      models.SubmitMissionPayload
// @Success 200 {string} submit successfully!
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | mission not existed
// @Failure 406 {string} File error
// @Failure 407 {string} submit failed
// @router /:Mid [post]
func (u *MissionController) SubmitMission() {
}

// @Title UploadFileOnlyZip
// @Description Upload the file.
// @Param   Token       header    string      true        "The token to conform"
// @Param   File        body      string      true        "The file you want to upload"
// @Success 200 {string} File URL
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | mission not existed
// @Failure 406 {string} File error
// @router /file/upload [post]
func (u *MissionController) MissionUploadFile() {
}

// @Title GiveUpMission
// @Description If you want to give up mission...
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to give up"
// @Success 200 {string} give up successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | mission not existed
// @Failure 407 {string} never gain this mission
// @router /:Mid [delete]
func (u *MissionController) GiveUpMission() {
}

// @Title GetMissionFeedback
// @Description Get the mission feedback.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to give up"
// @Success 200 {object} models.CompanyFeedback
// @Failure 401 {string} auth failed
// @Failure 402 {string} user | mission not existed
// @router /:Mid/feedback [get]
func (u *MissionController) MissionFeedback() {
}

// @Title CreateAndPostMission
// @Description Create a mission and post it.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Payload     body      object      models.CreateMissionPayload
// @Success 200 {int} MissionID
// @Failure 401 {string} auth failed
// @Failure 402 {string} company not existed
// @router / [post]
func (u *MissionController) CreateMission() {
}

// @Title ModifyMissionByID
// @Description Company modify mission information.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to modify"
// @Param   Payload     body      object      models.CreateMissionPayload
// @Success 200 {string} modify successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} mission not existed
// @router /:Mid [put]
func (u *MissionController) ModifyMission() {
}

// @Title DeleteMissionByID
// @Description Delete a mission by ID.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to delete"
// @Success 200 {string} delete successful
// @Failure 401 {string} auth failed
// @Failure 402 {string} mission not existed
// @router /:Mid [delete]
func (u *MissionController) DeleteMission() {
}

// @Title GetStaticMissionTagList
// @Description Get the mission tags list.
// @Success 200 {object} models.TagList
// @Failure 407 {string} something wrong
// @router /tag/list [get]
func (u *MissionController) MissionStaticTagList() {
}

// @Title MissionStatisticsByID
// @Description New user must sign up.
// @Param   Token       header    string      true        "The token to conform"
// @Param   Mid         path      int         true        "The mission id to get"
// @Success 200 {object} models.MissionStatistics
// @Failure 401 {string} sign up failed
// @router /:Mid/statistics [get]
func (u *MissionController) MissionStatistics() {
}