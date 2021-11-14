package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
	"os"
	"skeleton-echo/models"
	"skeleton-echo/request"
	"skeleton-echo/services"
	"strconv"
	"time"
)

type P3Controller struct {
	BaseFrontendController
	Controller
	service *services.P3Service
}

func NewP3Controller(services *services.P3Service) P3Controller {
	return P3Controller{
		service: services,
		BaseFrontendController: BaseFrontendController{
			Menu:        "Dashboard",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}
func (c *P3Controller) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/admin/v1/inventaris",
	}
	return Render(ctx, "Home", "p3a/index", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *P3Controller) Add(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/admin/v1/inventaris/add",
	}
	return Render(ctx, "Home", "p3a/create", c.Menu, append(c.BreadCrumbs, breadCrumbs), nil)
}

func (c *P3Controller) GetDetail(ctx echo.Context) error {

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
		action = `<a href="/admin/v1/inventaris/update/` + (v.ID) + `" class="btn btn-primary" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-edit"></i></a>
		<a href="/admin/v1/inventaris/detail/` + (v.ID) + `" class="btn btn-primary" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-eye"></i></a>
		<button onclick="Delete('` + v.ID + `')" class="btn btn-danger" title="Delete" style="text-decoration: none;font-weight: 100;color: white;/* width: 80px; */"><i class="fa fa-trash"></i></button>`
		//time := v.CreatedAt
		//createdAt = time.Format("2006-01-02")
		listOfData[k] = map[string]interface{}{
			"id_p3a":              v.ID,
			"no_urut":             v.NoUrut,
			"nama_p3a":            v.NamaP3A,
			"jumlah_p3a":          v.JumlahP3A,
			"nama_daerah_irigasi": v.DaerahIrigasi,
			"luas_wilayah":        v.LuasWilayah,
			"luas_layanan_p3a":    v.LuasLayananP3A,
			"keterangan":          v.Keterangan,
			//"created_at":          createdAt,
			"action": action,
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

//func (c *P3Controller) GetData(ctx echo.Context) error {
//	dataReq := models.Inventaris{}
//
//	data, err := c.service.GetAll(dataReq)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//
//	return c.Ok(ctx, data)
//}
//func (c *P3Controller) Detail(ctx echo.Context) error {
//	id := ctx.Param("id")
//	data, err := c.service.FindById(id)
//	if err != nil {
//		return c.InternalServerError(ctx, err)
//	}
//	return c.Ok(ctx, data)
//}

func (c *P3Controller) AddData(ctx echo.Context) error {
	var entity request.RequestInventaris
	if err := ctx.Bind(&entity); err != nil {
		log.Error("[Error] ", err)
		//return ctx.JSON(500, echo.Map{"message": "error binding data"})
	}

	name := []string{"lampiran_tahun_pembentukan", "diket_kep_dc", "lampiran_sk_bupati", "lampiran_akte_notaris", "lampiran_pendaftaran", "lampiran_ad_art", "lampiran_sekretariat"}
	var namaFile []string
	for i := range name {
		file, _ := ctx.FormFile(name[i])

		src, _ := file.Open()
		defer src.Close()

		// Destination
		t := time.Now().UnixNano()
		nf := name[i] + "_" + strconv.FormatInt(t, 10) + "_" + file.Filename
		nama := "static/image/" + nf
		dst, _ := os.Create(nama)
		defer dst.Close()

		// Copy
		_, err := io.Copy(dst, src)
		if err != nil {
			log.Error("[Error] ", err)
			return c.InternalServerError(ctx, err)
		}
		i++
		namaFile = append(namaFile, nf)
	}
	fmt.Println("Name File  : ", namaFile)

	//Store Data Status Legal
	statusLegal, err := c.service.CreateStatusLegal(entity, namaFile)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Store Data Kepengurusan
	pengurus, err := c.service.CreatePengurus(entity, namaFile)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Store Data Teknik Irigasi
	irigasi, err := c.service.CreateIrigasi(entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Store Data Teknik Pertanian
	pertanian, err := c.service.CreatePertanian(entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Store Data to Table p3a
	_, err = c.service.CreateDataP3a(entity, statusLegal.ID, pengurus.ID, irigasi.ID, pertanian.ID)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	return ctx.Redirect(http.StatusFound, "/admin/v1/inventaris")
}

func (c *P3Controller) GenerateExcel(ctx echo.Context) error {
	// Get Data Export
	data, err := c.service.GetDataExport()
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	//style := excelize.Style{
	//	Border:        nil,
	//	Fill:          excelize.Fill{},
	//	Font:          nil,
	//	Alignment:     nil,
	//	Protection:    nil,
	//	NumFmt:        0,
	//	DecimalPlaces: 0,
	//	CustomNumFmt:  nil,
	//	Lang:          "",
	//	NegRed:        false,
	//}
	f := excelize.NewFile()
	_, _ = f.NewConditionalStyle("center")
	style, _ := f.NewStyle(`
		{
			"alignment":{"horizontal":"center"
		}, 
        "fill": {
				"type": "pattern",
				"color": ["#E0EBF5"],
				"pattern": 1
			}`)

	f.SetCellStyle("Sheet1", "A1", "AF100", style)

	// Create a new sheet.
	index := f.NewSheet("Data Export")
	// Set value of a cell.
	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Provinsi/Kabupaten")
	_ = f.SetCellValue("Sheet1", "C1", "Kecamatan")
	_ = f.SetCellValue("Sheet1", "D1", "Nama Daerah Imigrasi")
	_ = f.SetCellValue("Sheet1", "E1", "Luas Wilayah (Ha)")
	_ = f.SetCellValue("Sheet1", "F1", "Jumlah P3A")
	_ = f.SetCellValue("Sheet1", "G1", "Nama P3A")
	_ = f.SetCellValue("Sheet1", "H1", "Luas Layanan P3A (Ha)")

	// Status Legal
	_ = f.SetCellValue("Sheet1", "I1", "Status Legal P3A")
	_ = f.SetCellValue("Sheet1", "I2", "Tahun Pembentukan")
	_ = f.SetCellValue("Sheet1", "J2", "Diketahui Kepala Desa / Camat")
	_ = f.SetCellValue("Sheet1", "K2", "SK Bupati")
	_ = f.SetCellValue("Sheet1", "L2", "Akte Notaris")
	_ = f.SetCellValue("Sheet1", "M2", "Terdaftar di Pengadilan Negeri / Kemenkum HAM")
	_ = f.MergeCell("Sheet1", "I1", "M1")

	// Kepengurusan
	_ = f.SetCellValue("Sheet1", "N1", "Kepengurusan")
	_ = f.SetCellValue("Sheet1", "N2", "Ketua")
	_ = f.SetCellValue("Sheet1", "O2", "Wakil Ketua")
	_ = f.SetCellValue("Sheet1", "P2", "Sekertaris")
	_ = f.SetCellValue("Sheet1", "Q2", "Bendahara")
	_ = f.SetCellValue("Sheet1", "R2", "Seksi (L/P)")
	_ = f.SetCellValue("Sheet1", "U2", "Jumlah Anggota")
	_ = f.SetCellValue("Sheet1", "N3", "L/P")
	_ = f.SetCellValue("Sheet1", "O3", "L/P")
	_ = f.SetCellValue("Sheet1", "P3", "L/P")
	_ = f.SetCellValue("Sheet1", "Q3", "L/P")
	_ = f.SetCellValue("Sheet1", "R3", "Teknik")
	_ = f.SetCellValue("Sheet1", "S3", "O & P")
	_ = f.SetCellValue("Sheet1", "T3", "Bisnis")
	_ = f.MergeCell("Sheet1", "N1", "U1")
	_ = f.MergeCell("Sheet1", "R2", "T2")
	_ = f.MergeCell("Sheet1", "U2", "U3")

	_ = f.SetCellValue("Sheet1", "V1", "AD/ART")
	_ = f.SetCellValue("Sheet1", "W1", "Sekertariat")
	_ = f.SetCellValue("Sheet1", "X1", "Persentase (%) perempuan")
	_ = f.SetCellValue("Sheet1", "Y1", "areal tersier (ha)")
	_ = f.SetCellValue("Sheet1", "Z1", "Pengisian Buku")
	_ = f.SetCellValue("Sheet1", "AA1", "Iuran")

	//Teknik irigasi
	_ = f.SetCellValue("Sheet1", "AB1", "Teknik Irigasi")
	_ = f.SetCellValue("Sheet1", "AB2", "Operasi")
	_ = f.SetCellValue("Sheet1", "AC2", "Partisipatif")
	_ = f.MergeCell("Sheet1", "AB1", "AC1")
	_ = f.MergeCell("Sheet1", "AB2", "AB3")
	_ = f.MergeCell("Sheet1", "AC2", "AC3")

	//Teknik Pertanian
	_ = f.SetCellValue("Sheet1", "AD1", "Teknik Pertanian")
	_ = f.SetCellValue("Sheet1", "AD2", "Pola Tanam")
	_ = f.SetCellValue("Sheet1", "AE2", "Usaha Tani")
	_ = f.MergeCell("Sheet1", "AD1", "AE1")
	_ = f.MergeCell("Sheet1", "AD2", "AD3")
	_ = f.MergeCell("Sheet1", "AE2", "AE3")

	// Keterangan
	_ = f.SetCellValue("Sheet1", "AF1", "Keterangan")
	_ = f.MergeCell("Sheet1", "AF1", "AF3")

	// ROW MERGE
	for i := 0; i < 9; i++ {
		_ = f.MergeCell("Sheet1", string(rune('A'-1+i))+"1", string(rune('A'-1+i))+"3")
	}

	for i := 0; i < 6; i++ {
		_ = f.MergeCell("Sheet1", string(rune('I'-1+i))+"2", string(rune('I'-1+i))+"3")
	}

	for i := 0; i < 5; i++ {
		_ = f.MergeCell("Sheet1", string(rune('W'-1+i))+"1", string(rune('W'-1+i))+"3")
	}
	_ = f.MergeCell("Sheet1", "AA1", "AA3")

	for i, v := range data {
		_ = f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+4), i+1)
		i = i + 4
		_ = f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), v.NamaProv)
		_ = f.SetCellValue("Sheet1", "C"+strconv.Itoa(i), v.NamaKec)
		_ = f.SetCellValue("Sheet1", "D"+strconv.Itoa(i), v.DaerahIrigasi)
		_ = f.SetCellValue("Sheet1", "E"+strconv.Itoa(i), v.LuasWilayah)
		_ = f.SetCellValue("Sheet1", "F"+strconv.Itoa(i), v.JumlahP3A)
		_ = f.SetCellValue("Sheet1", "G"+strconv.Itoa(i), v.NamaP3A)
		_ = f.SetCellValue("Sheet1", "H"+strconv.Itoa(i), v.LuasLayananP3A)
		_ = f.SetCellValue("Sheet1", "I"+strconv.Itoa(i), v.TahunPembentukan)
		_ = f.SetCellValue("Sheet1", "J"+strconv.Itoa(i), v.LamKplDesa)
		_ = f.SetCellValue("Sheet1", "K"+strconv.Itoa(i), v.SKBupati)
		_ = f.SetCellValue("Sheet1", "L"+strconv.Itoa(i), v.AkteNotaris)
		_ = f.SetCellValue("Sheet1", "M"+strconv.Itoa(i), v.NoPendaftaran)
		_ = f.SetCellValue("Sheet1", "N"+strconv.Itoa(i), v.Ketua)
		_ = f.SetCellValue("Sheet1", "O"+strconv.Itoa(i), v.Wakil)
		_ = f.SetCellValue("Sheet1", "P"+strconv.Itoa(i), v.Sekretaris)
		_ = f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i), v.Bendahara)
		_ = f.SetCellValue("Sheet1", "R"+strconv.Itoa(i), v.SekTeknik)
		_ = f.SetCellValue("Sheet1", "S"+strconv.Itoa(i), v.SekOP)
		_ = f.SetCellValue("Sheet1", "T"+strconv.Itoa(i), v.SekBisnis)
		_ = f.SetCellValue("Sheet1", "U"+strconv.Itoa(i), v.JumlahAnggota)
		_ = f.SetCellValue("Sheet1", "V"+strconv.Itoa(i), v.NoADRT)
		_ = f.SetCellValue("Sheet1", "W"+strconv.Itoa(i), v.Sekretariat)
		_ = f.SetCellValue("Sheet1", "X"+strconv.Itoa(i), v.PresentasiPerempuanP3A)
		_ = f.SetCellValue("Sheet1", "Y"+strconv.Itoa(i), v.ArealTersier)
		_ = f.SetCellValue("Sheet1", "Z"+strconv.Itoa(i), v.PengisianBuku)
		_ = f.SetCellValue("Sheet1", "AA"+strconv.Itoa(i), v.Iuran)
		_ = f.SetCellValue("Sheet1", "AB"+strconv.Itoa(i), v.Operasi)
		_ = f.SetCellValue("Sheet1", "AC"+strconv.Itoa(i), v.Partisipatif)
		_ = f.SetCellValue("Sheet1", "AD"+strconv.Itoa(i), v.PolaTanam)
		_ = f.SetCellValue("Sheet1", "AE"+strconv.Itoa(i), v.UsahaTani)
		_ = f.SetCellValue("Sheet1", "AF"+strconv.Itoa(i), v.Keterangan)

	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	t := time.Now()
	name := "Report - " + t.Format("2006-01-02") + ".xlsx"
	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
	}
	//Delete File
	defer os.Remove(name)

	return ctx.File(name)
}

