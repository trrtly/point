package page

import (
	"point/internal/handler/api/render"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Create struct {
	// 用户id
	UID int64 `validate:"required,number"`
	// 页面路径的 `path` 部分，例如：`http://api.youyao.com/user/point`，则 `uri` 为 `user/point`
	URI string `validate:"required"`
}

// @Summary 添加页面浏览积分
// @Description 添加页面浏览积分
// @Tags 页面浏览积分
// @Version 1.0
// @Accept json
// @Produce json
// @Param body body page.Create true "请求参数"
// @Success 200 object render.Response "成功返回值"
// @Failure 400 object render.Response "失败返回值"
// @Router /api/point/page [post]
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
