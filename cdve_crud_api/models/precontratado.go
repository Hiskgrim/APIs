package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Precontratado struct {
    Id string `orm:"column(id);pk"`
    PrimerNombre string `orm:"column(primer_nombre)"`
    SegundoNombre string `orm:"column(segundo_nombre)"`
    PrimerApellido string `orm:"column(primer_apellido)"`
    SegundoApellido string `orm:"column(segundo_apellido)"`
    Documento int `orm:"column(documento)"`
    Expedicion string `orm:"column(expedicion)"`
    Categoria string `orm:"column(categoria)"`
    Dedicacion string `orm:"column(dedicacion)"`
    HorasSemanales int `orm:"column(horas_semanales)"`
    Semanas int `orm:"column(semanas)"`
    ProyectoCurricular int `orm:"column(proyecto_curricular)"`
}

func init() {
	orm.RegisterModel(new(Precontratado))
}

func GetAllPrecontratado(idResolucion string) (arregloIDs []Precontratado) {
	o := orm.NewOrm()
	var temp []Precontratado
	_, err := o.Raw("SELECT cdve.vinculacion_docente.id id, agora.informacion_persona_natural.primer_nombre primer_nombre, agora.informacion_persona_natural.segundo_nombre segundo_nombre, agora.informacion_persona_natural.primer_apellido primer_apellido, agora.informacion_persona_natural.segundo_apellido segundo_apellido, agora.informacion_persona_natural.num_documento_persona documento, core.ciudad.nombre expedicion, cdve.escalafon.nombre_escalafon categoria, cdve.dedicacion.nombre_dedicacion dedicacion, cdve.vinculacion_docente.numero_horas_semanales horas_semanales, cdve.vinculacion_docente.numero_semanas semanas, cdve.vinculacion_docente.id_proyecto_curricular proyecto_curricular FROM agora.informacion_persona_natural, core.ciudad, cdve.escalafon_persona, cdve.escalafon, cdve.dedicacion, cdve.vinculacion_docente WHERE agora.informacion_persona_natural.id_ciudad_expedicion_documento=core.ciudad.id_ciudad AND agora.informacion_persona_natural.num_documento_persona=cdve.vinculacion_docente.id_persona AND cdve.escalafon_persona.id_persona_natural=agora.informacion_persona_natural.num_documento_persona AND cdve.vinculacion_docente.id_dedicacion=cdve.dedicacion.id_dedicacion AND cdve.escalafon_persona.id_escalafon=cdve.escalafon.id_escalafon AND cdve.vinculacion_docente.estado=TRUE AND cdve.vinculacion_docente.id_resolucion = "+idResolucion+";").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}