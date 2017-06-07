package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaPadreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:DependenciaTipoDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"] = append(beego.GlobalControllerRouter["dependencias/controllers:TipoDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
