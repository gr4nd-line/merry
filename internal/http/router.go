package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func AddRoutes(e *echo.Echo) {

	merry := e.Group("/merry")

	merry.GET("", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message":   "ok",
			"app_name":  "merry",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

}
