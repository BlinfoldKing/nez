package handler

import (
	"fmt"
	"nez/lib"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Result = lib.Result[any, error]

func Ok[T any](value T) Result {
	return lib.ResultOk[any, error](value)
}

func Err[E error](err E) Result {
	return lib.ResultErr[any, error](err)
}

func Handler[T any](handler func(request T) Result) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(T)
		if err := c.Bind(req); err == nil {
			lib.Try(func() Result {
				validate := validator.New()
				err := validate.Struct(req)
				if err != nil {
					return Err(err)
				}

				return handler(*req)
			}).Catch(func(err error) {
				fmt.Println(err)
				c.JSON(500, echo.Map{
					"error": err.Error(),
				})
			}).Finally(func(value any) Result {
				c.JSON(200, echo.Map{
					"data": value,
				})

				return Ok(value)
			})
		}

		return nil
	}
}
