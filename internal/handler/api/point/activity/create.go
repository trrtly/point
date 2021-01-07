package activity

import (
	"errors"
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
func HandlerCreate(
	activity core.ActivityStore,
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
		activity, err := activity.FindEventKey(req.ActivityCode)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{
					"uid":          req.UID,
					"activityCode": req.ActivityCode,
					"specialCode":  req.SpecialCode,
					"specialVal":   req.SpecialVal,
				},
			).Errorln("cannot find activity use the event_key")
			return render.Fail(c, err)
		}
		if !activity.IsActivite() {
			return render.Fail(c, errors.New("活动尚未发布或不在有效期内"))
		}

		return render.Success(c, activity)
	}
}
