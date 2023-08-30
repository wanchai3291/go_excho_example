package SampleRoute

import (
	SampleController "go_echo_example/src/sample/controller"

	"github.com/labstack/echo"
)

func InitSampleRoutes(e *echo.Echo) {
	e.POST("/sample", SampleController.Create)
}
