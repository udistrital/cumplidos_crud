// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/cumplidos_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/item_informe_tipo_contrato",
			beego.NSInclude(
				&controllers.ItemInformeTipoContratoController{},
			),
		),

		beego.NSNamespace("/item_informe",
			beego.NSInclude(
				&controllers.ItemInformeController{},
			),
		),

		beego.NSNamespace("/pago_mensual",
			beego.NSInclude(
				&controllers.PagoMensualController{},
			),
		),

		beego.NSNamespace("/cambio_estado_pago",
			beego.NSInclude(
				&controllers.CambioEstadoPagoController{},
			),
		),

		beego.NSNamespace("/estado_pago_mensual",
			beego.NSInclude(
				&controllers.EstadoPagoMensualController{},
			),
		),

		beego.NSNamespace("/soporte_pago_mensual",
			beego.NSInclude(
				&controllers.SoportePagoMensualController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
