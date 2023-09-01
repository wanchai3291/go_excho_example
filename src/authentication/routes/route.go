package AuthenticationRoute

import (
	AuthenticationController "go_echo_example/src/authentication/controller"

	"github.com/labstack/echo"
)

func InitAuthRoutes(e *echo.Echo) {
	e.POST("/sign-in", AuthenticationController.SignIn)
}
