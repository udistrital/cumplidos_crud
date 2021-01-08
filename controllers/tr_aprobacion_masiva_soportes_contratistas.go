package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_crud/models"
)

// operations for TrAprobacionMasivaSoportesContratistasController
type TrAprobacionMasivaSoportesContratistasController struct {
	beego.Controller
}

func (c *TrAprobacionMasivaSoportesContratistasController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// AprobarSoportesContratistas ...
// @Title AprobarSoportesContratistas
// @Description create AprobarSoportesContratistas
// @Param	body		body 	[]models.PagoMensual	true		"Tr_aprobacion_masiva_pagos: Aprueba documentosde n personas al tiempo"
// @Success 200 {object} []models.PagoMensual
// @Failure 400 body is empty
// @router / [post]
func (c *TrAprobacionMasivaSoportesContratistasController) Post() {
	var v []models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if err = models.AprobarSoportesContratistas(&v); err == nil {
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
