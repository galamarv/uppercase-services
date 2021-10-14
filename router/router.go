package router

import (
	"github.com/galamarv/uppercase-services/api"
	"github.com/galamarv/uppercase-services/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//set all middlewares
	middlewares.SetMainMiddleWares(e)

	//set main routes
	api.MainGroup(e)

	return e
}
