package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:CambioEstadoPagoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:EstadoPagoMensualController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ItemInformeTipoContratoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:PagoMensualController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:SoportePagoMensualController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadEspecificaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:ActividadRealizadaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:TrAprobacionMasivaDocumentosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:TrAprobacionMasivaDocumentosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:TrAprobacionMasivaPagosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:TrAprobacionMasivaPagosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:TrAprobacionMasivaSoportesContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_crud/controllers:TrAprobacionMasivaSoportesContratistasController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
