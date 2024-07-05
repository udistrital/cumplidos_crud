package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	_ "github.com/udistrital/cumplidos_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerrorv2"
)

func main() {
	AllowedOrigins := []string{"*.udistrital.edu.co"}
	orm.Debug = true
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		url.QueryEscape(beego.AppConfig.String("PGpass"))+"@"+
		beego.AppConfig.String("PGhost")+":"+
		beego.AppConfig.String("PGport")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschema")+"")
	if beego.BConfig.RunMode == "dev" {
		AllowedOrigins = []string{"*"}
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: AllowedOrigins,
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	apistatus.Init()
	beego.Run()
}
