package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"personalsite/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer handles rendering templates
type TemplateRenderer struct {
	templates *template.Template
}

// Render method for TemplateRenderer
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Static("/assets", "assets")

	// Register the template renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", handlers.HomeHandler)
	e.POST("/contactrequest", handlers.ContactRequestHandler)
	e.GET("/example/:name/:age", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	name := c.Param("name")
	age := c.Param("age")
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s! You are %v years old!", name, age))
}
