package main

import (
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/config"
	"github.com/TOUCHTHANAWAT/Mini_Spatial_Data_Platform/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitDB()

	e := echo.New()
	e.Use(middleware.CORS())

	// e.Use(middleware.Logger()) //เก็บและแสดง log
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\033[32m${status}\033[0m | ${latency_human} | ${method} ${uri}\n",
	}))
	e.Use(middleware.Recover()) //กัน server ล่ม

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
