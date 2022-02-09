package main

import (
	"webcache/interfaces"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := interfaces.NewHandler()

	e.GET("/", h.Root)

	e.Logger.Fatal(e.Start(":1323"))
}
