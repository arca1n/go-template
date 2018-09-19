package services

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var EchoServer *echo.Echo
var Logger echo.Logger

func init() {
	fmt.Println("Initializing Logger")
	Logger = log.New(os.Getenv("hostname"))
	Logger.Info("Server and Logger initialized")
	EchoServer = echo.New()
	EchoServer.Logger = Logger
}
