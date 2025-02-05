package handlers

import (
	"math/rand/v2"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title":        "Home Page",
		"randomNumber": rand.IntN(100),
	})
}

func ContactRequestHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"title":        "Home Page",
		"randomNumber": rand.IntN(100),
	})
}
