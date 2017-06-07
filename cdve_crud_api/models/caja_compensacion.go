package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CajaCompensacion struct {
	Estado                   string  `orm:"column(estado)"`
	Email                    string  `orm:"column(email)"`
	NombreRepresentanteLegal string  `orm:"column(nombre_representante_legal);null"`
	ExtencionFax             float64 `orm:"column(extencion_fax);null"`
	Fax                      float64 `orm:"column(fax);null"`
	ExtencionTelefono        float64 `orm:"column(extencion_telefono);null"`
	Telefono                 float64 `orm:"column(telefono)"`
	Direccion                string  `orm:"column(direccion)"`
	Nombre                   string  `orm:"column(nombre)"`
	IdUbicacion              float64 `orm:"column(id_ubicacion);null"`
	Id                       int     `orm:"column(nit);pk"`
}

func (t *CajaCompensacion) TableName() string {
	return "caja_compensacion"
}

func init() {
	orm.RegisterModel(new(CajaCompensacion))
}

// AddCajaCompensacion insert a new CajaCompensacion into database and returns
// last inserted Id on success.
func AddCajaCompensacion(m *CajaCompensacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCajaCompensacionById retrieves CajaCompensacion by Id. Returns error if
// Id doesn't exist
func GetCajaCompensacionById(id int) (v *CajaCompensacion, err error) {
	o := orm.NewOrm()
	v = &CajaCompensacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCajaCompensacion retrieves all CajaCompensacion matches certain condition. Returns empty list if
// no records exist
func GetAllCajaCompensacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CajaCompensacion))
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

	var l []CajaCompensacion
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

// UpdateCajaCompensacion updates CajaCompensacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateCajaCompensacionById(m *CajaCompensacion) (err error) {
	o := orm.NewOrm()
	v := CajaCompensacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCajaCompensacion deletes CajaCompensacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCajaCompensacion(id int) (err error) {
	o := orm.NewOrm()
	v := CajaCompensacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CajaCompensacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
