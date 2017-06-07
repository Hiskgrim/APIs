package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:DominioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:PredicadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"] = append(beego.GlobalControllerRouter["rulerAPI/controllers:TipoPredicadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
