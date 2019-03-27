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
// @Param   loginType   query   string  true        "The login type is email or telephone"
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
// @Success 200            {object}      models.ResultOfUserLogin
// @Failure 401 {string} sign up failed
// @router /user/signup [post]
func (u *AuthController) UserSignUp() {
}

// @Title CompanySignUp
// @Description New company must sign up.
// @Param   companyName       body    string              true        "The company name"
// @Param   password          body    string              true        "The password of account"
// @Param   emailAddress      body    string              true        "The email address of company account"
// @Param   confirmCode       body    string              true        "The confirm code we send to email"
// @Success 200 {object}      models.ResultOfCompanySignUp
// @Success 401 {object}      models.SignUpError
// @router /company/signup [post]
func (u *AuthController) CompanySignUp() {
}

// @Title CompanyInformationConfirm
// @Description Company account must be comfirmed
// @Param   licenseImage      body    string              true        "The company license image"
// @Param   organizationCode  body    string              true        "The company organization code"
// @Param   cardImage         body    string              true        "The company legal ID card image"
// @Param   legalName         body    string              true        "The legal body name"
// @Param   legalNumber       body    string              true        "The legal body number"
// @Param   Cid               path    int                 true        "The company ID"
// @Success 200 {string} post information successful
// @Success 401 {object| models.PostConfirmInformationError
// @router /company/:Cid/confirmInfo [post]
func (u *AuthController) CompanyInfoConfirm() {
}

// @Title CompanyInformationPost
// @Description Post the company base information.
// @Param   companyNature       body    string                    true        "The company nature"
// @Param   contactName         body    string                    true        "The contact person name"
// @Param   contactTel          body    string                    true        "The contact person telephone number"
// @Param   licenseLocation     body    string                    true        "The company license location"
// @Param   initTags            body    models.CompanyInitTags    true        "The company init tags(max = 6)"
// @Param   Cid                 path    int                       true        "The company ID"
// @Success 200 {string} company information post successful
// @Failure 401 {string} company information post failed
// @router /company/:Cid/info [post]
func (u *AuthController) CompanyInfoPost() {
}

// @Title StartConfirmCompanyInformation
// @Description Confirm request will be posted to reviewer.
// @Param   Cid               path    int                 true        "The company ID"
// @Success 200 {string} confirm request post successful
// @Failure 401 {string} confirm request post failed
// @router /company/:Cid/info [put]
func (u *AuthController) CompanyInfoPut() {
}

// @Title ReviewerSignUp
// @Description New reviewer must sign up.
// @Param   name               body    string             true        "The reviewer name"
// @Param   password           body    string             true        "The reviewer account password"
// @Param   intro              body    string             true        "The reviewer introduction"
// @Success 200 {object} models.ResultOfReviewerSignUp
// @Failure 401 {string} sign up failed
// @router /reviewer/signup [post]
func (u *AuthController) ReviewerSignUp() {
}

// @Title CompanySignIn
// @Description Company login.
// @Param   account            body    string              true        "The company account telephone or email"
// @Param   code               body    string              true        "The confirm code"
// @Param   loginType          query   string              true        "The type to login: tel or email"
// @Success 200 {object} models.ResultOfCompanyLogin
// @Success 401 {object} models.CompanyLoginError
// @router /company/signin [post]
func (u *AuthController) CompanySignIn() {
}

// @Title ReviewerSignIn
// @Description Reviewer login
// @Param   name               body    string              true        "The reviewer name"
// @Param   password           body    string              true        "The confirm code"
// @Success 200 {object} models.ResultOfReviewerLogin
// @Success 401 {object} models.ReviewerLoginError
// @router /reviewer/signin [post]
func (u *AuthController) ReviewerSignIn() {
}

// @Title GetTelephoneConfirmCode
// @Description Get telephone confirm code
// @Param   tel        query    string  true        "The telephone number ready for confirm"
// @Success 200 {string} send code successful
// @Failure 410 {string} please wait for minute
// @router /telephone/confirm [get]
func (u *AuthController) GetTelephoneConfirmCode() {
}

// @Title GetEmailConfirmCode
// @Description Get email confirm code.
// @Param   email       body    string  true        "The telephone number ready for confirm"
// @Success 200 {string} send code successful
// @Failure 410 {string} please wait for minute
// @router /email/confirm [post]
func (u *AuthController) GetEamilConfirmCode() {
}