package main

import (
	"github.com/go.standard.project.layout/project.web/fiber/crud.postgres.singleton/fmts"
	"github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/controller/route"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	db := dbs.dbStruct{db.Connect()}
	route.AllRoutes(db, app)

	app.Listen(8181)

	fmts.Println("Server Started!")
}
