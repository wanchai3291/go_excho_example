package UserRoutes

import (
	UserController "go_echo_example/src/user-management/controller"

	"github.com/labstack/echo"
)

func InitUserRoutes(e *echo.Echo) {
	e.GET("users", UserController.CreateUser)
}
