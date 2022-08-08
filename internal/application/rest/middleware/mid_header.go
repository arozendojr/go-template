package middleware

import (
	"fmt"

	"github.com/labstack/echo"
)

func Header(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Action server Header")
		return next(c)
	}
}
