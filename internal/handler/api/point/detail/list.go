package detail

import (
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type List struct {
	UID      int64 `query:"uid,required,number"`
	PType    int8  `query:"ptype,number"`
	Page     int   `query:"page,number"`
	PageSize int   `query:"pageSize,number"`
}

type response struct {
}

// @Summary 获取积分明细列表
// @Description 获取积分明细列表
// @Tags 积分明细列表
// @Version 1.0
// @Accept json
// @Produce json
// @Param uid query int true "uid"
// @Param ptype query int true "积分类型 1：消费积分，2：服务积分"
// @Param page query int false "当前页码"
// @Param page_size query int false "每页显示条数"
// @Success 200 object render.Response "成功返回值"
// @Failure 400 object render.Response "失败返回值"
// @Router /api/points [get]
func HandlerList(
	detail core.UserPointDetailStore,
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
		details, total, err := detail.List(req.UID, req.PType, req.Page, req.PageSize)
		if err != nil {
			return render.Fail(c, err)
		}

		res := map[string]interface{}{
			"list":      details,
			"page":      req.Page,
			"page_size": req.PageSize,
			"total":     total,
		}

		return render.Success(c, res)
	}
}
