package database

import (
	"main/model"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var db *pg.DB

// Init : init db
func Init() {
	// connect to db
	db = pg.Connect(&pg.Options{
		User:     "root",
		Password: "root",
	})

	// create table
	err := db.CreateTable(&model.Todo{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		panic(err)
	}
}

// Get : get db instance
func Get() *pg.DB {
	if db == nil {
		Init()
	}
	return db
}
