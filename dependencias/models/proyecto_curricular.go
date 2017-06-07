package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ProyectoCurricular struct {
    Id string `orm:"column(id);pk"`
    Nombre string `orm:"column(nombre)"`
}

func init() {
	orm.RegisterModel(new(ProyectoCurricular))
}

func GetAllProyectoCurricular(idFacultadStr string, nivelAcademico string) (arregloIDs []ProyectoCurricular) {
	o := orm.NewOrm()
	var temp []ProyectoCurricular
	_, err := o.Raw("SELECT oikos.dependencia.id, oikos.dependencia.nombre FROM oikos.dependencia, oikos.dependencia_tipo_dependencia tipo_dependencia_1, oikos.dependencia_tipo_dependencia tipo_dependencia_2, oikos.dependencia_padre WHERE oikos.dependencia.id = tipo_dependencia_1.dependencia_id AND tipo_dependencia_1.tipo_dependencia_id=2 AND oikos.dependencia.id=oikos.dependencia_padre.hija AND oikos.dependencia_padre.padre="+idFacultadStr+" AND oikos.dependencia.id = tipo_dependencia_2.dependencia_id AND tipo_dependencia_2.tipo_dependencia_id="+nivelAcademico+";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}

func GetOneProyectoCurricular(idProyectoCurricularStr string) (arregloIDs ProyectoCurricular) {
	o := orm.NewOrm()
	var temp []ProyectoCurricular
	_, err := o.Raw("SELECT oikos.dependencia.id, oikos.dependencia.nombre FROM oikos.dependencia, oikos.dependencia_tipo_dependencia WHERE oikos.dependencia.id = oikos.dependencia_tipo_dependencia.dependencia_id and oikos.dependencia_tipo_dependencia.tipo_dependencia_id=2 AND oikos.dependencia.id="+idProyectoCurricularStr+";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp[0]
}