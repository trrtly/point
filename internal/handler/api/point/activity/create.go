package activity

import (
	"errors"
	"point/internal/handler/api/render"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Create struct {
	// 用户 id
	UID int64 `json:"uid" validate:"required,number"`
	// 事件编号
	ActivityCode string `json:"activityCode" validate:"required"`
	// 特例编号
	SpecialCode string `json:"specialCode"`
	// 特例数值
	SpecialVal string `json:"specialVal"`
}

// @Summary 添加事件积分
// @Description 添加事件积分
// @Tags 事件积分
// @Version 1.0
// @Accept json
// @Produce json
// @Param body body activity.Create true "请求参数"
// @Success 200 object render.Response "成功返回值"
// @Failure 400 object render.Response "失败返回值"
// @Router /api/point/activity [post]
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
		if req.ActivityCode != "question_answer" {
			return render.Fail(c, errors.New("不支持的事件编号"))
		}

		return render.Success(c, "ok")
	}
}