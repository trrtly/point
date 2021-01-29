package health

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Handler returns a new health handler.
func Handler(r fiber.Router) fiber.Router {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
	return r
}
