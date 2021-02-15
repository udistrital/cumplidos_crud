package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type PagoMensualAuditoria struct {
	Pago              *PagoMensual
	CargoEjecutor     string
	DocumentoEjecutor string
}

func UpdatePagoMensualAuditoriaById(m *PagoMensualAuditoria) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			err = o.Rollback()
			return
		} else {
			o.Commit()
			return
		}

	}()

	//v := PagoMensual{Id: m.Pago.Id}
	// ascertain id exists in the database
	//if err = o.Read(&v); err == nil {
	//m.Pago.FechaCreacion = v.FechaCreacion
	if _, err = o.Update(m.Pago, "estado_pago_mensual_id", "documento_responsable_id", "cargo_responsable", "fecha_modificacion"); err != nil {
		panic(err)
	} else {
		s := seguimientoAuditoria(m.Pago, m.CargoEjecutor, m.DocumentoEjecutor)
		// Se cambian los estados anteriores activos
		if _, err := o.QueryTable("cambio_estado_pago").Filter("pago_mensual_id", m.Pago.Id).Filter("activo", true).Update(orm.Params{
			"activo":             false,
			"fecha_modificacion": time.Now(),
		}); err != nil {
			panic(err)
		} else {

			if _, err := o.Insert(&s); err != nil {
				panic(err)
			}
		}
	}
	//}
	return
}

func AddPagoMensualAuditoria(m *PagoMensualAuditoria) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			o.Rollback()
			return
		} else {
			o.Commit()
			return
		}

	}()

	if id, err = o.Insert(m.Pago); err != nil {
		logs.Error(err)
		panic(err)
	} else {
		v := seguimientoAuditoria(m.Pago, m.CargoEjecutor, m.DocumentoEjecutor)

		if id, err = o.Insert(&v); err != nil {
			logs.Error(err)
			panic(err)
		}
	}
	return
}

// funcion para agregar un cambio de estado pago de manera invisible para llevar una auditoria
/*func seguimientoAuditoria2(m *PagoMensualAuditoria) (v CambioEstadoPago) {
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
}*/

// funcion para agregar un cambio de estado pago de manera invisible para llevar una auditoria
func seguimientoAuditoria(m *PagoMensual, CargoEjecutor string, DocumentoEjecutor string) (v CambioEstadoPago) {
	v = CambioEstadoPago{
		EstadoPagoMensualId:    m.EstadoPagoMensualId.Id,
		DocumentoResponsableId: DocumentoEjecutor,
		CargoResponsable:       CargoEjecutor,
		PagoMensualId:          m,
		Activo:                 true,
		FechaCreacion:          time.Time{},
		FechaModificacion:      time.Time{},
	}
	return v
}
