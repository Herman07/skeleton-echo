package home

import (
	"github.com/labstack/echo"
	"net/http"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Ini adalah Halaman Home")


}

//func Page(c echo.Context) error {
//	//render only file, must full name with extension
//	fmt.Println("masuk gak ya")
//	return c.Render(http.StatusOK, "index", echo.Map{"title": "Page file title!!"})
//}