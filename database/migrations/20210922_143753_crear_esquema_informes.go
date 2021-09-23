package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaInformes_20210922_143753 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaInformes_20210922_143753{}
	m.Created = "20210922_143753"

	migration.Register("CrearEsquemaInformes_20210922_143753", m)
}

// Run the migrations
func (m *CrearEsquemaInformes_20210922_143753) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210922_143753_crear_esquema_informes_up.sql")

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
func (m *CrearEsquemaInformes_20210922_143753) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210922_143753_crear_esquema_informes_down.sql")

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
