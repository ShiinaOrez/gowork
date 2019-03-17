package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:AuthController"],
        beego.ControllerComments{
            Method: "UserSignIn",
            Router: `/user/login`,
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

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ShiinaOrez/gowork/autoAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
