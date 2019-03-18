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
// @Param	email		body 	string	true		"The email for login"
// @Param   tel         body    string  true        "The telephone number of login"
// @Param	password	body 	string	true		"The password for login"
// @Param   loginType   path    string  true        "The login type is email or telephone"
// @Success 200 {string} login success
// @Failure 401 {string}
// @router /user/signin [post]
func (u *AuthController) UserSignIn() {
}

// @Title UserSignUp
// @Description New user must sign up.
// @Param   realName       body    string              true        "The real name of user"
// @Param   sex            body    int                 true        "The sex of user"
// @Param   birthYear      body    string              true        "The birth year of user"
// @Param   birthMonth     body    string              true        "The birth month of user"
// @Param   city           body    string              true        "The city of user"
// @Param   university     body    string              true        "The university of user"
// @Param   enrollYear     body    string              true        "The enroll year of user"
// @Param   graduateYear   body    string  	           true        "The graduate year of user"
// @Param   educationLevel body    string              true        "The education level of user"
// @Param   major          body    string              true        "The major of user"
// @Param   status         body    string              true        "The status of user"
// @Param   initTags       body    models.AuthInitTags true        "The initialize tags of user"
// @Param   intro          body    string              true        "The introduction of user"
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

// @Title GetTelephoneConfirmCode
// @Description New user must sign up.
// @Success 200 {string} sign up successful
// @Failure 401 {string} sign up failed
// @router /telephone/confirm [post]
func (u *AuthController) GetTelephoneConfirmCode() {
}