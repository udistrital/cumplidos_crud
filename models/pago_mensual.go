package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type PagoMensualAuditoria struct {
	Pago              *PagoMensual
	CargoEjecutor     string
	DocumentoEjecutor string
}

type PagoMensual struct {
	Id                     int                `orm:"column(id);pk;auto"`
	NumeroContrato         string             `orm:"column(numero_contrato)"`
	VigenciaContrato       float64            `orm:"column(vigencia_contrato)"`
	Mes                    float64            `orm:"column(mes)"`
	DocumentoPersonaId     string             `orm:"column(documento_persona_id)"`
	EstadoPagoMensualId    *EstadoPagoMensual `orm:"column(estado_pago_mensual_id);rel(fk)"`
	DocumentoResponsableId string             `orm:"column(documento_responsable_id)"`
	CargoResponsable       string             `orm:"column(cargo_responsable)"`
	Ano                    float64            `orm:"column(ano)"`
	Activo                 bool               `orm:"column(activo)"`
	FechaCreacion          time.Time          `orm:"auto_now;column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion      time.Time          `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone)"`
}

func (t *PagoMensual) TableName() string {
	return "pago_mensual"
}

func init() {
	orm.RegisterModel(new(PagoMensual))
}

// AddPagoMensual insert a new PagoMensual into database and returns
// last inserted Id on success.
func AddPagoMensual(m *PagoMensual) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPagoMensualById retrieves PagoMensual by Id. Returns error if
// Id doesn't exist
func GetPagoMensualById(id int) (v *PagoMensual, err error) {
	o := orm.NewOrm()
	v = &PagoMensual{Id: id}
	if err = o.Read(v); err == nil {
		fmt.Println("modelo: ", v)
		return v, nil
	}
	return nil, err
}

// GetAllPagoMensual retrieves all PagoMensual matches certain condition. Returns empty list if
// no records exist
func GetAllPagoMensual(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PagoMensual)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "in") {
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

	var l []PagoMensual
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

// UpdatePagoMensual updates PagoMensual by Id and returns error if
// Usar update con auditoria par actualizar un pago mensual the record to be updated doesn't exist
func UpdatePagoMensualById(m *PagoMensual) (err error) {
	o := orm.NewOrm()
	v := PagoMensual{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePagoMensual deletes PagoMensual by Id and returns error if
// the record to be deleted doesn't exist
func DeletePagoMensual(id int) (err error) {
	o := orm.NewOrm()
	v := PagoMensual{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PagoMensual{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// funcion para agregar un cambio de estado pago de manera invisible para llevar una auditoria
func seguimientoAuditoria(m *PagoMensualAuditoria) (v CambioEstadoPago) {
	v = CambioEstadoPago{
		EstadoPagoMensualId:    m.Pago.EstadoPagoMensualId.Id,
		DocumentoResponsableId: m.DocumentoEjecutor,
		CargoResponsable:       m.CargoEjecutor,
		PagoMensualId:          m.Pago,
		Activo:                 true,
		FechaCreacion:          time.Time{},
		FechaModificacion:      time.Time{},
	}
	return v
}

func UpdatePagoMensualAuditoriaById(m *PagoMensualAuditoria) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			o.Rollback()
		} else {
			o.Commit()
		}

	}()

	v := PagoMensual{Id: m.Pago.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Update(m.Pago); err != nil {
			panic(err)
		} else {
			s := seguimientoAuditoria(m)
			if _, err := o.Insert(&s); err != nil {
				panic(err)
			}
		}
	}
	return
}

func AddPagoMensualAuditoria(m *PagoMensualAuditoria) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			o.Rollback()
		} else {
			o.Commit()
		}

	}()

	if id, err = o.Insert(m.Pago); err != nil {
		logs.Error(err)
		panic(err)
	} else {
		v := seguimientoAuditoria(m)

		if id, err = o.Insert(&v); err != nil {
			logs.Error(err)
			panic(err)
		}
	}
	return
}
