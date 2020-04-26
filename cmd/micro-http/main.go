package main

import (
	"go-micro-http/internal/app/micro-http/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	service.UseDummyRouting(e)

	e.Logger.Fatal(e.Start(":8080"))
}
