package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "site-backend/controller"
)

func HandleRequests(api *fiber.App) {
	api.Post("/api/form", controller.CreateForm)
	api.Get("/api/form/:user", controller.GetForm)
}
