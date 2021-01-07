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
	Private  bool
}

func New(
	users core.UserAssetsStore,
	userz core.UserAssetsService,
	activity core.ActivityStore,
) Server {
	return Server{
		Users:    users,
		Userz:    userz,
		Activity: activity,
	}
}

// Handler defines api handler
func (s Server) Handler() *fiber.App {
	r := fiber.New()
	r.Get("/assets/:uid", assets.HandleFind(s.Users))
	r.Post("/point/activity", activity.HandlerCreate(s.Activity))
	return r
}
