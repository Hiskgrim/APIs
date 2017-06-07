package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PuntoSalarial struct {
	Decreto       string    `orm:"column(decreto);null"`
	ValorPunto    float64   `orm:"column(valor_punto)"`
	FechaRegistro time.Time `orm:"column(fecha_registro);type(date)"`
	Vigencia      int       `orm:"column(vigencia)"`
	Id            int       `orm:"column(id);pk"`
}

func (t *PuntoSalarial) TableName() string {
	return "punto_salarial"
}

func init() {
	orm.RegisterModel(new(PuntoSalarial))
}

// AddPuntoSalarial insert a new PuntoSalarial into database and returns
// last inserted Id on success.
func AddPuntoSalarial(m *PuntoSalarial) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPuntoSalarialById retrieves PuntoSalarial by Id. Returns error if
// Id doesn't exist
func GetPuntoSalarialById(id int) (v *PuntoSalarial, err error) {
	o := orm.NewOrm()
	v = &PuntoSalarial{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPuntoSalarial retrieves all PuntoSalarial matches certain condition. Returns empty list if
// no records exist
func GetAllPuntoSalarial(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PuntoSalarial))
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

	var l []PuntoSalarial
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

// UpdatePuntoSalarial updates PuntoSalarial by Id and returns error if
// the record to be updated doesn't exist
func UpdatePuntoSalarialById(m *PuntoSalarial) (err error) {
	o := orm.NewOrm()
	v := PuntoSalarial{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePuntoSalarial deletes PuntoSalarial by Id and returns error if
// the record to be deleted doesn't exist
func DeletePuntoSalarial(id int) (err error) {
	o := orm.NewOrm()
	v := PuntoSalarial{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PuntoSalarial{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
