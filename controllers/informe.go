package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/udistrital/cumplidos_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// InformeController operations for Informe
type InformeController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Informe
// @Param	body		body 	models.Informe	true		"body for Informe content"
// @Success 201 {int} models.Informe
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *InformeController) Post() {
	var test map[string]interface{}
	var v models.Informe
	json.Unmarshal(c.Ctx.Input.RequestBody, &test)
	fmt.Println(test)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("informe completo", v)
		fmt.Println(v.ActividadesEspecificas)
		if len(v.ActividadesEspecificas) == 0 {
			if _, err := models.AddInforme(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registration successful", "Data": v}
			} else {
				logs.Error(err)
				c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
				c.Abort("400")
			}
		} else {
			fmt.Println("informe completo")
			if _, err := models.AddInformeCompleto(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registration successful", "Data": v}
			} else {
				logs.Error(err)
				c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
				c.Abort("400")
			}
		}
	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Informe by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Informe
// @Failure 404 not found resource
// @router /:id [get]
func (c *InformeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetInformeById(id)
	if err != nil {
		logs.Error(err)
		c.Data["mesaage"] = "Error service GetOne: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Informe
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Informe
// @Failure 404 not found resource
// @router / [get]
func (c *InformeController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllInforme(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["mesaage"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Informe
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Informe	true		"body for Informe content"
// @Success 200 {object} models.Informe
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *InformeController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Informe{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if len(v.ActividadesEspecificas) == 0 {
			if err := models.UpdateInformeById(&v); err == nil {
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update successful", "Data": v}
			} else {
				logs.Error(err)
				c.Data["mesaage"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
				c.Abort("400")
			}
		} else {
			fmt.Println("update Informe completo")
			if err := models.UpdateInformeCompletoById(&v); err == nil {
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update successful", "Data": v}
			} else {
				logs.Error(err)
				c.Data["mesaage"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
				c.Abort("400")
			}
		}

	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Informe
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *InformeController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteInforme(id); err == nil {
		d := map[string]interface{}{"Id": id}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Delete successful", "Data": d}
	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Delete: Request contains incorrect parameter"
		c.Abort("404")
	}
	c.ServeJSON()
}
