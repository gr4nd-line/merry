package http

import (
	"github.com/gr4nd-line/merry/internal/http/handlers"
	"github.com/labstack/echo/v5"
)

func AddRoutes(e *echo.Echo) {

	merry := e.Group("/merry")

	merry.GET("", handlers.HealthCheck)

}
