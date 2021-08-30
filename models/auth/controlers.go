package auth

import (
	"github.com/labstack/echo"
	"net/http"
)

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "auth/login.html", map[string]interface{}{
	})

}
