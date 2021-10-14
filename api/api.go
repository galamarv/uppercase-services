package api

import (
	"uppercase/handler"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.POST("/api/uppercase", handler.UpperCase)

}
