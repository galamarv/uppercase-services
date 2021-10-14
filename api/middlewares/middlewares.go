package middlewares

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMainMiddleWares(e *echo.Echo) {
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("oval")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("KompasGramedia")) == 1 {
			return true, nil
		}
		return false, echo.NewHTTPError(http.StatusUnauthorized, "username atau password salah")
	}))
}
