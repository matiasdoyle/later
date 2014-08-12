package models

import (
	"github.com/coopernurse/gorp"
)

var (
	db *gorp.DbMap
)

func Init(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	dbmap.AddTableWithName(Item{}, "items").SetKeys(true, "Id")

	err := dbmap.CreateTablesIfNotExists()
	if err != nil {
		panic(err)
	}

	db = dbmap
}
