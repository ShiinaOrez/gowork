// @APIVersion 0.0.1
// @Title Elite API
// @Description Elite is a application which due to deal with the relationship between enterprises and students
// @Contact Shiinaorez@gmail.com
// @TermsOfServiceUrl http://www.muxixyz.com
package routers

import (
	"github.com/ShiinaOrez/gowork/autoAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1.0",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/company",
			beego.NSInclude(
				&controllers.CompanyController{},
			),
		),
		beego.NSNamespace("/mission",
			beego.NSInclude(
				&controllers.MissionController{},
			),
		),
		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
		beego.NSNamespace("/elite",
			beego.NSInclude(
				&controllers.EliteController{},
			),
		),
		beego.NSNamespace("/missionList",
			beego.NSInclude(
				&controllers.MissionListController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
