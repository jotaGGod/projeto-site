package controller

import (
	"github.com/gofiber/fiber/v2"
	"site-backend/database/entities"
	"site-backend/service"
	"site-backend/utils"
)

func CreateForm(c *fiber.Ctx) error {
	var form entities.FormEntity
	if err := c.BodyParser(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := utils.Validate.Struct(form)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	createdForm, err := service.CreateForm(form)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(createdForm)
}

func GetForm(c *fiber.Ctx) error {

	return nil
}
