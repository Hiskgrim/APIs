package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type InformacionPersonaNatural struct {
	IdCiudadExpedicionDocumento       float64            `orm:"column(id_ciudad_expedicion_documento)"`
	FechaExpedicionDocumento          time.Time          `orm:"column(fecha_expedicion_documento);type(date)"`
	IdNitCajaCompensacion             *CajaCompensacion  `orm:"column(id_nit_caja_compensacion);rel(fk)"`
	IdNitFondoPension                 *FondoPension      `orm:"column(id_nit_fondo_pension);rel(fk)"`
	IdNitEps                          *Eps               `orm:"column(id_nit_eps);rel(fk)"`
	IdNitArl                          *Arl               `orm:"column(id_nit_arl);rel(fk)"`
	IdCajaCompensacion                *ParametroEntidad  `orm:"column(id_caja_compensacion);rel(fk)"`
	IdFondoPension                    *ParametroEntidad  `orm:"column(id_fondo_pension);rel(fk)"`
	IdEps                             *ParametroEntidad  `orm:"column(id_eps);rel(fk)"`
	IdArl                             *ParametroEntidad  `orm:"column(id_arl);rel(fk)"`
	IdNucleoBasico                    *SniesNucleoBasico `orm:"column(id_nucleo_basico);rel(fk)"`
	DependientePadreOHermano          bool               `orm:"column(dependiente_padre_o_hermano)"`
	DependienteConyuge                bool               `orm:"column(dependiente_conyuge)"`
	DependienteHijoMas23Discapacitado bool               `orm:"column(dependiente_hijo_mas23_discapacitado)"`
	DependienteHijoMenos23Estudiando  bool               `orm:"column(dependiente_hijo_menos23_estudiando)"`
	DependienteHijoMenorEdad          bool               `orm:"column(dependiente_hijo_menor_edad)"`
	InteresViviendaAfc                float64            `orm:"column(interes_vivienda_afc);null"`
	IdEntidadBancariaAfc              *Banco             `orm:"column(id_entidad_bancaria_afc);rel(fk)"`
	NumCuentaBancariaAfc              string             `orm:"column(num_cuenta_bancaria_afc);null"`
	CuentaAhorroAfc                   bool               `orm:"column(cuenta_ahorro_afc)"`
	ValorUvtPrepagada                 float64            `orm:"column(valor_uvt_prepagada);null"`
	MedicinaPrepagada                 bool               `orm:"column(medicina_prepagada)"`
	DeclaranteRenta                   bool               `orm:"column(declarante_renta)"`
	TipoDiscapacidad                  string             `orm:"column(tipo_discapacidad);null"`
	Discapacitado                     bool               `orm:"column(discapacitado)"`
	EstadoCivil                       string             `orm:"column(estado_civil)"`
	NumeroPersonasACargo              float64            `orm:"column(numero_personas_a_cargo);null"`
	PersonasACargo                    bool               `orm:"column(personas_a_cargo)"`
	CabezaFamilia                     bool               `orm:"column(cabeza_familia)"`
	ComunidadLgbt                     bool               `orm:"column(comunidad_lgbt)"`
	GrupoEtnico                       string             `orm:"column(grupo_etnico);null"`
	Genero                            string             `orm:"column(genero);null"`
	MontoCapitalAutorizado            float64            `orm:"column(monto_capital_autorizado);null"`
	Especialidad                      string             `orm:"column(especialidad);null"`
	Profesion                         string             `orm:"column(profesion);null"`
	Perfil                            *ParametroEstandar `orm:"column(perfil);rel(fk)"`
	IdPaisNacimiento                  *Pais              `orm:"column(id_pais_nacimiento);rel(fk)"`
	Cargo                             string             `orm:"column(cargo);null"`
	SegundoNombre                     string             `orm:"column(segundo_nombre);null"`
	PrimerNombre                      string             `orm:"column(primer_nombre)"`
	SegundoApellido                   string             `orm:"column(segundo_apellido);null"`
	PrimerApellido                    string             `orm:"column(primer_apellido)"`
	DigitoVerificacion                float64            `orm:"column(digito_verificacion)"`
	Id                                int                `orm:"column(num_documento_persona);pk"`
	TipoDocumento                     *ParametroEstandar `orm:"column(tipo_documento);rel(fk)"`
}

func (t *InformacionPersonaNatural) TableName() string {
	return "informacion_persona_natural"
}

func init() {
	orm.RegisterModel(new(InformacionPersonaNatural))
}

// AddInformacionPersonaNatural insert a new InformacionPersonaNatural into database and returns
// last inserted Id on success.
func AddInformacionPersonaNatural(m *InformacionPersonaNatural) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInformacionPersonaNaturalById retrieves InformacionPersonaNatural by Id. Returns error if
// Id doesn't exist
func GetInformacionPersonaNaturalById(id int) (v *InformacionPersonaNatural, err error) {
	o := orm.NewOrm()
	v = &InformacionPersonaNatural{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInformacionPersonaNatural retrieves all InformacionPersonaNatural matches certain condition. Returns empty list if
// no records exist
func GetAllInformacionPersonaNatural(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InformacionPersonaNatural))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []InformacionPersonaNatural
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateInformacionPersonaNatural updates InformacionPersonaNatural by Id and returns error if
// the record to be updated doesn't exist
func UpdateInformacionPersonaNaturalById(m *InformacionPersonaNatural) (err error) {
	o := orm.NewOrm()
	v := InformacionPersonaNatural{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInformacionPersonaNatural deletes InformacionPersonaNatural by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInformacionPersonaNatural(id int) (err error) {
	o := orm.NewOrm()
	v := InformacionPersonaNatural{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&InformacionPersonaNatural{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
