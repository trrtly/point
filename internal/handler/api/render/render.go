package render

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"point/internal/handler/api/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// indent the json-encoded API responses
var indent bool

func init() {
	indent, _ = strconv.ParseBool(
		os.Getenv("HTTP_JSON_INDENT"),
	)
}

// Response defines reponse context for the api
type Response struct {
	// `code` 错误码
	// 全局错误码说明：
	// `1001` 用户不存在
	Code int64 `json:"code" example:"200"`
	// `msg` 错误信息
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = errors.New("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = errors.New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = errors.New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = errors.New("Not Found")

	// ErrNotImplemented is returned when an endpoint is not implemented.
	ErrNotImplemented = errors.New("Not Implemented")
)

// ErrorCode writes the json-encoded error message to the response.
func ErrorCode(c *fiber.Ctx, err error, status int) {
	JSON(c, &errors.Error{Message: err.Error()}, status)
}

// InternalError writes the json-encoded error message to the response
// with a 500 internal server error.
func InternalError(c *fiber.Ctx, err error) {
	ErrorCode(c, err, 500)
}

// InternalErrorf writes the json-encoded error message to the response
// with a 500 internal server error.
func InternalErrorf(c *fiber.Ctx, format string, a ...interface{}) {
	ErrorCode(c, fmt.Errorf(format, a...), 500)
}

// NotImplemented writes the json-encoded error message to the
// response with a 501 not found status code.
func NotImplemented(c *fiber.Ctx, err error) {
	ErrorCode(c, err, 501)
}

// NotFound writes the json-encoded error message to the response
// with a 404 not found status code.
func NotFound(c *fiber.Ctx, err error) {
	ErrorCode(c, err, 404)
}

// NotFoundf writes the json-encoded error message to the response
// with a 404 not found status code.
func NotFoundf(c *fiber.Ctx, format string, a ...interface{}) {
	ErrorCode(c, fmt.Errorf(format, a...), 404)
}

// Unauthorized writes the json-encoded error message to the response
// with a 401 unauthorized status code.
func Unauthorized(c *fiber.Ctx, err error) {
	ErrorCode(c, err, 401)
}

// Forbidden writes the json-encoded error message to the response
// with a 403 forbidden status code.
func Forbidden(c *fiber.Ctx, err error) {
	ErrorCode(c, err, 403)
}

// BadRequest writes the json-encoded error message to the response
// with a 400 bad request status code.
func BadRequest(c *fiber.Ctx, err error) {
	ErrorCode(c, err, 400)
}

// BadRequestf writes the json-encoded error message to the response
// with a 400 bad request status code.
func BadRequestf(c *fiber.Ctx, format string, a ...interface{}) {
	ErrorCode(c, fmt.Errorf(format, a...), 400)
}

// JSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func JSON(c *fiber.Ctx, v interface{}, status int) error {
	return c.Status(status).JSON(v)
}

// Success reponse an json-encoded api success data
func Success(c *fiber.Ctx, v interface{}) error {
	logrus.WithFields(
		logrus.Fields{
			"uri":     string(c.Request().URI().RequestURI()),
			"request": c.Request(),
			"data":    v,
		},
	).Infoln("api success")
	return JSON(c, &Response{
		Code: http.StatusOK,
		Data: v,
	}, http.StatusOK)
}

// Fail reponse an json-encoded api fail data
func Fail(c *fiber.Ctx, err error, code ...int64) error {
	if len(code) == 0 {
		code = append(code, errors.ApiError)
	}
	logrus.WithFields(
		logrus.Fields{
			"uri":     string(c.Request().URI().RequestURI()),
			"request": c.Request(),
			"err":     err,
			"code":    code,
		},
	).Errorln("api fail")
	return JSON(c, &Response{
		Code: code[0],
		Msg:  err.Error(),
	}, http.StatusOK)
}
