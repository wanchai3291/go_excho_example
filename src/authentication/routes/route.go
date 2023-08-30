package AuthenticationRoute

import (
	AuthenticationController "go_echo_example/src/authentication/controller"

	"github.com/labstack/echo"
)

func InitAuthRoutes(e *echo.Echo) {
	e.GET("users", AuthenticationController.SignIn)
}
