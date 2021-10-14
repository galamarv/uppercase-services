package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"uppercase/vo"

	"github.com/labstack/echo/v4"
)

func UpperCase(c echo.Context) error {
	data := vo.UpperCaseData{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	//log.Printf(err)
	strbaru := data.Text
	res := strings.ToUpper(strbaru)

	fmt.Printf(res)
	resp := vo.UpperCaseResponse{
		Text: res,
	}
	return c.JSON(http.StatusOK, resp)
}
