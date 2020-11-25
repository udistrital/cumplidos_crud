package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type LlenarTablasParametricas_20201120_151503 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &LlenarTablasParametricas_20201120_151503{}
	m.Created = "20201120_151503"

	migration.Register("LlenarTablasParametricas_20201120_151503", m)
}

// Run the migrations
func (m *LlenarTablasParametricas_20201120_151503) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20201120_151503_llenar_tablas_parametricas_up.sql")

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
func (m *LlenarTablasParametricas_20201120_151503) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20201120_151503_llenar_tablas_parametricas_down.sql")

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
