package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/PingMe-backend/services/user/pkg/database"
	"github.com/radekkrejcirik01/PingMe-backend/services/user/pkg/model/people"
)

// CreatePeopleInvitation POST /create/people/invitation
func CreatePeopleInvitation(c *fiber.Ctx) error {
	t := &people.PeopleInvitationTable{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	message, err := people.CreatePeopleInvitation(database.DB, t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(PeopleResponse{
		Status:  "succes",
		Message: message,
	})
}

// GetPeople POST /get/people
func GetPeople(c *fiber.Ctx) error {
	t := &people.People{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	people, err := people.GetPeople(database.DB, t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(PeopleResponse{
		Status:  "succes",
		Message: "People succesfully got!",
		Data:    people,
	})
}

// AcceptPeopleInvitation POST /accept/people/invitation
func AcceptPeopleInvitation(c *fiber.Ctx) error {
	t := &people.AcceptInvite{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := people.AcceptInvitation(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Invitation succesfully accepted!",
	})
}

// CheckInvitations POST /check/people/invitations
func CheckInvitations(c *fiber.Ctx) error {
	t := &people.CheckIfFriend{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	record, err := people.CheckInvitations(database.DB, t)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(Response{
			Status:  "succes",
			Message: "No record found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(CheckInvitationsResponse{
		Status:  "succes",
		Message: "Invitations succesfully checked",
		Data:    record,
	})
}

// RemoveFriend POST /remove/friend
func RemoveFriend(c *fiber.Ctx) error {
	t := &people.Remove{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := people.RemoveFriend(database.DB, t); err != nil {
		return c.Status(fiber.StatusOK).JSON(Response{
			Status:  "succes",
			Message: "No record found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Friend removed",
	})
}
