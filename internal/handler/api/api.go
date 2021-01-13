package api

import (
	"point/internal/core"
	"point/internal/handler/api/assets"
	"point/internal/handler/api/point/activity"
	"point/internal/handler/api/point/detail"
	"point/internal/handler/api/point/goods"
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
	GoodsOrder core.ExchangeGoodsOrderStore
}

func New(
	assets core.UserAssetsStore,
	activity core.ActivityStore,
	special core.ActivitySpecialStore,
	detail core.UserPointDetailStore,
	goods core.ExchangeGoodsStore,
	hashid *hd.HD,
	goodsOrder core.ExchangeGoodsOrderStore,
) Server {
	return Server{
		Assets:     assets,
		Activity:   activity,
		Special:    special,
		Detail:     detail,
		HD:         hashid,
		Goods:      goods,
		GoodsOrder: goodsOrder,
	}
}

// Handler defines api handler
func (s Server) Handler(r fiber.Router) fiber.Router {
	r.Use(logger.New())
	r.Get("/assets", assets.HandleFind(s.Assets))
	r.Post("/point/activity", activity.HandlerCreate(s.Activity, s.Special, s.Detail, s.Assets))
	r.Post("/point/goods", goods.HandlerCreate(s.HD, s.Goods, s.GoodsOrder, s.Assets))
	r.Get("/point/goods", goods.HandlerList(s.Goods, s.GoodsOrder, s.Assets))
	r.Get("/point/details", detail.HandlerList(s.Detail))
	return r
}
