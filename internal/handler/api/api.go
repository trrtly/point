package api

import (
	"point/internal/core"
	"point/internal/handler/api/assets"
	"point/internal/handler/api/point/page"
	"point/internal/handler/api/point/activity"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Users   core.UserAssetsStore
	Userz   core.UserAssetsService
	Private bool
}

func New(
	users core.UserAssetsStore,
	userz core.UserAssetsService,
) Server {
	return Server{
		Users: users,
		Userz: userz,
	}
}

// Handler defines api handler
func (s Server) Handler() *fiber.App {
	r := fiber.New()
	r.Get("/assets/:uid", assets.HandleFind(s.Users))
	r.Post("/point/page", page.HandlerCreate())
	r.Post("/point/activity", activity.HandlerCreate())
	return r
}
