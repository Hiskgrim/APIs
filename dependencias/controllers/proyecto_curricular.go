package controllers

import (
  "dependencias/models"
  "github.com/astaxie/beego"
)

type ProyectoCurricularController struct {
	beego.Controller
}

func (c *ProyectoCurricularController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
}


func (c *ProyectoCurricularController) GetAll() {
	idFacultadStr := c.Ctx.Input.Param(":idFacultad")
	nivelAcademicoStr := c.Ctx.Input.Param(":nivelAcademico")
	var nivelAcademico string
	if(nivelAcademicoStr=="pregrado"){
		nivelAcademico = "3"
	}else if(nivelAcademicoStr=="posgrado"){
		nivelAcademico = "4"
	}
    listaProyectosCurriculares := models.GetAllProyectoCurricular(idFacultadStr, nivelAcademico)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = listaProyectosCurriculares
    c.ServeJSON()
}

func (c *ProyectoCurricularController) GetOne() {
	idProyectoCurricularStr := c.Ctx.Input.Param(":idProyectoCurricular")
    proyectoCurricular := models.GetOneProyectoCurricular(idProyectoCurricularStr)
    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = proyectoCurricular
    c.ServeJSON()
}