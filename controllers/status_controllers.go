package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"skeleton-echo/models"
	"skeleton-echo/request"
	"skeleton-echo/services"
)

type StatusDataController struct {
	BaseFrontendController
	Controller
	service *services.StatusDataService
}

func NewStatusDataController(services *services.StatusDataService) StatusDataController {
	return StatusDataController{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Status Legal",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *StatusDataController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Status Legal",
		"link": "/inventaris/v1/master-data/status-legal/tables",
	}
	return Render(ctx, "Home", "status-legal/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *StatusDataController) Store(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/master-data/status-legal/add",
	}
	return Render(ctx, "Home", "status-legal/add", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *StatusDataController) Detail(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/master-data/status-legal/detail/:id",
	}
	dataStatus := models.StatusLegal{
		ID:                  data.ID,
		TahunPembentukan:    data.TahunPembentukan,
		LamTahunPembentukan: data.LamTahunPembentukan,
		LamKplDesa:          data.LamKplDesa,
		SKBupati:            data.SKBupati,
		LamSKBupati:         data.LamSKBupati,
		AkteNotaris:         data.AkteNotaris,
		LamAkteNotaris:      data.LamAkteNotaris,
		NoPendaftaran:       data.NoPendaftaran,
		LamPendaftaran:      data.LamPendaftaran,
	}
	return Render(ctx, "Home", "status-legal/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), dataStatus)
}

func (c *StatusDataController) AddData(ctx echo.Context) error {
	var entity request.StatusLegalReq

	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	_, err := c.service.Create(entity)
	//entity.CreatedAt = time.Now()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(302, "/inventaris/v1/master-data/status-legal")
}

func (c *StatusDataController) DoUpdate(ctx echo.Context) error {
	var entity request.StatusLegalReq
	id := ctx.Param("id")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	data, err := c.service.UpdateById(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	fmt.Println(data)
	return ctx.Redirect(302, "/inventaris/v1/master-data/status-legal")
}

func (c *StatusDataController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx,nil)
}
