package api

import (
	"point/internal/core"
	"point/internal/handler/api/assets"
	"point/internal/handler/api/point/activity"
	"point/internal/handler/api/point/detail"
	"point/internal/handler/api/point/goods"
	"point/internal/handler/api/wechat"
	"point/internal/pkg/hd"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Assets     core.UserAssetsStore
	Activity   core.ActivityStore
	Special    core.ActivitySpecialStore
	Detail     core.UserPointDetailStore
	HD         *hd.HD
	Goods      core.ExchangeGoodsStore
}

func New(
	assets core.UserAssetsStore,
	activity core.ActivityStore,
	special core.ActivitySpecialStore,
	detail core.UserPointDetailStore,
	goods core.ExchangeGoodsStore,
	hashid *hd.HD,
) Server {
	return Server{
		Assets:     assets,
		Activity:   activity,
		Special:    special,
		Detail:     detail,
		HD:         hashid,
		Goods:      goods,
	}
}

// Handler defines api handler
func (s Server) Handler(r fiber.Router) fiber.Router {
	r.Use(logger.New())
	r.Get("/assets", assets.HandleFind(s.Assets))
	r.Post("/point/activity", activity.HandlerCreate(s.Activity, s.Special, s.Detail, s.Assets))
	r.Post("/point/goods", goods.HandlerExchange(s.HD, s.Goods, s.Assets, s.Detail))
	r.Get("/point/goods", goods.HandlerList(s.Goods, s.Assets))
	r.Get("/point/details", detail.HandlerList(s.Detail))
	r.Post("/wechat/bind", wechat.HandleBind(s.Detail))
	return r
}
