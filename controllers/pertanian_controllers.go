package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"skeleton-echo/models"
	"skeleton-echo/request"
	"skeleton-echo/services"
)

type TaniDataController struct {
	BaseFrontendController
	Controller
	service *services.TaniDataService
}

func NewTaniDataController(services *services.TaniDataService) TaniDataController {
	return TaniDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Pertanian",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *TaniDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/tk-tani",
	}
	return Render(ctx, "Home", "teknik-pertanian/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *TaniDataController) Update(ctx echo.Context) error {
	id := ctx.Param("id")

	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	dataTani := models.TeknikPertanian{
		ID:        data.ID,
		PolaTanam: data.PolaTanam,
		UsahaTani: data.UsahaTani,
	}
	return ctx.JSON(http.StatusOK, &dataTani)
}

func (c *TaniDataController) AddData(ctx echo.Context) error {
	var entity request.TeknikTaniReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	_, err := c.service.Create(entity)
	//entity.CreatedAt = time.Now()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(302, "/inventaris/v1/master-data/tk-tani")
}

func (c *TaniDataController) DoUpdate(ctx echo.Context) error {
	var entity request.TeknikTaniReq
	id := ctx.Param("id")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	data, err := c.service.UpdateById(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	fmt.Println(data)
	return ctx.Redirect(302, "/inventaris/v1/master-data/tk-tani")
}

func (c *TaniDataController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx, nil)
}
