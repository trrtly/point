package activity

import (
	"errors"
	"point/internal/core"
	"point/internal/handler/api/render"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Create struct {
	// 用户 id
	UID int64 `json:"uid" validate:"required,number"`
	// 事件编号
	Key string `json:"key" validate:"required"`
	// 特例的类型 1 表示地域 2表示 角色  3表示产品  4表示医院  5表示特定数据
	// 问卷跟页面浏览都传 5
	Type int8 `json:"type" validate:"number"`
	// 特例数值
	Val string `json:"val"`
}

// @Summary 添加事件积分
// @Description 添加事件积分，问卷事件 `key`: `question_answer`；页面浏览事件 `key`: `page_view`
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
	special core.ActivitySpecialStore,
	point core.UserPointDetailStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Create)

		if err := c.BodyParser(req); err != nil {
			return render.Fail(c, err)
		}
		logrus.WithFields(
			logrus.Fields{
				"uid":  req.UID,
				"key":  req.Key,
				"type": req.Type,
				"val":  req.Val,
			},
		).Infoln("receive request")
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, err)
		}
		activity, err := activity.FindEventKey(req.Key)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{
					"uid":  req.UID,
					"key":  req.Key,
					"type": req.Type,
					"val":  req.Val,
				},
			).Errorln("cannot find activity use the event_key")
			return render.Fail(c, err)
		}
		if !activity.IsActivite() {
			return render.Fail(c, errors.New("事件尚未发布或不在有效期内"))
		}
		if activity.HasSpecial() {
			special.FindSVal(activity.ID, req.Val)
		} else {
			if activity.ServicePoint > 0 || activity.MoneyPoint > 0 {
				detail := &core.UserPointDetail{
					UID:          req.UID,
					ActivityID:   activity.ID,
					MoneyPoint:   activity.MoneyPoint,
					ServicePoint: activity.ServicePoint,
					Type:         1,
					Status:       1,
					Desc:         "",
					CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
				}
				point.Create(detail)
			}
		}

		return render.Success(c, "ok")
	}
}
