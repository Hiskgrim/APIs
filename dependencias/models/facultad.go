package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Facultad struct {
    Id string `orm:"column(id);pk"`
    Nombre string `orm:"column(nombre)"`
}

func init() {
	orm.RegisterModel(new(Facultad))
}

func GetAllFacultad() (arregloIDs []Facultad) {
	o := orm.NewOrm()
	var temp []Facultad
	_, err := o.Raw("SELECT oikos.dependencia.id, oikos.dependencia.nombre FROM oikos.dependencia, oikos.dependencia_tipo_dependencia WHERE oikos.dependencia.id = oikos.dependencia_tipo_dependencia.dependencia_id and oikos.dependencia_tipo_dependencia.tipo_dependencia_id=1;").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}

func GetOneFacultad(idFacultadStr string) (arregloIDs Facultad) {
	o := orm.NewOrm()
	var temp []Facultad
	_, err := o.Raw("SELECT oikos.dependencia.id, oikos.dependencia.nombre FROM oikos.dependencia, oikos.dependencia_tipo_dependencia WHERE oikos.dependencia.id = oikos.dependencia_tipo_dependencia.dependencia_id and oikos.dependencia_tipo_dependencia.tipo_dependencia_id=1 AND oikos.dependencia.id="+idFacultadStr+";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp[0]
}