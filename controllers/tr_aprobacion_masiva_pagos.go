package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_crud/models"
)

// Tr_aprobacion_masiva_pagosController operations for Tr_aprobacion_masiva_pagos
type TrAprobacionMasivaPagosController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrAprobacionMasivaPagosController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Aprobacion masiva de pagos
// @Description create Tr_aprobacion_masiva_pagos
// @Param	body		body 	models.ArregloPagoMensualAuditoria	true		"Tr_aprobacion_masiva_pagos: Aprueba n pagos al tiempo"
// @Success 200 {object} []models.PagoMensual
// @Failure 400 body is empty
// @router / [post]
func (c *TrAprobacionMasivaPagosController) Post() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "TrAprobacionMasivaPagosController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var v models.ArregloPagoMensualAuditoria
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if err = models.AprobarPagos(&v); err == nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful modification", "Data": v}

		} else {
			logs.Error(err)
			c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		fmt.Println(err)
		logs.Error(err)
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}
