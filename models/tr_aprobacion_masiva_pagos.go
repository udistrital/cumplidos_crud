package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ArregloPagoMensualAuditoria struct {
	Pagos             *[]PagoMensual
	CargoEjecutor     string
	DocumentoEjecutor string
}

//funcion para la aprobación masiva de pagos
func AprobarPagos(m *ArregloPagoMensualAuditoria) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			o.Rollback()
			outputError := map[string]interface{}{"funcion": "/AprobarPagos", "err": err, "status": "500"}
			panic(outputError)

		} else {
			err = o.Commit()
			return
		}

	}()

	for _, v := range *m.Pagos {

		v.EstadoPagoMensualId.Id = 5
		if _, err = o.Update(&v, "estado_pago_mensual_id", "fecha_modificacion"); err != nil {
			panic(err)
		} else {
			s := seguimientoAuditoria(&v, m.CargoEjecutor, m.DocumentoEjecutor)
			/*if _, err := o.Insert(&s); err != nil {
				panic(err)
			}*/
			// Se cambian los estados anteriores activos
			if _, err := o.QueryTable("cambio_estado_pago").Filter("pago_mensual_id", v.Id).Filter("activo", true).Update(orm.Params{
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
	}

	return
}
