package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ContratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:tipoDedicacion`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PrecontratadoController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PrecontratadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["cdve_crud_api/controllers:PersonaEscalafonController"] = append(beego.GlobalControllerRouter["cdve_crud_api/controllers:PersonaEscalafonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})


}