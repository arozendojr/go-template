package rest

import (
	"github.com/arozendojr/go-rest/internal/application/rest/controllers"
	"github.com/arozendojr/go-rest/internal/application/rest/middleware"

	"github.com/labstack/echo"
	middlewar "github.com/labstack/echo/middleware"
)

func Start() {

	server := echo.New()

	server.Use(middlewar.Logger())
	server.Use(middleware.Header)
	server.Use(middlewar.Recover())

	server.GET("/", controllers.Root)

	server.Logger.Fatal(server.Start(":8080"))
}
