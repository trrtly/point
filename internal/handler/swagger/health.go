package swagger

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	_ "point/docs"
)

// Handler returns a new swagger handler.
func Handler() *fiber.App {
	r := fiber.New()
	r.Get("/*", swagger.Handler)
	return r
}
