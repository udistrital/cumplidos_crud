package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//funcion para la aprobación masiva de pagos
func AprobarPagos(m *[]PagoMensual) (err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range *m {
		v.EstadoPagoMensualId.Id = 5
		if _, err = o.Update(&v); err != nil {
			fmt.Println("Pago mensual, pago aprobado: ", &v)
			err = o.Rollback()
		} else {
			fmt.Println(err)
		}
	}
	err = o.Commit()

	return
}
