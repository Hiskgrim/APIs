package controllers

import (
  "cdve_crud_api/models"
  "github.com/astaxie/beego"
)

type ContratadoController struct {
	beego.Controller
}

func (c *ContratadoController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)

}


func (c *ContratadoController) GetAll() {
	tipoDedicacionStr := c.Ctx.Input.Param(":tipoDedicacion")
    listaContratados := models.GetAllContratado(tipoDedicacionStr)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaContratados
    c.ServeJSON()
}