func (c *P3Controller) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindById(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/inventaris/v1/update/:id",
	}

	return Render(ctx, "Home", "p3a/update", c.Menu, append(c.BreadCrumbs, breadCrumbs), data)
}

func (c *P3Controller) DoUpdate(ctx echo.Context) error {
	var entity request.UpdateInventaris
	id := ctx.Param("id")
	if err := ctx.Bind(&entity); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	var name []string
	fmt.Println("Data Request : ", entity)
	if entity.LamTahunPembentukan != nil {
		name = append(name, "lampiran_tahun_pembentukan")
	}
	if entity.LamKplDesa != nil {
		name = append(name, "diket_kep_dc")
	}
	if entity.LamSKBupati != nil {
		name = append(name, "lampiran_sk_bupati")
	}
	if entity.LamAkteNotaris != nil{
		name = append(name, "lampiran_akte_notaris")
	}
	if entity.LamPendaftaran != nil {
		name = append(name, "lampiran_pendaftaran")
	}
	if entity.LampiranADRT != nil{
		name = append(name, "lampiran_ad_art")
	}
	if entity.LampiranSekretariat != nil {
		name = append(name, "lampiran_sekretariat")
	}
	var namaFile []string
	for i := range name {
		file, _ := ctx.FormFile(name[i])

		src, _ := file.Open()
		defer src.Close()

		// Destination
		t := time.Now().UnixNano()
		nf := name[i] + "_" + strconv.FormatInt(t, 10) + "_" + file.Filename
		nama := "static/image/" + nf
		dst, _ := os.Create(nama)
		defer dst.Close()

		// Copy
		_, err := io.Copy(dst, src)
		if err != nil {
			log.Error("[Error] ", err)
			return c.InternalServerError(ctx, err)
		}
		i++
		namaFile = append(namaFile, nf)
	}
	fmt.Println("Name File  : ", namaFile)
	entity.LamTahunPembentukan = &namaFile[0]
	entity.LamKplDesa = &namaFile[1]
	entity.LamSKBupati = &namaFile[2]
	entity.LamAkteNotaris = &namaFile[3]
	entity.LamPendaftaran = &namaFile[4]
	entity.LampiranADRT = &namaFile[5]
	entity.LampiranSekretariat = &namaFile[6]
	// Update Data Status Legal
	_, err := c.service.UpdateStatusLegal(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Update Data Kepengurusan
	_, err = c.service.UpdatePengurus(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Update Data Teknik Irigasi
	_, err = c.service.UpdateIrigasi(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	// Update Data Teknik Pertanian
	_, err = c.service.UpdatePertanian(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}

	//Update Data to Table p3a
	_, err = c.service.UpdateById(id, entity)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return ctx.Redirect(http.StatusFound, "/admin/v1/inventaris")
}

func (c *P3Controller) Delete(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.service.Delete(id)
	if err != nil {
		return c.InternalServerError(ctx, err)
	}
	return c.Ok(ctx, nil)
}
