package goods

import (
	"point/internal/core"
	"point/internal/core/status"
	"point/internal/handler/api/render"
	"point/internal/pkg/hd"

	"github.com/pkg/errors"

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
	GoodsNum int32 `json:"goodsNum" validate:"required,number"`
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
func HandlerExchange(
	hashid *hd.HD,
	goods core.ExchangeGoodsStore,
	assets core.UserAssetsStore,
	detail core.UserPointDetailStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Create)

		if err := c.BodyParser(req); err != nil {
			return render.Fail(c, err)
		}
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, err)
		}
		gid, err := hashid.DecodeWithError(req.GoodsYyid)
		if err != nil {
			return render.Fail(c, errors.New("商品 id 有误"))
		}
		egoods, err := goods.Find(uint32(gid[0]))
		if err != nil {
			return render.Fail(c, errors.New("兑换商品不存在"))
		}
		if !egoods.IsActivite() {
			return render.Fail(c, errors.New("商品尚未发布或不在有效期内"))
		}
		uassets, err := assets.Find(req.UID)
		if err != nil {
			return render.Fail(c, err)
		}
		if uassets.MoneyPoint < egoods.MoneyPoint {
			return render.Fail(c, errors.New("消费积分不足"), 5001)
		}
		if uassets.ServicePoint < egoods.ServicePoint {
			return render.Fail(c, errors.New("服务积分不足"), 5002)
		}
		detailm := &core.UserPointDetail{
			UID:      req.UID,
			Type:     core.ActivityTypeUse,
			Status:   status.UserPointDetailArrived,
			GoodsID:  egoods.ID,
			GoodsNum: req.GoodsNum,
			Desc:     egoods.UserDetailDesc,
		}
		detailm.MoneyPoint = egoods.MoneyPoint * float64(req.GoodsNum)
		detailm.ServicePoint = egoods.ServicePoint * float64(req.GoodsNum)

		err = detail.Create(detailm)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{
					"request": req,
					"detailm": detailm,
				},
			).Errorln("/api/point/goods 创建积分详情失败", err)
			return render.Fail(c, errors.New("兑换失败，请稍后重试"), 5003)
		}

		err = assets.DecrPoint(req.UID, detailm.MoneyPoint, detailm.ServicePoint)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{
					"req":     req,
					"egoods":  egoods,
					"uassets": uassets,
				},
			).Errorln("/api/point/goods 用户积分资产扣除失败", err)
			return render.Fail(c, errors.New("兑换失败，请稍后重试"), 5003)
		}

		return render.Success(c, detailm)
	}
}
