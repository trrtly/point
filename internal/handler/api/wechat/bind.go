package wechat

import (
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Bind struct {
	// 用户 id
	UID int64 `json:"uid" validate:"required,number"`
	// wechat_user 表 id
	WechatUserID int64 `json:"wechatUserId" validate:"required,number"`
}

// @Summary 绑定微信关联关系
// @Description 绑定微信关联关系
// @Tags 事件积分
// @Version 1.0
// @Accept json
// @Produce json
// @Param body body wechat.Bind true "请求参数"
// @Success 200 object render.Response "成功返回值"
// @Failure 400 object render.Response "失败返回值"
// @Router /api/wechat/bind [post]
func HandleBind(
	detail core.UserPointDetailStore,
	assets core.UserAssetsStore,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(Bind)
		if err := c.BodyParser(req); err != nil {
			return render.Fail(c, errors.Wrap(err, "参数解析失败"))
		}
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			return render.Fail(c, errors.Wrap(err, "参数验证失败"))
		}
		// 是否绑定过
		hasBind := detail.HasBindUIDWechatUID(req.WechatUserID)
		if hasBind {
			return render.Fail(c, errors.New("该用户已绑定过关联关系"))
		}
		_, err := assets.Find(req.UID)
		// 是否存在资产记录
		notHasAssets := errors.Is(err, gorm.ErrRecordNotFound)
		if err != nil && !notHasAssets {
			return render.Fail(c, errors.Wrap(err, "用户资产异常"))
		}
		// 不存在则创建
		if notHasAssets {
			err = assets.Create(&core.UserAssets{
				UID: req.UID,
			})
		}
		// 绑定关联关系
		err = detail.BindUIDWechatUID(req.UID, req.WechatUserID)
		if err != nil {
			return render.Fail(c, errors.Wrap(err, "绑定失败"))
		}
		totalmoney, totalService, err := detail.FindMoneyServicePointSum(req.WechatUserID)
		if err != nil {
			return render.Success(c, req)
		}
		err = assets.IncrPoint(req.UID, totalmoney, totalService)
		if err != nil {
			return render.Fail(c, errors.Wrap(err, "积分添加失败"))
		}

		return render.Success(c, req)
	}
}
