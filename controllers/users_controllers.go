package controllers

import (
	"github.com/labstack/echo/v4"
	"skeleton-echo/services"
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
		"link": "/inventaris/v1/users",
	}
	return Render(ctx, "Home", "users/detail", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

//func (c *KabDataController) Store(ctx echo.Context) error {
//	breadCrumbs := map[string]interface{}{
//		"menu": "Home",
//		"link": "/inventaris/v1/master-data/kab/add",
//	}
//	return Render(ctx, "Home", "master-data/kabupaten/add", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
//}
//func (c *KabDataController) Update(ctx echo.Context) error {
//	id := ctx.Param("id")
//	data, err := c.service.FindById(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//
//	breadCrumbs := map[string]interface{}{
//		"menu": "Home",
//		"link": "/inventaris/v1/master-data/kab/update/:id",
//	}
//	dataKab := models.MasterDataKab{
//		ID:         data.ID,
//		Kabupaten:   data.Kabupaten,
//		IDProv:  data.IDProv,
//	}
//	return Render(ctx, "Home", "master-data/kabupaten/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), dataKab)
//}
//
//func (c *KabDataController) GetDetail(ctx echo.Context) error {
//
//	draw, err := strconv.Atoi(ctx.Request().URL.Query().Get("draw"))
//	start, err := strconv.Atoi(ctx.Request().URL.Query().Get("start"))
//	search := ctx.Request().URL.Query().Get("search[value]")
//	length, err := strconv.Atoi(ctx.Request().URL.Query().Get("length"))
//	order, err := strconv.Atoi(ctx.Request().URL.Query().Get("order[0][column]"))
//	orderName := ctx.Request().URL.Query().Get("columns[" + strconv.Itoa(order) + "][name]")
//	orderAscDesc := ctx.Request().URL.Query().Get("order[0][dir]")
//
//	recordTotal, recordFiltered, data ,err := c.service.QueryDatatable(search,orderAscDesc, orderName, length, start)
//	if err != nil {
//		return ctx.JSON(http.StatusInternalServerError, err.Error())
//	}
//	//var createdAt string
//	var action string
//	listOfData := make([]map[string]interface{}, len(data))
//	for k, v := range data {
//		action = `<a data-toggle="modal" data-target="#modal-update" class="btn btn-success btn-bold btn-upper" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-edit"></i></a>
//		<a href="javascript:;" onclick="Delete('` + v.ID + `')" class="btn btn-danger btn-bold btn-upper" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fas fa-trash"></i></a>`
//		//time := v.CreatedAt
//		//createdAt = time.Format("2006-01-02")
//		listOfData[k] = map[string]interface{}{
//			"nama_kab":    v.Kabupaten,
//			"id_kab":          v.ID,
//			"action": action,
//		}
//	}
//	result := models.ResponseDatatable{
//		Draw:            draw,
//		RecordsTotal:    recordTotal,
//		RecordsFiltered: recordFiltered,
//		Data:            listOfData,
//	}
//	return ctx.JSON(http.StatusOK, &result)
//}
//
//func (c *KabDataController) AddData(ctx echo.Context) error {
//	var entity request.KabReq
//
//	if err := ctx.Bind(&entity); err != nil {
//		return ctx.JSON(400, echo.Map{"message": "error binding data"})
//	}
//	_, err := c.service.Create(entity)
//	//entity.CreatedAt = time.Now()
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	return ctx.Redirect(302, "/inventaris/v1/master-data/kab")
//}
//
//func (c *KabDataController) DoUpdate(ctx echo.Context) error {
//	var entity request.KabReq
//	id := ctx.Param("id_kab")
//	if err := ctx.Bind(&entity); err != nil {
//		return ctx.JSON(400, echo.Map{"message": "error binding data"})
//	}
//	data, err := c.service.UpdateById(id, entity)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	fmt.Println(data)
//	return ctx.Redirect(302, "/inventaris/v1/master-data/kab")
//}
//
//func (c *KabDataController) Delete(ctx echo.Context) error {
//	id := ctx.Param("id_kab")
//
//	err := c.service.Delete(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	return c.Ok(ctx,nil)
//}
