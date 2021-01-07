package goods

import (
	"point/internal/handler/api/render"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Create struct {
	// 用户 id
	UID int64 `json:"uid" validate:"required,number"`
	// 商品编号
	GoodsYyid string `json:"goodsYyid" validate:"required"`
	// 商品数量
	GoodsNum int64 `json:"goodsNum" validate:"required,number"`
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
func HandlerCreate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Create)

		if err := c.BodyParser(req); err != nil {
			return render.Fail(c, err)
		}
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, err)
		}

		return render.Success(c, "ok")
	}
}
