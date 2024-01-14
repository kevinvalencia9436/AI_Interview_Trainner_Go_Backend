package routes

import (
	handler "up-it-aps-api/app/handlers"
	service "up-it-aps-api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserRoutes(api fiber.Router, store *session.Store) {
	userService := &service.UserService{}
	userHandler := handler.NewUserHandler(userService, store)
	user := api.Group("/users")

	user.Get("/", userHandler.GetUserByEmail)
	user.Get("/settings", userHandler.GetUserSettingsByEmail)
	user.Post("/settings", userHandler.UpdateUserSettings)
}
