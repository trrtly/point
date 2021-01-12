package assets

import (
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/gofiber/fiber/v2"
)

type Find struct {
	UID int64 `query:"uid,required,number"`
}

// HandleFind returns an fiber.Handler that writes a json-encoded
// user assets details to the response body.
func HandleFind(
	assets core.UserAssetsStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Find)
		if err := c.QueryParser(req); err != nil {
			return render.Fail(c, err)
		}
		assets, err := assets.Find(req.UID)
		if err != nil {
			return render.Fail(c, err)
		}
		return render.Success(c, assets)
	}
}
