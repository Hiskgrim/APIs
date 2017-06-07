package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SniesNucleoBasico struct {
	Estado      string  `orm:"column(estado)"`
	Descripcion string  `orm:"column(descripcion);null"`
	Nombre      string  `orm:"column(nombre)"`
	Id          int     `orm:"column(id_nucleo);pk"`
	IdArea      float64 `orm:"column(id_area)"`
}

func (t *SniesNucleoBasico) TableName() string {
	return "snies_nucleo_basico"
}

func init() {
	orm.RegisterModel(new(SniesNucleoBasico))
}

// AddSniesNucleoBasico insert a new SniesNucleoBasico into database and returns
// last inserted Id on success.
func AddSniesNucleoBasico(m *SniesNucleoBasico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSniesNucleoBasicoById retrieves SniesNucleoBasico by Id. Returns error if
// Id doesn't exist
func GetSniesNucleoBasicoById(id int) (v *SniesNucleoBasico, err error) {
	o := orm.NewOrm()
	v = &SniesNucleoBasico{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSniesNucleoBasico retrieves all SniesNucleoBasico matches certain condition. Returns empty list if
// no records exist
func GetAllSniesNucleoBasico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SniesNucleoBasico))
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

	var l []SniesNucleoBasico
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

// UpdateSniesNucleoBasico updates SniesNucleoBasico by Id and returns error if
// the record to be updated doesn't exist
func UpdateSniesNucleoBasicoById(m *SniesNucleoBasico) (err error) {
	o := orm.NewOrm()
	v := SniesNucleoBasico{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSniesNucleoBasico deletes SniesNucleoBasico by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSniesNucleoBasico(id int) (err error) {
	o := orm.NewOrm()
	v := SniesNucleoBasico{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SniesNucleoBasico{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
