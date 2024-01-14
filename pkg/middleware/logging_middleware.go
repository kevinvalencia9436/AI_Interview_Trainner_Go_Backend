package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	log.Printf("%s request to %s", c.Method(), c.Path())
	return c.Next()
}
