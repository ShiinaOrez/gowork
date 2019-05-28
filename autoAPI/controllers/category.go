package controllers

import "github.com/astaxie/beego"

// Operations about Category
type CategoryController struct {
	beego.Controller
}

// @Title CreateCategory
// @Description Create a category, provide the category name.
// @Param   Payload      body     models.CreateCategoryPayload    true    "The payload of create category."
// @Param   Token       header    string      true        "The token to conform"
// @Success 200          {object} models.CreateCategoryReturn
// @Failure 400          {string} bad request
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router / [post]
func (u *CategoryController) CategoryCreate() {
}

// @Title CategoryGetInformation
// @Description Get a category information by category id.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id           path     int                             true    "The id to query category."
// @Success 200  {object} models.GetCategoryReturn
// @Failure 401 {string} auth failed
// @Failure 404 {string} user not existed
// @router /:id [get]
func (u *CategoryController) GetCategoryInfo() {
}



// @Title  AddTagsToCategory
// @Description Add a list of tag to category.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id           path     int                             true    "The id to query category."
// @Param   Payload      body     models.AddTagsToCategoryPayload true    "The payload is a list of tags."
// @Success 200 {string} OK
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /:id/addTags [put]
func (u *CategoryController) AddTagToCategory() {
}


// @Title RemoveTagFromCategory
// @Description  Remove tag from category.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The id to query category."
// @Param   Payload      body     models.RemoveTagFromCategoryPayload true  "The payload is a tag id."
// @Success 200 {string} OK
// @Failure 400          {string} bad request
// @Failure 401 {string} auth failed
// @Failure 404 {string} category not found
// @router /:id/removeTag [put]
func (u *CategoryController) RemoveTagFromCategory() {
}

// @Title DeleteCategory
// @Description Delete a category by id.
// @Param   Token       header    string      true        "The token to conform"
// @Param   id         path      int         true        "The id to query category."
// @Success 200 {string} OK
// @Failure 400          {string} bad request
// @Failure 401 {string} auth failed
// @Failure 404 {string} user | company not existed
// @router /:id [delete]
func (u *CategoryController) DeleteCategory() {
}


// @Title GetCategoryForElite
// @Description  获取用于人才搜索的Category & tag
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CategoryForSearch
// @Failure 401 {string} auth failed
// @router /forElite [get]
func (u *CategoryController) CategoryForElite() {
}

// @Title  GetCategoryForMission
// @Description 获取用于任务搜索的Category & tag
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CategoryForSearch
// @Failure 401 {string} auth failed
// @router /forMission [get]
func (u *CategoryController) CategoryForMission() {
}


// @Title GetCategoryForCompany
// @Description 获取用于公司搜索的Category & tag
// @Param   Token       header    string      true        "The token to conform"
// @Success 200 {object} models.CompanySearchList
// @Failure 401 {string} auth failed
// @router /forCompany [get]
func (u *CategoryController) CategoryForCompany() {
}