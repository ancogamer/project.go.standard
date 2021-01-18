package hcustomer

import (
	fmts "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.singleton/fmts"
	mcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/models/customer"
	mErrors "github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/models/errors"
	rcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.postgresa/repo/customer"
	"github.com/gofiber/fiber"
)

func Post(c *fiber.Ctx) {
	var input mcustomer.Customer
	var Errors mErrors.Errors
	if err := c.BodyParser(&input); err != nil {
		Errors.Msg = err.Error()
		c.Status(503).JSON(Errors)
		return
	}
	input.ImpIp = c.IP()
	if err := rcustomer.Post(input); err != nil {
		Errors.Msg = fmts.Concat("Error: ", err.Error())
		//503 Serviço indisponivel
		c.Status(503).JSON(Errors)
		return
	}
	c.Status(200)
	return
}
