package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ResolucionVinculacion struct {
    Id int `orm:"column(id);pk"`
    Numero int `orm:"column(numero)"`
    Vigencia int `orm:"column(vigencia)"`
    Facultad string `orm:"column(facultad)"`
    NivelAcademico string `orm:"column(nivel_academico)"`
    Dedicacion string `orm:"column(dedicacion)"`
}

func init() {
	orm.RegisterModel(new(ResolucionVinculacion))
}

func GetAllResolucionVinculacion() (arregloIDs []ResolucionVinculacion) {
	o := orm.NewOrm()
	var temp []ResolucionVinculacion
	_, err := o.Raw("SELECT r.id_resolucion id, r.numero_resolucion numero, r.vigencia vigencia, d.nombre facultad, rv.nivel_academico nivel_academico, rv.dedicacion dedicacion FROM cdve.resolucion r, cdve.resolucion_vinculacion_docente rv, oikos.dependencia d WHERE rv.id_facultad=d.id AND r.id_resolucion=rv.id_resolucion;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}