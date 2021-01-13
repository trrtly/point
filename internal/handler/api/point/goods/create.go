package goods

import (
	"github.com/pkg/errors"
	"point/internal/core"
	"point/internal/core/status"
	"point/internal/handler/api/render"
	"point/internal/pkg/hd"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Create struct {
	// 用户 id
	UID int64 `json:"uid" validate:"required,number"`
	// 商品编号
	GoodsYyid string `json:"goodsYyid" validate:"required"`
	// 商品数量
	GoodsNum uint32 `json:"goodsNum" validate:"required,number"`
}

// @Summary 积分兑换商品
// @Description 积分兑换商品
// @Tags 兑换商品
// @Version 1.0
// @Accept json
// @Produce json
// @Param body body goods.Create true "请求参数"
// @Success 200 object render.Response "成功返回值"
// @Failure 400 object render.Response "失败返回值"
// @Router /api/point/goods [post]
func HandlerCreate(
	hashid *hd.HD,
	goods core.ExchangeGoodsStore,
	gorders core.ExchangeGoodsOrderStore,
	assets core.UserAssetsStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Create)

		logrus.WithFields(
			logrus.Fields{
				"data": string(c.Body()),
			},
		).Errorln("/api/point/goods request data")

		if err := c.BodyParser(req); err != nil {
			return render.Fail(c, err)
		}
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, err)
		}
		gid, err := hashid.DecodeWithError(req.GoodsYyid)
		if err != nil {
			return render.Fail(c, err)
		}
		egoods, err := goods.Find(uint32(gid[0]))
		if err != nil {
			return render.Fail(c, err)
		}
		if !egoods.IsActivite() {
			return render.Fail(c, errors.New("商品尚未发布或不在有效期内"))
		}
		uassets, err := assets.Find(req.UID)
		if err != nil {
			return render.Fail(c, err)
		}
		if uassets.MoneyPoint < egoods.MoneyPoint {
			return render.Fail(c, errors.New("消费积分不足"))
		}
		if uassets.ServicePoint < egoods.ServicePoint {
			return render.Fail(c, errors.New("服务积分不足"))
		}
		gorder := &core.ExchangeGoodsOrder{
			UID:      req.UID,
			GoodsID:  egoods.ID,
			GoodsNum: req.GoodsNum,
			Status:   status.Regular,
		}
		gorder.MoneyPoint = egoods.MoneyPoint * float64(req.GoodsNum)
		gorder.ServicePoint = egoods.ServicePoint * float64(req.GoodsNum)

		err = assets.DecrPoint(req.UID, gorder.MoneyPoint, gorder.ServicePoint)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{
					"req":     req,
					"egoods":  egoods,
					"uassets": uassets,
				},
			).Errorln("/api/point/goods decr fail")
			return render.Fail(c, errors.New("兑换失败，请稍后重试"))
		}

		err = gorders.Create(gorder)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{
					"request":        req,
					"exchange_goods": egoods,
					"user_assets":    uassets,
					"goods_order":    gorder,
				},
			).Errorln("/api/point/goods create goods order fail")
		}

		return render.Success(c, "ok")
	}
}
