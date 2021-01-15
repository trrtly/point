package goods

import (
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type List struct {
	UID      int64 `query:"uid,required,number"`
	Page     int   `query:"page,number"`
	PageSize int   `query:"page_size,number"`
}

type response struct {
	render.Response
	Data struct {
		// 商品列表
		List []*core.ExchangeGoods `json:"list"`
		// 页码值
		Page int `json:"page" example:"1"`
		// 每页显示条数
		PageSize int `json:"page_size" example:"20"`
		// 总条数
		Total int64 `json:"total" example:"100"`
	} `json:"data"`
}

// @Summary 积分兑换商品列表
// @Description 积分兑换商品列表
// @Tags 兑换商品
// @Version 1.0
// @Accept json
// @Produce json
// @Param uid query int true "uid"
// @Param page query int false "当前页码"
// @Param page_size query int false "每页显示条数"
// @Success 200 object response "成功返回值"
// @Failure 400 object render.Response "失败返回值"
// @Router /api/point/goods [get]
func HandlerList(
	goods core.ExchangeGoodsStore,
	assets core.UserAssetsStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(List)
		if err := c.QueryParser(req); err != nil {
			return render.Fail(c, err)
		}
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, err)
		}
		goods, total, err := goods.ListActivate(req.Page, req.PageSize)
		if err != nil {
			return render.Fail(c, err)
		}

		res := map[string]interface{}{
			"list":      goods,
			"page":      req.Page,
			"page_size": req.PageSize,
			"total":     total,
		}

		return render.Success(c, res)
	}
}
