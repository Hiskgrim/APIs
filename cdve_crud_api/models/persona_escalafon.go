package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type PersonaEscalafon struct {
    Id int `orm:"column(id);pk"`
    PrimerNombre string `orm:"column(primer_nombre)"`
    SegundoNombre string `orm:"column(segundo_nombre)"`
    PrimerApellido string `orm:"column(primer_apellido)"`
    SegundoApellido string `orm:"column(segundo_apellido)"`
    Escalafon string `orm:"column(escalafon)"`
}

func init() {
	orm.RegisterModel(new(PersonaEscalafon))
}

func GetAllPersonaEscalafon() (arregloIDs []PersonaEscalafon) {
	o := orm.NewOrm()
	var temp []PersonaEscalafon
	_, err := o.Raw("SELECT p.num_documento_persona id, primer_nombre, segundo_nombre, primer_apellido, segundo_apellido, nombre_escalafon escalafon FROM agora.informacion_persona_natural p, cdve.escalafon_persona ep, cdve.escalafon e WHERE p.num_documento_persona=ep.id_persona_natural AND ep.id_escalafon=e.id_escalafon;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}