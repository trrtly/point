package assets

import (
	"point/internal/core"
	"point/internal/handler/api/render"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// HandleFind returns an fiber.Handler that writes a json-encoded
// user assets details to the response body.
func HandleFind(
	assets core.UserAssetsStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		uid, err := strconv.Atoi(c.Params("uid"))
		if err != nil {
			return render.Fail(c, err)
		}
		assets, err := assets.Find(int64(uid))
		if err != nil {
			return render.Fail(c, err)
		}
		return render.Success(c, assets)
	}
}
