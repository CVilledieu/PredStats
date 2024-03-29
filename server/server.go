package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func GetMatchData(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
