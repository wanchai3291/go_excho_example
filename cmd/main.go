package main

import (
	"go_echo_example/configs/database/mongodb"
	"go_echo_example/configs/database/supabase"
	"go_echo_example/pkg/handler"
	AuthenticationRoute "go_echo_example/src/authentication/routes"
	SampleRoute "go_echo_example/src/sample/routes"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	handler.EnvHandler()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	mongodb.Connect()

	supabase.Connect()
	supa := supabase.DB()
	supabaseDB, err := supa.DB()
	if err != nil {
		panic(err)
	}
	supabaseDB.Ping()
	AuthenticationRoute.InitAuthRoutes(e)
	SampleRoute.InitSampleRoutes(e)
	// UserRoutes.InitUserRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
}
