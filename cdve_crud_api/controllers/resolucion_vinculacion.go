package controllers

import (
  "cdve_crud_api/models"
  "github.com/astaxie/beego"
)

type ResolucionVinculacionController struct {
	beego.Controller
}

func (c *ResolucionVinculacionController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)

}


func (c *ResolucionVinculacionController) GetAll() {
    listaResoluciones := models.GetAllResolucionVinculacion()
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaResoluciones
    c.ServeJSON()
}