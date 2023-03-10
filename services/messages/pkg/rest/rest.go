package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/PingMe-backend/services/messages/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	app.Post("/create/conversation", controller.CreateConversation)
	app.Post("/get/conversations/:page", controller.GetConversations)
	app.Post("/get/conversation/details", controller.GetConversationDetails)
	app.Post("/get/conversation/usernames", controller.GetConversationUsernames)
	app.Post("/update/conversation", controller.UpdateConversation)
	app.Post("/add/conversation/users", controller.AddConversationUsers)
	app.Post("/remove/conversation", controller.RemoveConversation)
	app.Post("/remove/conversation/user", controller.RemoveUserFromConversation)
	app.Post("/delete/conversation", controller.DeleteConversation)

	app.Post("/get/messages", controller.GetMessages)

	app.Post("/update/read", controller.UpdatLastRead)
	app.Post("/react/message", controller.MessageReacted)
	app.Post("/send/message", controller.SendMessage)
	app.Post("/send/typing", controller.SendTyping)

	return app
}
