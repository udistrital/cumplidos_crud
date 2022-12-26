package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Informe struct {
	Id                     int                    `orm:"column(id);pk;auto"`
	Activo                 bool                   `orm:"column(activo);null"`
	FechaCreacion          time.Time              `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion      time.Time              `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone);null"`
	PeriodoInformeInicio   time.Time              `orm:"column(periodo_informe_inicio);type(timestamp without time zone);null"`
	PeriodoInformeFin      time.Time              `orm:"column(periodo_informe_fin);type(timestamp without time zone);null"`
	Proceso                string                 `orm:"column(proceso)"`
	PagoMensualId          *PagoMensual           `orm:"column(pago_mensual_id);rel(fk)"`
	ActividadesEspecificas []*ActividadEspecifica `orm:"reverse(many)"`
}

func (t *Informe) TableName() string {
	return "informe"
}

func init() {
	orm.RegisterModel(new(Informe))
}

// AddInforme insert a new Informe into database and returns
// last inserted Id on success.
func AddInforme(m *Informe) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//AddInformeCompleto insert a new Informe with actividades_especificas and actividades_realizadas
func AddInformeCompleto(m *Informe) (id int64, err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		o.Rollback()
		err = errors.New("Error al guardar el informe, se revierte la transaccion")
		return id, err
	}
	if id, err := o.Insert(m); err != nil {
		o.Rollback()
		err = errors.New("Error al guardar el informe, se revierte la transaccion")
		return id, err
	} else {
		for _, actEsp := range m.ActividadesEspecificas {
			var inf Informe
			inf.Id = int(id)
			actEsp.InformeId = &inf
			actEsp.Activo = true
			if id_actesp, err := o.Insert(actEsp); err != nil {
				o.Rollback()
				err = errors.New("Error al guardar el informe, se revierte la transaccion")
				return id, err
			} else {
				for _, actRea := range actEsp.ActividadesRealizadas {
					var act_esp ActividadEspecifica
					act_esp.Id = int(id_actesp)
					actRea.ActividadEspecificaId = &act_esp
					actRea.Activo = true
					if _, err := o.Insert(actRea); err != nil {
						o.Rollback()
						err = errors.New("Error al guardar el informe, se revierte la transaccion")
						return id, err
					}
				}
			}
		}
	}
	err = o.Commit()
	return id, err
}

// GetInformeById retrieves Informe by Id. Returns error if
// Id doesn't exist
func GetInformeById(id int) (v *Informe, err error) {
	o := orm.NewOrm()
	v = &Informe{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInforme retrieves all Informe matches certain condition. Returns empty list if
// no records exist
func GetAllInforme(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Informe))
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

	var l []Informe
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

// UpdateInforme updates Informe by Id and returns error if
// the record to be updated doesn't exist
func UpdateInformeById(m *Informe) (err error) {
	o := orm.NewOrm()
	v := Informe{Id: m.Id}
	v.Activo = true
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateInforme updates Informe by Id with actividades_especificas and actividades_realizadas and returns error if
// the record to be updated doesn't exist
func UpdateInformeCompletoById(m *Informe) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		o.Rollback()
		err = errors.New("Error al guardar el informe, se revierte la transaccion")
		return err
	}
	v := Informe{Id: m.Id}
	v.Activo = true
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
			for _, actEsp := range m.ActividadesEspecificas {
				var inf Informe
				inf.Id = int(m.Id)
				actEsp.InformeId = &inf
				ae := &ActividadEspecifica{Id: actEsp.Id}
				if err = o.Read(ae); err == nil { //la actividad existe y sera actualizada
					if _, err := o.Update(actEsp); err != nil {
						o.Rollback()
						err = errors.New("Error al guardar el informe, se revierte la transaccion")
						return err
					} else { //procede a insertar o actualizar las actividades realizadas
						for _, actRea := range actEsp.ActividadesRealizadas {
							var act_esp ActividadEspecifica
							act_esp.Id = actEsp.Id
							actRea.ActividadEspecificaId = &act_esp
							ar := &ActividadRealizada{Id: actRea.Id}
							if err = o.Read(ar); err == nil { //Se valida si la actividad realizada existe
								if _, err := o.Update(actRea); err != nil {
									o.Rollback()
									err = errors.New("Error al guardar el informe, se revierte la transaccion")
									return err
								}
							} else { //la actividad realizada no existe y tiene que ser creada
								actRea.Activo = true
								if _, err := o.Insert(actRea); err != nil {
									o.Rollback()
									err = errors.New("Error al guardar el informe, se revierte la transaccion")
									return err
								}
							}
						}
					}
				} else { //la actividad especifica no existe y debe ser creada
					actEsp.Activo = true
					if id_actesp, err := o.Insert(actEsp); err != nil {
						o.Rollback()
						err = errors.New("Error al guardar el informe, se revierte la transaccion")
						return err
					} else {
						for _, actRea := range actEsp.ActividadesRealizadas {
							var act_esp ActividadEspecifica
							act_esp.Id = int(id_actesp)
							actRea.ActividadEspecificaId = &act_esp
							actRea.Activo = true
							if _, err := o.Insert(actRea); err != nil {
								o.Rollback()
								err = errors.New("Error al guardar el informe, se revierte la transaccion")
								return err
							}
						}
					}

				}
			}
		} else {
			o.Rollback()
			return err
		}
	} else {
		o.Rollback()
		return err
	}
	err = o.Commit()
	return
}

// DeleteInforme deletes Informe by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInforme(id int) (err error) {
	o := orm.NewOrm()
	v := Informe{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Informe{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
