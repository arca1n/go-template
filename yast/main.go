package main

import (
	"app/yast/models"
	"app/yast/services"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance

	// Middleware
	services.EchoServer.Use(middleware.Logger())
	services.EchoServer.Use(middleware.Recover())

	// Routes
	services.EchoServer.GET("/", hello)

	// DB Migrations
	models.BootstrapModels()

	// Start server
	services.EchoServer.Logger.Fatal(services.EchoServer.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
