// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/agenda_v1/Agenda_V1/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/contacto",
			beego.NSInclude(
				&controllers.ContactoController{},
			),
		),

		beego.NSNamespace("/telefono",
			beego.NSInclude(
				&controllers.TelefonoController{},
			),
		),

		beego.NSNamespace("/correo",
			beego.NSInclude(
				&controllers.CorreoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
