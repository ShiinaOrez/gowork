package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type AuthController struct {
	beego.Controller
}

// @Title UserSignIn
// @Description Account sign in.
// @Param	username		body 	string	true		"The username for login"
// @Param	password		body 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /user/signin [post]
func (u *AuthController) UserSignIn() {
}

// @Title UserSignUp
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /user/signup [post]
func (u *AuthController) UserSignUp() {
}

// @Title CompanySignUp
// @Description New company must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /company/signup [post]
func (u *AuthController) CompanySignUp() {
}

// @Title CompanyInformation
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /company/info [post]
func (u *AuthController) CompanyInfoPost() {
}

// @Title ConfirmCompanyInformation
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /company/info [post]
func (u *AuthController) CompanyInfoPut() {
}

// @Title ReviewerSignUp
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /reviewer/signup [post]
func (u *AuthController) ReviewerSignUp() {
}

// @Title CompanySignIn
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /company/signin [post]
func (u *AuthController) CompanySignIn() {
}

// @Title ReviewerSignIn
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /reviewer/signin [post]
func (u *AuthController) ReviewerSignIn() {
}
