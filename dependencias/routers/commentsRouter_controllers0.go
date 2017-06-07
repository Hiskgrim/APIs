package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dependencias/controllers:FacultadController"] = append(beego.GlobalControllerRouter["dependencias/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:FacultadController"] = append(beego.GlobalControllerRouter["dependencias/controllers:FacultadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idFacultad`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["dependencias/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:nivelAcademico/:idFacultad`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:ProyectoCurricularController"] = append(beego.GlobalControllerRouter["dependencias/controllers:ProyectoCurricularController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idProyectoCurricular`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
