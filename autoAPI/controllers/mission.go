package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Missions
type MissionController struct {
	beego.Controller
}

// @Title MissionInformationByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid/info [post]
func (u *MissionController) MissionInfo() {
}

// @Title SearchForMission
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /search [post]
func (u *MissionController) MissionSearch() {
}

// @Title GainMission
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid [put]
func (u *MissionController) GainMission() {
}

// @Title GetMissionStatusByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid/status [post]
func (u *MissionController) MissionStatus() {
}

// @Title SubmitMissionByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid [post]
func (u *MissionController) SubmitMission() {
}

// @Title UploadFileOnlyZip
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid/file/upload [post]
func (u *MissionController) MissionUploadFile() {
}

// @Title GiveUpMission
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid [delete]
func (u *MissionController) GiveUpMission() {
}

// @Title GetMissionFeedback
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid/feedback [post]
func (u *MissionController) MissionFeedback() {
}

// @Title CreateAndPostMission
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router / [post]
func (u *MissionController) CreateMission() {
}

// @Title ModifyMissionByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid [put] type company
func (u *MissionController) ModifyMission() {
}

// @Title DeleteMissionByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid [post] type company
func (u *MissionController) DeleteMission() {
}

// @Title GetStaticMissionTagList
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /tag/list [post]
func (u *MissionController) MissionStaticTagList() {
}

// @Title MissionStatisticsByID
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /:Mid/statistics [post]
func (u *MissionController) MissionStatistics() {
}