package detail

import (
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	render.Response
	Data respData `json:"data"`
}

type respData struct {
	// 积分列表
	List []*core.UserPointDetail `json:"list"`
	// 页码值
	Page int `json:"page" example:"1"`
	// 每页显示条数
	PageSize int `json:"page_size" example:"20"`
	// 总条数
	Total int64 `json:"total" example:"100"`
}

// @Summary 获取积分明细列表
// @Description 获取积分明细列表
// @Tags 积分明细列表
// @Version 1.0
// @Accept json
// @Produce json
// @Param uid query int true "uid"
// @Param fetchService query bool true "是否获取服务积分，true：返回服务积分，false：返回消费积分"
// @Param page query int false "当前页码"
// @Param pageSize query int false "每页显示条数"
// @Param type query int false "类型 1：发放，2：使用， 0或不传为全部"
// @Success 200 object response "成功返回值"
// @Failure 400 object render.Response "失败返回值"
// @Router /api/point/details [get]
func HandlerList(
	detail core.UserPointDetailStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(core.UserPointDetailListRequest)

		if err := c.QueryParser(req); err != nil {
			return render.Fail(c, err)
		}
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, err)
		}
		details, total, err := detail.List(req)
		if err != nil {
			return render.Fail(c, err)
		}

		data := new(respData)
		data.List = details
		data.Page = req.Page
		data.PageSize = req.PageSize
		data.Total = total

		return render.Success(c, data)
	}
}
