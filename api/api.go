package api

import (
	"github.com/galamarv/uppercase-services/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)

}
