package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Tags
type TagController struct {
	beego.Controller
}

// @Title TagIndustryList
// @Description Get industry tag list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {string} none tag exist
// @router /industry/list [get]
func (u *TagController) TagIndustryList() {
}

// @Title TagPositionList
// @Description Get position tag list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {object} none tag exist
// @router /position/list [get]
func (u *TagController) TagPositionList() {
}

// @Title TagUniversityList
// @Description Get university tags list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {string} none tag exist
// @router /university/list [get]
func (u *TagController) TagUniversityList() {
}

// @Title TagAbilityList
// @Description Get ability tags list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {string} none tag exist
// @router /ability/list [get]
func (u *TagController) TagAbilityList() {
}

// @Title TagCharacterList
// @Description Get character tags list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {string} none tag exist
// @router /character/list [get]
func (u *TagController) TagCharacterList() {
}

// @Title TagGradeList
// @Description Get grade tags list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {string} none tag exist
// @router /grade/list [get]
func (u *TagController) TagGradeList() {
}

// @Title TagAddressList
// @Description Get address tags list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {string} none tag exist
// @router /address/list [get]
func (u *TagController) TagAddressList() {
}

// @Title TagWelfareList
// @Description Get welfare tags list.
// @Param   page    query    int        true        "The page number."
// @Success 200 {object} models.TagList
// @Success 201 {string} none tag exist
// @router /welfare/list [get]
func (u *TagController) TagWelfareList() {
}
