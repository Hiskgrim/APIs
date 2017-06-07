package controllers

import (
  "dependencias/models"
  "github.com/astaxie/beego"
)

type FacultadController struct {
	beego.Controller
}

func (c *FacultadController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
}


func (c *FacultadController) GetAll() {
    listaFacultades := models.GetAllFacultad()
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaFacultades
    c.ServeJSON()
}

func (c *FacultadController) GetOne() {
	idFacultadStr := c.Ctx.Input.Param(":idFacultad")
    facultad := models.GetOneFacultad(idFacultadStr)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = facultad
    c.ServeJSON()
}