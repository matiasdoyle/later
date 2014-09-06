package main

import (
	"database/sql"
	"os"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/matiasdoyle/later/models"
	"github.com/matiasdoyle/later/routes"
	"github.com/matiasdoyle/later/routes/api"
)

func main() {
	m := martini.Classic()
	setupDB()

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", routes.RenderInbox)

	m.Post("/signup", binding.Bind(models.User{}), routes.Signup)

	m.Post("/api/later", binding.Bind(models.Item{}), api.CreateCheckoutItem)
	m.Get("/api/verify-token/:token", api.VerifyToken)

	m.Run()
}

func setupDB() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	models.Init(dbmap)
}
