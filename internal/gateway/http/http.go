package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

func (e Error) Error() string {
	return e.Err
}
func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}
func ErrUnAuthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "unauthorized request",
	}
}

func ErrNotResourceNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  res + " resource not found",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid JSON request",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid id given",
	}
}

func ServerInit() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if apiError, ok := err.(Error); ok {
				return c.Status(apiError.Code).JSON(apiError)
			}
			apiError := NewError(http.StatusInternalServerError, err.Error())
			return c.Status(apiError.Code).JSON(apiError)
		},
	})
	registerRoutes(app)
	if err := app.Listen(*listenAddr); err != nil {
		panic(err)
	}
}
