package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ArlController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:CajaCompensacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:DepartamentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EpsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:EscalafonPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:FondoPensionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:LugarEjecucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PaisController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
