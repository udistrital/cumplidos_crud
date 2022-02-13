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

		beego.NSNamespace("/tr_aprobacion_masiva_documentos",
			beego.NSInclude(
				&controllers.TrAprobacionMasivaDocumentosController{},
			),
		),

		beego.NSNamespace("/tr_aprobacion_masiva_pagos",
			beego.NSInclude(
				&controllers.TrAprobacionMasivaPagosController{},
			),
		),

		beego.NSNamespace("/tr_aprobacion_masiva_soportes_contratistas",
			beego.NSInclude(
				&controllers.TrAprobacionMasivaSoportesContratistasController{},
			),
		),

		beego.NSNamespace("/informe",
			beego.NSInclude(
				&controllers.InformeController{},
			),
		),

		beego.NSNamespace("/actividad_especifica",
			beego.NSInclude(
				&controllers.ActividadEspecificaController{},
			),
		),

		beego.NSNamespace("/actividad_realizada",
			beego.NSInclude(
				&controllers.ActividadRealizadaController{},
			),
		),

		beego.NSNamespace("/fechas_carga_cumplidos",
			beego.NSInclude(
				&controllers.FechasCargaCumplidosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
