package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Contratado struct {
    Id string `orm:"column(id);pk"`
    Nombre string `orm:"column(nombre)"`
    Documento string `orm:"column(documento)"`
    Expedicion string `orm:"column(expedicion)"`
    Categoria string `orm:"column(categoria)"`
    Dedicacion string `orm:"column(dedicacion)"`
    HorasSemanales int `orm:"column(horas_semanales)"`
    PeriodoVinculacion int `orm:"column(periodo_vinculacion)"`
    ValorContrato float64 `orm:"column(valor_contrato)"`
    IdProyectoCurricular int `orm:"column(id_proyecto_curricular)"`
}

func init() {
	orm.RegisterModel(new(Contratado))
}

func GetAllContratado(tipoDedicacion string) (arregloIDs []Contratado) {
	o := orm.NewOrm()
	var temp []Contratado
	_, err := o.Raw("SELECT agora.informacion_persona_natural.num_documento_persona id, (agora.informacion_persona_natural.primer_nombre || ' ' || agora.informacion_persona_natural.segundo_nombre || ' ' || agora.informacion_persona_natural.primer_apellido || ' ' || agora.informacion_persona_natural.segundo_apellido) nombre, agora.informacion_persona_natural.num_documento_persona documento, core.ciudad.nombre expedicion, cdve.escalafon.nombre_escalafon categoria, cdve.dedicacion.nombre_dedicacion dedicacion, cdve.contrato_docente.numero_horas_semanales horas_semanales, cdve.contrato_docente.numero_semanas periodo_vinculacion, argo.contrato_general.valor_contrato valor_contrato, cdve.contrato_docente.id_proyecto_curricular id_proyecto_curricular FROM agora.informacion_persona_natural, core.ciudad, cdve.escalafon_persona, cdve.escalafon, cdve.dedicacion, cdve.contrato_docente, argo.contrato_general WHERE agora.informacion_persona_natural.id_ciudad_expedicion_documento=core.ciudad.id_ciudad AND agora.informacion_persona_natural.num_documento_persona=cdve.contrato_docente.id_persona AND cdve.contrato_docente.numero_contrato=argo.contrato_general.numero_contrato AND cdve.contrato_docente.vigencia=argo.contrato_general.vigencia AND cdve.escalafon_persona.id_persona_natural=agora.informacion_persona_natural.num_documento_persona AND cdve.contrato_docente.id_dedicacion=cdve.dedicacion.id_dedicacion AND cdve.escalafon_persona.id_escalafon=cdve.escalafon.id_escalafon AND cdve.dedicacion.nombre_dedicacion like '"+tipoDedicacion+"';").QueryRows(&temp)
	if err == nil {
		fmt.Println("Consulta exitosa")
	}

	return temp
}