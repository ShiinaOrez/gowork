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
// @Param   PostBody 	body    models.AuthEmailSignIn true "依照验证码登录所需PostBody"
// @Success 200 {object} models.ResultOfUserLogin 
// @Failure 401 {string}
// @router /user/signin/email [post]
func (u *AuthController) UserSignInByEmail() {
}

// @Title UserSignIn
// @Description Account sign in.
// @Param   tel         body    string  true        "The telephone number of login"
// @Param	code	    body 	string	true		"The confirm code"
// @Success 200 {object} models.ResultOfUserLogin 
// @Failure 401 {string}
// @router /user/signin/tel [post]
func (u *AuthController) UserSignInByTel() {
}

// @Title UserSignUp
// @Description New user must sign up.
// @Param   real_name       body    string              true        "The real name of user"
// @Param   sex            body    bool                true        "The sex of user"
// @Param   birth_year      body    int                 true        "The birth year of user"
// @Param   birth_month     body    int                 true        "The birth month of user"
// @Param   city           body    string              true        "The city of user"
// @Param   university     body    string              true        "The university of user"
// @Param   enroll_year     body    int                 true        "The enroll year of user"
// @Param   graduate_year   body    int   	       true        "The graduate year of user"
// @Param   education_level body    string              true        "The education level of user"
// @Param   major          body    string              true        "The major of user"
// @Param   status         body    string              true        "The status of user"
// @Param   init_tags       body    models.AuthInitTags true        "The initialize tags of user"
// @Param   intro          body    string              true        "The introduction of user"
// @Success 200            {object}      models.ResultOfUserLogin
// @Failure 401 {string} sign up failed
// @router /user/signup [post]
func (u *AuthController) UserSignUp() {
}

// @Title CompanySignUp
// @Description New company must sign up.
// @Param   company_name       body    string              true        "The company name"
// @Param   password           body    string              true        "The password of account"
// @Param   email_address      body    string              true        "The email address of company account"
// @Param   confirm_code       body    string              true        "The confirm code we send to email"
// @Success 200 {object}      models.ResultOfCompanySignUp
// @Success 401 {object}      models.SignUpError
// @router /company/signup [post]
func (u *AuthController) CompanySignUp() {
}

// @Title CompanyInformationConfirm
// @Description Company account must be comfirmed
// @Param   company_name       body    string              true        "The company name"
// @Param   license_image      body    string              true        "The company license image"
// @Param   organization_code  body    string              true        "The company organization code"
// @Param   origin_image       body    string              true        "The company organization origin image"
// @Param   legal_identify     body    string              true        "The company legal's identify"
// @Param   card_image         body    string              true        "The company legal ID card image"
// @Param   legal_name         body    string              true        "The legal body name"
// @Param   legal_number       body    string              true        "The legal body number"
// @Param   Cid                path    int                 true        "The company ID"
// @Success 200 {string} post information successful
// @Success 401 {object| models.PostConfirmInformationError
// @router /company/:Cid/confirmInfo [post]
func (u *AuthController) CompanyInfoConfirm() {
}

// @Title CompanyInformationPost
// @Description Post the company base information.
// @Param   company_nature       body    string                    true        "The company nature"
// @Param   contact_name         body    string                    true        "The contact person name"
// @Param   contact_tel          body    string                    true        "The contact person telephone number"
// @Param   license_location     body    string                    true        "The company license location"
// @Param   init_tags            body    models.CompanyInitTags    true        "The company init tags(max = 6)"
// @Param   Cid                  path    int                       true        "The company ID"
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
// @Param   email              body    string             true        "The reviewer email"
// @Param   password           body    string             true        "The reviewer account password"
// @Param   intro              body    string             true        "The reviewer introduction"
// @Success 200 {object} models.ResultOfReviewerSignUp
// @Failure 401 {string} sign up failed
// @router /reviewer/signup [post]
func (u *AuthController) ReviewerSignUp() {
}

// @Title CompanySignIn
// @Description Company login.
// @Param   PostBody 	body    models.AuthEmailSignIn true "依照验证码登录所需PostBody"
// @Success 200 {object} models.ResultOfCompanyLogin
// @Success 401 {object} models.CompanyLoginError
// @router /company/signin/email [post]
func (u *AuthController) CompanySignInByEmail() {
}

// @Title CompanySignIn
// @Description Company login.
// @Param   tel                body    string              true        "The company account telephone"
// @Param   code               body    string              true        "The confirm code"
// @Success 200 {object} models.ResultOfCompanyLogin
// @Success 401 {object} models.CompanyLoginError
// @router /company/signin/tel [post]
func (u *AuthController) CompanySignInByTel() {
}

// @Title ReviewerSignIn
// @Description Reviewer login
// @Param   email              body    string              true        "The reviewer email"
// @Param   password           body    string              true        "The confirm code"
// @Success 200 {object} models.ResultOfReviewerLogin
// @Success 401 {object} models.ReviewerLoginError
// @router /reviewer/signin [post]
func (u *AuthController) ReviewerSignIn() {
}

// @Title GetTelephoneConfirmCode
// @Description Get telephone confirm code
// @Param   tel        body    string  true        "The telephone number ready for confirm"
// @Success 200 {string} send code successful
// @Failure 410 {string} please wait for minute
// @router /tel/confirm [post]
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