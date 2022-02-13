package main

import (
	"webcache/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	h := interfaces.NewHandler()
	e.Use(middleware.Logger())

	e.GET("/", h.Root)
	e.GET("/cache.png", h.WithCacheHeader)
	e.GET("/png", h.PNG)

	e.Logger.Fatal(e.Start(":1323"))
}
