package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("./Site/*.html")),
	}
}

type count struct {
	Count int
}

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())

	count := count{Count: 0}
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "newindex", count)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func GetMatchData(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
