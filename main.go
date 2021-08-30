package main

import (
	"fmt"
	"github.com/foolin/echo-template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"skeleton-echo/routers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = echotemplate.Default()

	e.Static("/skeleton/assets", "static/assets")
	e.Static("/admin/template" , "static/assets/admin_template")

	//register router
	routers.Api(e)

	//e = routers.Main()
	e.GET("/page", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "dashboard/dashboard", nil)
	})
	e.GET("/table", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "table/table", nil)
	})

	dsn := "host=localhost user=postgres password=admin dbname=skeleton port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database Success :", db)
	e.Logger.Fatal(e.Start(":1323"))

}
