package main

import (
	"database/sql"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/matiasdoyle/checkout/models"
	"github.com/matiasdoyle/checkout/routes"
	"github.com/matiasdoyle/checkout/routes/api"
)

func main() {
	m := martini.Classic()
	setupDB()

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", routes.RenderInbox)

	m.Post("/signup", binding.Bind(models.User{}), routes.Signup)

	m.Post("/api/checkout", binding.Bind(models.Item{}), api.CreateCheckoutItem)

	m.Run()
}

func setupDB() {
	db, err := sql.Open("mysql", "root@/checkout?parseTime=true")
	if err != nil {
		panic(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	models.Init(dbmap)
}
