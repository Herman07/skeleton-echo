package routers

import (
	"github.com/labstack/echo"
	"net/http"
	"skeleton-echo/models/auth"
)

func Api(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hellow World")
	})
	e.GET("/login", auth.Login)
	//e.GET("/page", home.Page)
}
