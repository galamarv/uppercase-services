package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/galamarv/uppercase-services/vo"
	"github.com/labstack/echo/v4"
)

func AddCat(c echo.Context) error {
	data := vo.UpperCaseData{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is yout cat %#v", data)
	return c.String(http.StatusOK, "We got your Cat!!!")
}
