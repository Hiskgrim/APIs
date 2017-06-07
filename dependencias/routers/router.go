// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dependencias/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&controllers.DependenciaController{},
			),
		),

		beego.NSNamespace("/dependencia_tipo_dependencia",
			beego.NSInclude(
				&controllers.DependenciaTipoDependenciaController{},
			),
		),

		beego.NSNamespace("/dependencia_padre",
			beego.NSInclude(
				&controllers.DependenciaPadreController{},
			),
		),

		beego.NSNamespace("/tipo_dependencia",
			beego.NSInclude(
				&controllers.TipoDependenciaController{},
			),
		),
		beego.NSNamespace("/facultad",
			beego.NSInclude(
				&controllers.FacultadController{},
			),
		),
		beego.NSNamespace("/proyecto_curricular",
			beego.NSInclude(
				&controllers.ProyectoCurricularController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
