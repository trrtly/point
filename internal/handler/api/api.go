package api

import (
	"point/internal/core"
	"point/internal/handler/api/assets"
	"point/internal/handler/api/point/activity"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Users    core.UserAssetsStore
	Userz    core.UserAssetsService
	Activity core.ActivityStore
	Special  core.ActivitySpecialStore
	Detail   core.UserPointDetailStore
	Private  bool
}

func New(
	users core.UserAssetsStore,
	userz core.UserAssetsService,
	activity core.ActivityStore,
	special core.ActivitySpecialStore,
	detail core.UserPointDetailStore,
) Server {
	return Server{
		Users:    users,
		Userz:    userz,
		Activity: activity,
		Special:  special,
		Detail:   detail,
	}
}

// Handler defines api handler
func (s Server) Handler() *fiber.App {
	r := fiber.New()
	r.Get("/assets/:uid", assets.HandleFind(s.Users))
	r.Post("/point/activity", activity.HandlerCreate(s.Activity, s.Special, s.Detail))
	return r
}
