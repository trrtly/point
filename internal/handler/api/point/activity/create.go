package activity

import (
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/pkg/errors"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Create struct {
	// 用户 id
	UID int64 `json:"uid,omitempty" validate:"number"`
	// 微信 wechatUserId，针对未登录用户，wechat_user 表 id
	WechatUserID int64 `json:"wechatUserId"`
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
// @Description 登录用户传 `uid` ，未登录用户传 `wechatUserId`
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
	assets core.UserAssetsStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Create)

		if err := c.BodyParser(req); err != nil {
			return render.Fail(c, errors.Wrap(err, "参数解析失败"))
		}

		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, errors.Wrap(err, "参数验证失败"))
		}
		if req.UID <= 0 && req.WechatUserID <= 0 {
			return render.Fail(c, errors.New("uid与openid不能同时为空"))
		}
		activity, err := activity.FindEventKey(req.Key)
		if err != nil {
			return render.Fail(c, errors.Wrap(err, "活动不存在"))
		}
		if !activity.IsActivite() {
			return render.Fail(c, errors.New("事件尚未发布或不在有效期内"))
		}
		detail := &core.UserPointDetail{
			UID:          req.UID,
			WechatUserID: req.WechatUserID,
			ActivityID:   activity.ID,
			Type:         core.ActivityTypeGain,
			Status:       core.UserPointDetailArrived,
		}
		if activity.HasSpecial() {
			aspecial, err := special.FindSVal(activity.ID, req.Val)
			if err == nil && aspecial.IsActivite() && aspecial.IsPointGtZero() {
				detail.MoneyPoint = aspecial.MoneyPoint
				detail.ServicePoint = aspecial.ServicePoint
				detail.Desc = aspecial.PointDesc
				detail.ActivitySpecialID = aspecial.ID
			}
		}
		if !detail.IsPointGtZero() && activity.IsPointGtZero() {
			detail.MoneyPoint = activity.MoneyPoint
			detail.ServicePoint = activity.ServicePoint
			detail.Desc = activity.PointDesc
		}
		if !detail.IsPointGtZero() {
			return render.Success(c, detail)
		}
		if detail.UID > 0 {
			uassets, err := assets.FindOrCreate(req.UID)
			if err != nil {
				return render.Fail(c, errors.Wrap(err, "用户资产异常"))
			}
			err = assets.IncrPoint(detail.UID, uassets, &core.UserAssets{
				MoneyPoint:   detail.MoneyPoint,
				ServicePoint: detail.ServicePoint,
			})
			if err != nil {
				return render.Fail(c, errors.Wrap(err, "用户积分更新失败"))
			}
		}

		point.Create(detail)

		return render.Success(c, detail)
	}
}
