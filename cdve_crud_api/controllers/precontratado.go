package controllers

import (
  "cdve_crud_api/models"
  "github.com/astaxie/beego"
)

type PrecontratadoController struct {
	beego.Controller
}

func (c *PrecontratadoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)

}


func (c *PrecontratadoController) GetAll() {
	idResolucion := c.Ctx.Input.Param(":idResolucion")
    listaPrecontratados := models.GetAllPrecontratado(idResolucion)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaPrecontratados
    c.ServeJSON()
}