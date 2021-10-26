package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"skeleton-echo/models"
	"skeleton-echo/request"
	"skeleton-echo/services"
	"strconv"
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

//
func (c *StatusDataController) GetDetail(ctx echo.Context) error {

	draw, err := strconv.Atoi(ctx.Request().URL.Query().Get("draw"))
	start, err := strconv.Atoi(ctx.Request().URL.Query().Get("start"))
	search := ctx.Request().URL.Query().Get("search[value]")
	length, err := strconv.Atoi(ctx.Request().URL.Query().Get("length"))
	order, err := strconv.Atoi(ctx.Request().URL.Query().Get("order[0][column]"))
	orderName := ctx.Request().URL.Query().Get("columns[" + strconv.Itoa(order) + "][name]")
	orderAscDesc := ctx.Request().URL.Query().Get("order[0][dir]")

	recordTotal, recordFiltered, data, err := c.service.QueryDatatable(search, orderAscDesc, orderName, length, start)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	//var createdAt string
	var action string
	listOfData := make([]map[string]interface{}, len(data))
	for k, v := range data {
		action = `<a href="/inventaris/v1/master-data/status-legal/detail` + (v.ID) + `" class="btn btn-success btn-bold btn-upper" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-edit"></i></a>
		<a href="javascript:;" onclick="Delete('` + v.ID + `')" class="btn btn-danger btn-bold btn-upper" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-trash"></i></a>`
		//time := v.CreatedAt
		//createdAt = time.Format("2006-01-02")
		listOfData[k] = map[string]interface{}{
			"id_status_legal":   v.ID,
			"tahun_pembentukan": v.TahunPembentukan,
			"no_sk_bupati":      v.SKBupati,
			"akte_notaris":      v.AkteNotaris,
			"no_pendaftaran":    v.NoPendaftaran,
			"action":            action,
		}
	}
	result := models.ResponseDatatable{
		Draw:            draw,
		RecordsTotal:    recordTotal,
		RecordsFiltered: recordFiltered,
		Data:            listOfData,
	}
	return ctx.JSON(http.StatusOK, &result)
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
