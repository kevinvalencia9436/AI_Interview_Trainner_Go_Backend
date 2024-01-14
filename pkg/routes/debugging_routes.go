package routes

import (
	handler "up-it-aps-api/app/handlers"
	service "up-it-aps-api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func DebuggingRoutes(api fiber.Router, store *session.Store) {
	userService := &service.UserService{}
	debuggingHandler := handler.NewDebuggingHandler()
	userHandler := handler.NewUserHandler(userService, store)
	debugging := api.Group("/debugging")

	debugging.Post("/", debuggingHandler.Debugging)
	debugging.Post("/get-user-details", debuggingHandler.GetUserDetails)
	debugging.Get("/get-all-users", userHandler.GetAllUsers)
	debugging.Post("/update-tokens-for-user", debuggingHandler.UpdateUserTokens)
}
