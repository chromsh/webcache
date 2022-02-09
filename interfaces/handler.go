package interfaces

import (
	"net/http"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Root(c echo.Context) error
}
type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

const (
	CookieCount = "count"
)

func (h *handler) Root(c echo.Context) error {
	cookie, err := c.Cookie(CookieCount)
	if err != nil {
		cookie = new(http.Cookie)
		cookie.Name = "count"
		cookie.Value = "0"
	}
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return internalServerError(c, err)
	}
	cookie.Value = strconv.Itoa(count + 1)
	c.SetCookie(cookie)
	return c.File(path.Join("static", "html", "root.html"))
}

func internalServerError(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
