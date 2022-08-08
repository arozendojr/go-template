package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Text string
	Code string
}

func Root(c echo.Context) error {
	m := Message{"Hello World", "200"}
	return c.JSON(http.StatusOK, m)

}
