package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/PingMe-backend/services/user/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	app.Post("/create", controller.CreateUser)
	app.Post("/get", controller.GetUser)

	app.Post("/create/people/invitation", controller.CreatePeopleInvitation)
	app.Post("/get/people", controller.GetPeople)

	app.Post("/create/hangout", controller.CreateHangout)
	app.Post("/get/hangouts", controller.GetHangouts)

	return app
}
