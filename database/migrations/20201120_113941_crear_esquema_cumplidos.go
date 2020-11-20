package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaCumplidos_20201120_113941 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaCumplidos_20201120_113941{}
	m.Created = "20201120_113941"

	migration.Register("CrearEsquemaCumplidos_20201120_113941", m)
}

// Run the migrations
func (m *CrearEsquemaCumplidos_20201120_113941) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201120_113941_crear_esquema_cumplidos_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *CrearEsquemaCumplidos_20201120_113941) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201120_113941_crear_esquema_cumplidos_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
