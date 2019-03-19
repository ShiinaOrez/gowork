package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanyInfoConfirm",
            Router: `/company/:Cid/info`,
            AllowHTTPMethods: []string{"post"},
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
            Method: "CompanyInfoPut",
            Router: `/company/:Cid/info`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CompanySignIn",
            Router: `/company/signin`,
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
            Router: `/telephone/confirm`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "UserSignIn",
            Router: `/user/signin`,
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

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyInfo",
            Router: `/:Cid/info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "GetCompanyPositionList",
            Router: `/:Cid/position/list`,
            AllowHTTPMethods: []string{"post"},
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
            Method: "CompanyCancelAttention",
            Router: `/attention`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyAttentionList",
            Router: `/attention/list`,
            AllowHTTPMethods: []string{"post"},
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
            Router: `/category/:CateID/tags`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompanyEmail",
            Router: `/email`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "ConfirmCompanyEmail",
            Router: `/email/confirm`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyBaseInfo",
            Router: `/info`,
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
            Method: "UpdateCompanyLogo",
            Router: `/logo`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "GetCompanyMissionList",
            Router: `/mission/list`,
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
            Method: "CreateCompanyOffer",
            Router: `/offer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "DeleteCompanyOffer",
            Router: `/offer/:Oid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "GetCompanyOfferInfo",
            Router: `/offer/:Oid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CreateCompanyPosition",
            Router: `/position`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "EditCompanyPosition",
            Router: `/position/:Pid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "DeleteCompanyPosition",
            Router: `/position/:Pid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "ResumeToCompanyPosition",
            Router: `/position/:Pid/resume`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "CompanyPositionSearch",
            Router: `/position/:Pid/search`,
            AllowHTTPMethods: []string{"post"},
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
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:CompanyController"],
        beego.ControllerComments{
            Method: "UpdateCompanyTel",
            Router: `/tel`,
            AllowHTTPMethods: []string{"put"},
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
            Router: `/:MsgID`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "ReadMessage",
            Router: `/:MsgID`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetCompanyMessageList",
            Router: `/company/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetFeedbackMessageList",
            Router: `/feedback/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MessageController"],
        beego.ControllerComments{
            Method: "GetPlatformMessageList",
            Router: `/platform/list`,
            AllowHTTPMethods: []string{"post"},
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
            Method: "DeleteMission",
            Router: `/:Mid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "GainMission",
            Router: `/:Mid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "ModifyMission",
            Router: `/:Mid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "SubmitMission",
            Router: `/:Mid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "GiveUpMission",
            Router: `/:Mid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionFeedback",
            Router: `/:Mid/feedback`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionUploadFile",
            Router: `/:Mid/file/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionInfo",
            Router: `/:Mid/info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionStatistics",
            Router: `/:Mid/statistics`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionStatus",
            Router: `/:Mid/status`,
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

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionController"],
        beego.ControllerComments{
            Method: "MissionStaticTagList",
            Router: `/tag/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "GetMissionListInformation",
            Router: `/:MlID`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "DelMissionList",
            Router: `/:MlID`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "GetEliteByMissionList",
            Router: `/:MlID/getElite`,
            AllowHTTPMethods: []string{"post"},
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
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:MissionListController"],
        beego.ControllerComments{
            Method: "GetMissionListAttentionList",
            Router: `/attention/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "TagAbilityList",
            Router: `/ability/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "TagIndustryList",
            Router: `/industry/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "TagPositionList",
            Router: `/position/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:TagController"],
        beego.ControllerComments{
            Method: "TagUniversityList",
            Router: `/university/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfo",
            Router: `/:Uid/info`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserAttentionCompany",
            Router: `/attention/:Cid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserCancelAttentionCompany",
            Router: `/attention/:Cid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "AttentionCompanyStatu",
            Router: `/attention/:Cid/status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserAttentionList",
            Router: `/attention/list`,
            AllowHTTPMethods: []string{"post"},
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
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserEditPassword",
            Router: `/editPassword`,
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
            Method: "EditSetting",
            Router: `/info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserMissionInfo",
            Router: `/mission/info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "PlatformSetting",
            Router: `/platform/setting`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "EditPlatformSetting",
            Router: `/platform/setting`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}