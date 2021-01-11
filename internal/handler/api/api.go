package api

import (
	"point/internal/core"
	"point/internal/handler/api/assets"
	"point/internal/handler/api/point/activity"
	"point/internal/handler/api/point/goods"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Assets   core.UserAssetsStore
	Activity core.ActivityStore
	Special  core.ActivitySpecialStore
	Detail   core.UserPointDetailStore
	Private  bool
}

func New(
	assets core.UserAssetsStore,
	activity core.ActivityStore,
	special core.ActivitySpecialStore,
	detail core.UserPointDetailStore,
) Server {
	return Server{
		Assets:   assets,
		Activity: activity,
		Special:  special,
		Detail:   detail,
	}
}

// Handler defines api handler
func (s Server) Handler(r fiber.Router) fiber.Router {
	r.Use(logger.New())
	r.Get("/assets/:uid", assets.HandleFind(s.Assets))
	r.Post("/point/activity", activity.HandlerCreate(s.Activity, s.Special, s.Detail, s.Assets))
	r.Post("/point/goods", goods.HandlerCreate())
	return r
}
