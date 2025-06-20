package main

import (
	"PersonalPortfolio/handlers"

	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New();

	e.GET("/", handlers.HomeHandler)
	e.Static("/assets", "static/assets")

	e.Logger.Fatal(e.Start(":8080"))
}


