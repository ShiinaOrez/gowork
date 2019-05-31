package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanyInfoConfirm",
            Router: `/company/:Cid/confirmInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanyInfoPut",
            Router: `/company/:Cid/info`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanyInfoPost",
            Router: `/company/:Cid/info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanySignInByEmail",
            Router: `/company/signin/email`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanySignInByTel",
            Router: `/company/signin/tel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanySignUp",
            Router: `/company/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "GetEamilConfirmCode",
            Router: `/email/confirm`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "ReviewerSignIn",
            Router: `/reviewer/signin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "ReviewerSignUp",
            Router: `/reviewer/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "GetTelephoneConfirmCode",
            Router: `/tel/confirm`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "UserSignInByEmail",
            Router: `/user/signin/email`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "UserSignInByTel",
            Router: `/user/signin/tel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "UserSignUp",
            Router: `/user/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "CategoryCreate",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetCategoryInfo",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "DeleteCategory",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "AddTagToCategory",
            Router: `/:id/addTags`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "RemoveTagFromCategory",
            Router: `/:id/removeTag`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "CategoryForCompany",
            Router: `/forCompany`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "CategoryForElite",
            Router: `/forElite`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "CategoryForMission",
            Router: `/forMission`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyCancelAttention",
            Router: `/attention`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyPayAttention",
            Router: `/attention`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyAttentionList",
            Router: `/attention/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CreateCompanyCategory",
            Router: `/category`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyCategoryAddTag",
            Router: `/category/:id/tags`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "EditCompanyInfo",
            Router: `/info`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyInfo",
            Router: `/info/:id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyMessageSetting",
            Router: `/info/email/allow`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompanyEmailStep1",
            Router: `/info/email/step1`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompanyEmailStep2",
            Router: `/info/email/step2`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompanyLogo",
            Router: `/info/logo`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompanyTelStep1",
            Router: `/info/tel/step1`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompanyTelStep2",
            Router: `/info/tel/step2`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "GetCompanyMissionList",
            Router: `/mission/:id/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyMissionNumberOverview",
            Router: `/mission/:id/overview`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "GetFeedbackList",
            Router: `/mission/{id}/submissions/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CreateCompanyMissionList",
            Router: `/missionList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "DeleteCompanyMissionList",
            Router: `/missionList/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyAllMissionList",
            Router: `/missionList/:id/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "RemoveMissionInMissionList",
            Router: `/missionList/mission`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "AddMissionToMissionList",
            Router: `/missionList/mission`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyChangePassword",
            Router: `/password/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanySearchFast",
            Router: `/search/fast`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanySearchIntelligent",
            Router: `/search/intelligent`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:EliteController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:EliteController"],
        beego.ControllerComments{
            Method: "SearchEliteIntelligent",
            Router: `/industry/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:EliteController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:EliteController"],
        beego.ControllerComments{
            Method: "SearchElite",
            Router: `/search`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetMessageInfo",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "ReadMessage",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "DeleteMessage",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetPlatformMessageList",
            Router: `/:id/platform/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetCompanyMessageList",
            Router: `/company/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetFeedbackMessageList",
            Router: `/feedback/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "CreateMission",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "GiveUpMission",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "DeleteMission",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "GainMission",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "ModifyMission",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "SubmitMission",
            Router: `/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "SubmissionFeedback",
            Router: `/:id/feedback`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionInfo",
            Router: `/:id/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionStatistics",
            Router: `/:id/statistics`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionStatus",
            Router: `/:id/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionStaticTagList",
            Router: `/:id/tag/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionUploadFile",
            Router: `/file/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionSearch",
            Router: `/search`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "GetMissionListInformation",
            Router: `/:MlID`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "DelMissionList",
            Router: `/:MlID`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "GetEliteByMissionList",
            Router: `/:MlID/getElite`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "AttentionMissionList",
            Router: `/attention/:MlID`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "CancelAttentionMissionList",
            Router: `/attention/:MlID`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "GetMissionListStatus",
            Router: `/attention/:MlID/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "GetMissionListAttentionList",
            Router: `/attention/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "MissionListSearch",
            Router: `/search`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "CreateTag",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "DeleteTag",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "UpdateTag",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "GetTag",
            Router: `/:id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserAttentionCompany",
            Router: `/attention/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserCancelAttentionCompany",
            Router: `/attention/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "AttentionCompanyStatu",
            Router: `/attention/:id/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserAttentionList",
            Router: `/attentions`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateAvatar",
            Router: `/avatar`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetSuitabilityList",
            Router: `/company/suitability/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "ConfirmEmail",
            Router: `/email/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "EditSetting",
            Router: `/info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserBaseInfo",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfo",
            Router: `/info/:ID`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfoInCompanyView",
            Router: `/infoCompany/:ID`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserMissionInfo",
            Router: `/mission/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/password`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "PersonalSetting",
            Router: `/personal/setting`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "EditPersonalSetting",
            Router: `/personal/setting`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "ConfirmTel",
            Router: `/tel/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
