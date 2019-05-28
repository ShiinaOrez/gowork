package controllers

import "github.com/astaxie/beego"

// Operations about Category
type CategoryController struct {
	beego.Controller
}

// @Title CreateCategory
// @Description Create a category, provide the category name.
// @Param   Payload      body     models.CreateCategoryPayload    true    "The payload of create category."
// @Success 200          {object} models.CreateCategoryReturn
// @Failure 400          {string} bad request
// @Failure 401          {string} authentication failed
// @router  /  [post]
func CreateCategory() {

}

// @Title GetCategoryInformation
// @Description Get a category information by category id.
// @Param   Token        header   string                          true    "The token to verify."
// @Param   id           path     int                             true    "The id to query category."
// @Success 200          {object} models.CategoryAndTags          true    "The category information."
// @Failure 400          {string} bad request
// @Failure 401          {string} authentication failed
// @Failure 404          {string} category not found
// @router  /:id  [get]
func GetCategory() {

}

// @Title AddTagsToCategory
// @Description Add a list of tag to category.
// @Param   Token        header   string                          true    "The token to verify."
// @Param   id           path     int                             true    "The id to query category."
// @Param   Payload      body     models.AddTagsToCategoryPayload true    "The payload is a list of tags."
// @Success 200          {string} OK
// @Failure 400          {string} bad request
// @Failure 401          {string} authentication failed
// @Failure 404          {string} category not found
// @router  /:id/addTags  [put]         
func AddTagsToCategory() {

}

// @Title RemoveTagFromCategory
// @Description Remove tag from category.
// @Param   Token        header   string                          true    "The token to verify."
// @Param   id           path     int                             true    "The id to query category."
// @Param   Payload      body     models.RemoveTagFromCategoryPayload true  "The payload is a tag id."
// @Success 200          {string} OK
// @Failure 400          {string} bad request
// @Failure 401          {string} authentication failed
// @Failure 404          {string} category not found
// @router  /:id/removeTag  [put]
func RemoveTagFromCategory() {

}

// @Title DeleteCategory
// @Description Delete a category by id.
// @Param   Token        header   string                          true    "The token to verify."
// @Param   id           path     int                             true    "The id to query category."
// @Success 200          {string} OK
// @Failure 400          {string} bad request
// @Failure 401          {string} authentication failed
// @Failure 404          {string} category not found
// @router  /:id  [delete]
func DeleteCategory() {

}

// @Title GetCategoryForElite
// @Description Get the category list for elite search.
// @Param   Token        header   string                          true    "The token to verify."
// @Success 200          {object} models.CategoryForSearch
// @Failure 401          {string} authentication failed
// @router  /forElite  [get]
func GetCategoryForElite() {

}

// @Title GetCategoryForMission
// @Description Get the category list for elite search.
// @Param   Token        header   string                          true    "The token to verify."
// @Success 200          {object} models.CategoryForSearch
// @Failure 401          {string} authentication failed
// @router  /forMission  [get]
func GetCategoryForMission() {

}

// @Title GetCategoryForCompany
// @Description Get the category list for elite search.
// @Param   Token        header   string                          true    "The token to verify."
// @Success 200          {object} models.CategoryForSearch
// @Failure 401          {string} authentication failed
// @router  /forCompany  [get]
func GetCategoryForCompany() {

}

func Change