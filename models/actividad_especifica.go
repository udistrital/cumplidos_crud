package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ActividadEspecifica struct {
	Id                    int                   `orm:"column(id);pk;auto"`
	ActividadEspecifica   string                `orm:"column(actividad_especifica)"`
	Avance                float64               `orm:"column(avance)"`
	FechaCreacion         time.Time             `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion     time.Time             `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone);null"`
	Activo                bool                  `orm:"column(activo);null"`
	InformeId             *Informe              `orm:"column(informe_id);rel(fk)"`
	ActividadesRealizadas []*ActividadRealizada `orm:"reverse(many)"`
}

func (t *ActividadEspecifica) TableName() string {
	return "actividad_especifica"
}

func init() {
	orm.RegisterModel(new(ActividadEspecifica))
}

// AddActividadEspecifica insert a new ActividadEspecifica into database and returns
// last inserted Id on success.
func AddActividadEspecifica(m *ActividadEspecifica) (id int64, err error) {
	o := orm.NewOrm()
	m.Activo = true
	id, err = o.Insert(m)
	return
}

// GetActividadEspecificaById retrieves ActividadEspecifica by Id. Returns error if
// Id doesn't exist
func GetActividadEspecificaById(id int) (v *ActividadEspecifica, err error) {
	o := orm.NewOrm()
	v = &ActividadEspecifica{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllActividadEspecifica retrieves all ActividadEspecifica matches certain condition. Returns empty list if
// no records exist
func GetAllActividadEspecifica(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ActividadEspecifica))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.HasSuffix(k, "in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				// rewrite dot-notation to Object__Attribute
				v = strings.Replace(v, ".", "__", -1)
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
				// rewrite dot-notation to Object__Attribute
				v = strings.Replace(v, ".", "__", -1)
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

	var l []ActividadEspecifica
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

// UpdateActividadEspecifica updates ActividadEspecifica by Id and returns error if
// the record to be updated doesn't exist
func UpdateActividadEspecificaById(m *ActividadEspecifica) (err error) {
	o := orm.NewOrm()
	v := ActividadEspecifica{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteActividadEspecifica deletes ActividadEspecifica by Id and returns error if
// the record to be deleted doesn't exist
func DeleteActividadEspecifica(id int) (err error) {
	o := orm.NewOrm()
	v := ActividadEspecifica{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ActividadEspecifica{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
