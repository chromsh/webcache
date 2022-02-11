package interfaces

import (
	"net/http"
	"path"
	"strconv"
	"webcache/images"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Root(c echo.Context) error
	NoCache(c echo.Context) error
	PNG(c echo.Context) error
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

func (h *handler) NoCache(c echo.Context) error {
	return responseWithHeaders(c, "no-cache")
}

func (h *handler) PNG(c echo.Context) error {
	n, err := strconv.Atoi(c.QueryParam("n"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid n")
	}
	data, err := images.PNG(n)
	if err != nil {
		return internalServerError(c, err)
	}
	return c.Blob(http.StatusOK, "image/png", data)
}

func responseWithHeaders(c echo.Context, cacheControl string) error {
	cookie, err := c.Cookie(CookieCount)
	if err != nil {
		return internalServerError(c, err)
	}
	num, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return internalServerError(c, err)
	}
	data, err := images.PNG(num)
	if err != nil {
		return internalServerError(c, err)
	}
	c.Response().Header().Set("Cache-Control", cacheControl)
	return c.Blob(http.StatusOK, "image/png", data)
}

func internalServerError(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
