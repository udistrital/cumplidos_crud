package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

//funcion para la aprobación masiva de soportes de contratistas
func AprobarSoportesContratistas(m *[]PagoMensual) (err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range *m {
		v.EstadoPagoMensualId.Id = 13
		if _, err = o.Update(&v); err != nil {
			fmt.Println("Pago mensual, soportes de contratistas aprobados", &v)
			err = o.Rollback()
		} else {
			fmt.Println(err)
			//seguimiento auditoria
		}
	}
	err = o.Commit()

	return
}
