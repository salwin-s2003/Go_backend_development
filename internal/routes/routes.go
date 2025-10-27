package routes

import (
	"github.com/gofiber/fiber/v2"
	"task/internal/handler"
)

func RegisterUserRoutes(app *fiber.App, h *handler.UserHandler) {
	api := app.Group("/users")
	api.Post("/", h.CreateUser)
	api.Get("/", h.ListUsers)
	api.Get("/:id", h.GetUserByID)
	api.Put("/:id", h.UpdateUser)
	api.Delete("/:id", h.DeleteUser)
}
