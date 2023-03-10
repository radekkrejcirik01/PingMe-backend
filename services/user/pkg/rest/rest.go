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

	app.Post("/upload/photo", controller.UploadPhoto)

	app.Post("/create/people/invitation", controller.CreatePeopleInvitation)
	app.Post("/get/people", controller.GetPeople)
	app.Post("/accept/people/invitation", controller.AcceptPeopleInvitation)
	app.Post("/check/people/invitations", controller.CheckInvitations)
	app.Post("/remove/friend", controller.RemoveFriend)

	app.Post("/create/hangout/group", controller.CreateGroupHangout)
	app.Post("/create/hangout", controller.CreateHangout)
	app.Post("/get/hangouts", controller.GetHangouts)
	app.Post("/get/hangout", controller.GetHangout)
	app.Post("/update/hangout", controller.UpdateHangout)
	app.Post("/remove/hangout/user", controller.RemoveUserFromHangout)
	app.Post("/delete/hangout", controller.DeleteHangout)
	app.Post("/accept/hangout/invitation", controller.AcceptHangoutInvitation)
	app.Post("/get/hangout/usernames", controller.GetHangoutUsernames)
	app.Post("/send/hangout/invitation", controller.SendHangoutInvitation)

	app.Post("/get/notifications", controller.GetNotifications)

	return app
}
