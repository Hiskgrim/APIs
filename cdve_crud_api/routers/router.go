// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cdve_crud_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/informacion_persona_natural",
			beego.NSInclude(
				&controllers.InformacionPersonaNaturalController{},
			),
		),

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/arl",
			beego.NSInclude(
				&controllers.ArlController{},
			),
		),

		beego.NSNamespace("/tipo_resolucion",
			beego.NSInclude(
				&controllers.TipoResolucionController{},
			),
		),

		beego.NSNamespace("/escalafon_persona",
			beego.NSInclude(
				&controllers.EscalafonPersonaController{},
			),
		),

		beego.NSNamespace("/fondo_pension",
			beego.NSInclude(
				&controllers.FondoPensionController{},
			),
		),

		beego.NSNamespace("/parametro_entidad",
			beego.NSInclude(
				&controllers.ParametroEntidadController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),

		beego.NSNamespace("/componente_resolucion",
			beego.NSInclude(
				&controllers.ComponenteResolucionController{},
			),
		),

		beego.NSNamespace("/pais",
			beego.NSInclude(
				&controllers.PaisController{},
			),
		),

		beego.NSNamespace("/departamento",
			beego.NSInclude(
				&controllers.DepartamentoController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion_docente",
			beego.NSInclude(
				&controllers.ResolucionVinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/vinculacion_docente",
			beego.NSInclude(
				&controllers.VinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/caja_compensacion",
			beego.NSInclude(
				&controllers.CajaCompensacionController{},
			),
		),

		beego.NSNamespace("/punto_salarial",
			beego.NSInclude(
				&controllers.PuntoSalarialController{},
			),
		),

		beego.NSNamespace("/escalafon",
			beego.NSInclude(
				&controllers.EscalafonController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/lugar_ejecucion",
			beego.NSInclude(
				&controllers.LugarEjecucionController{},
			),
		),

		beego.NSNamespace("/eps",
			beego.NSInclude(
				&controllers.EpsController{},
			),
		),

		beego.NSNamespace("/contrato",
			beego.NSInclude(
				&controllers.ContratoController{},
			),
		),

		beego.NSNamespace("/contrato_estado",
			beego.NSInclude(
				&controllers.ContratoEstadoController{},
			),
		),

		beego.NSNamespace("/contrato_general",
			beego.NSInclude(
				&controllers.ContratoGeneralController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/supervisor_contrato",
			beego.NSInclude(
				&controllers.SupervisorContratoController{},
			),
		),

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

		beego.NSNamespace("/parametros",
			beego.NSInclude(
				&controllers.ParametrosController{},
			),
		),

		beego.NSNamespace("/snies_nucleo_basico",
			beego.NSInclude(
				&controllers.SniesNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/contratado",
			beego.NSInclude(
				&controllers.ContratadoController{},
			),
		),

		beego.NSNamespace("/precontratado",
			beego.NSInclude(
				&controllers.PrecontratadoController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion",
			beego.NSInclude(
				&controllers.ResolucionVinculacionController{},
			),
		),

		beego.NSNamespace("/persona_escalafon",
			beego.NSInclude(
				&controllers.PersonaEscalafonController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
