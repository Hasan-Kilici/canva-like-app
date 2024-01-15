package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2"
	"time"
)

var RateLimit = limiter.New(limiter.Config{
	Max:        20,
	Expiration: 1 * time.Minute,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Rate limit exceeded. Please try again later.",
		})
	},
})