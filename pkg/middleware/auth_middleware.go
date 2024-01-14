package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// This or another middleware function needs to be checking the validity of the jwt
func WithAuthenticatedUserApi(c *fiber.Ctx) error {
	log.Println("WithAuthenticatedUserApi")
	apiKey := c.Get("x-api-key")
	log.Println("Api Key is " + apiKey)
	if apiKey != os.Getenv("API_KEY") {
		log.Println("Wrong API key")
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	log.Println("Correct API key")
	return c.Next()
}
