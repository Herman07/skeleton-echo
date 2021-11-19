package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"skeleton-echo/request"
	"skeleton-echo/services"
	"strconv"
	"time"
)

type UsersDataController struct {
	BaseFrontendController
	Controller
	service *services.UsersDataService
}

func NewUsersDataController(services *services.UsersDataService) UsersDataController {
	return UsersDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Users",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *UsersDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Users",
		"link": "/admin/v1/users",
	}
	return Render(ctx, "Home", "users/create", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *UsersDataController) AddData(ctx echo.Context) error {
	var entity request.UsersReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	user, err := c.service.Create(entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	file, _ := ctx.FormFile("foto")

	src, _ := file.Open()
	defer src.Close()

	// Destination
	t := time.Now().UnixNano()
	nf := entity.Nama + "_" + strconv.FormatInt(t, 10) + "_" + file.Filename
	nama := "static/image/" + nf
	dst, _ := os.Create(nama)
	defer dst.Close()

	// Copy
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Error("[Error] ", err)
		return c.InternalServerError(ctx, err)
	}
	_, err = c.service.CreateAkun(entity, nf, user.ID)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(302, "/admin/v1/inventaris")
}

