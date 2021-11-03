package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"skeleton-echo/models"
	"skeleton-echo/request"
	"skeleton-echo/services"
)

type IrigasiDataController struct {
	BaseFrontendController
	Controller
	service *services.IrigasiDataService
}

func NewIrigasiDataController(services *services.IrigasiDataService) IrigasiDataController {
	return IrigasiDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Teknik Irigasi",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *IrigasiDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Teknik Irigasi",
		"link": "/inventaris/v1/master-data/tk-irigasi",
	}
	return Render(ctx, "Home", "teknik-irigasi/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *IrigasiDataController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	dataIrigasi := models.TeknikIrigasi{
		ID:           data.ID,
		Operasi:      data.Operasi,
		Partisipatif: data.Partisipatif,
	}
	return ctx.JSON(http.StatusOK, &dataIrigasi)
}

func (c *IrigasiDataController) AddData(ctx echo.Context) error {
	var entity request.TeknikIrigasiReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}

	_, err := c.service.Create(entity)
	//entity.CreatedAt = time.Now()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(302, "/inventaris/v1/master-data/tk-irigasi")
}

func (c *IrigasiDataController) DoUpdate(ctx echo.Context) error {
	id := ctx.Param("id")

	var entity request.TeknikIrigasiReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	data, err := c.service.UpdateById(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	fmt.Println(data)

	return c.Ok(ctx, data)
}

func (c *IrigasiDataController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx, nil)
}
