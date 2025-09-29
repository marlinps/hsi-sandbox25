package routes

import (
	"authservice-sandbox/api/handlers"

	"authservice-sandbox/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router, userService *user.UserService) {
	api.Post("/login", handlers.LoginHandler(userService))
}
