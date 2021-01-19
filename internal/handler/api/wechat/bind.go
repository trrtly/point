package wechat

import (
	"point/internal/core"
	"point/internal/handler/api/render"

	"github.com/pkg/errors"

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
		err := detail.BindUIDWechatUID(req.UID, req.WechatUserID)
		if err != nil {
			return render.Fail(c, errors.Wrap(err, "绑定失败"))
		}
		return render.Success(c, req)
	}
}
