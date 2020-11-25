package controllers

import (
	"github.com/udistrital/cumplidos_crud/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
)

// operations for TrAprobacionMasivaDocumentosController
type TrAprobacionMasivaDocumentosController struct {
	beego.Controller
}

func (c *TrAprobacionMasivaDocumentosController) URLMapping() {
	c.Mapping("AprobarDocumentos", c.Post)
}



// AprobarDocumentos ...
// @Title Aprobación masiva de documentos
// @Description create Tr_aprobacion_masiva_documentos
// @Param	body		body 	[]models.PagoMensual	true		"Tr_aprobacion_masiva_documentos: Aprueba documentos al tiempo para n personas"
// @Success 200 {object} []models.PagoMensual
// @Failure 403 body is empty
// @router / [post]
func (c *TrAprobacionMasivaDocumentosController) Post() {

	var v []models.PagoMensual
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {


		if err = models.AprobarDocumentos(&v); err == nil {
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
