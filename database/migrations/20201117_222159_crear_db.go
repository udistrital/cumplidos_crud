package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearDb_20201117_222159 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearDb_20201117_222159{}
	m.Created = "20201117_222159"

	migration.Register("CrearDb_20201117_222159", m)
}

// Run the migrations
func (m *CrearDb_20201117_222159) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CrearDb_20201117_222159) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
