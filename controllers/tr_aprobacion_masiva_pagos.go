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
// @Param	body		body 	[]models.PagoMensual	true		"Tr_aprobacion_masiva_pagos: Aprueba n pagos al tiempo"
// @Success 200 {object} []models.PagoMensual
// @Failure 400 body is empty
// @router / [post]
func (c *TrAprobacionMasivaPagosController) Post() {
	var v []models.PagoMensual
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
