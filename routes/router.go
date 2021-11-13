package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"skeleton-echo/config"
	"skeleton-echo/middleware"
	"skeleton-echo/utils/session"
)

func Api(e *echo.Echo, db *gorm.DB) {
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(db)

	adminGroup := e.Group("/admin", middleware.SessionMiddleware(session.Manager))
	g := adminGroup.Group("/v1", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))


	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	})
	authController := config.InjectAuthController(db)
	e.GET("/login", authController.Index)
	e.POST("/do-login", authController.Login)
	e.POST("/logout", authController.Logout)

	{
		invenGroup := g.Group("/inventaris")
		dashboardController := config.InjectDashboardController(db)
		invenGroup.GET("", dashboardController.Index)
		invenGroup.GET("/add", dashboardController.Add)
		invenGroup.POST("/store", dashboardController.AddData)
		e.GET("/admin/v1/inventaris/generate", dashboardController.GenerateExcel)
		invenGroup.GET("/tables", dashboardController.GetDetail)
		invenGroup.DELETE("/delete/:id", dashboardController.Delete)
		//g.GET("/table", dashboardController.GetData)
		//g.GET("/table/:id", dashboardController.Detail)
		e.GET("/admin/v1/inventaris/update/:id", dashboardController.Update)
		invenGroup.POST("/do-update/:id", dashboardController.DoUpdate)

	}


	m := g.Group("/master-data", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	provController := config.InjectMasterController(db)
	p := m.Group("/provinsi", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	p.GET("", provController.Index)
	//p.GET("/add", masterController.Store)
	p.GET("/table", provController.GetDetail)
	p.GET("/update/:id", provController.Update)
	e.POST("/prov/doupdate/:id", provController.DoUpdate)
	e.POST("/prov/add", provController.AddData)
	e.POST("/prov/:id",provController.FindByID)
	e.DELETE("/prov/:id/delete",provController.Delete)


	kabController := config.InjectKabController(db)
	kb := m.Group("/kab", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	kb.GET("", kabController.Index)
	kb.GET("/table", kabController.GetDetail)
	kb.GET("/update/:id", kabController.Update)
	e.POST("/kab/doupdate/:id", kabController.DoUpdate)
	e.POST("/kab/addkab", kabController.AddData)
	e.POST("kab/:id",kabController.FindByID)
	e.DELETE("/kab/:id/delete",kabController.Delete)

	kecController := config.InjectKecController(db)
	kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	kc.GET("", kecController.Index)
	kc.GET("/table", kecController.GetDetail)
	kc.GET("/update/:id", kecController.Update)
	e.POST("/kec/doupdate/:id", kecController.DoUpdate)
	e.POST("/kec/addkec", kecController.AddData)
	e.POST("/kec/:id",kecController.FindByID)
	e.DELETE("/kec/:id/delete",kecController.Delete)

	//statusController := config.InjectStatusController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	//m.GET("/status-legal", statusController.Index)
	//m.GET("/status-legal/add", statusController.Store)
	//m.GET("/status-legal/update",statusController.Detail)
	//m.POST("/status-legal/update/:id",kecController.DoUpdate)
	//m.POST("/status-legal/create",kecController.AddData)
	//m.DELETE("/status-legal/delete/:id", kecController.Delete)

	pengurusController := config.InjectPengurusController(db)
	pengurusGroup := m.Group("/pengurus", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	//m.GET("/kepengurus", pengurusController.Index)
	pengurusGroup.GET("/add",pengurusController.Store)
	//kc.GET("/table",kecController.GetDetail)
	//kc.GET("/update/:id",kecController.Update)
	//kc.POST("/update/:id",kecController.DoUpdate)
	//kc.POST("/create",kecController.AddData)

	//irigasiController := config.InjectIrigasiController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	//m.GET("/tk-irigasi", irigasiController.Index)
	//m.GET("/tk-irigasi/update/:id",irigasiController.Update)
	//m.POST("/tk-irigasi/updates/:id",irigasiController.DoUpdate)
	//m.POST("/tk-irigasi/create",irigasiController.AddData)
	//m.DELETE("/tk-irigasi/delete/:id", irigasiController.Delete)

	//taniController := config.InjectPertanianController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	//m.GET("/tk-tani", taniController.Index)
	//m.GET("/tk-tani/update/:id",taniController.Update)
	//m.POST("/tk-tani/updates/:id",taniController.DoUpdate)
	//m.POST("/tk-tani/create",taniController.AddData)
	//m.DELETE("/tk-tani/delete/:id", taniController.Delete)

	//usersController := config.InjectUsersController(db)
	//kc := m.Group("/kec", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	//m.GET("/users", usersController.Index)
	//kc.GET("/add",kecController.Store)
	//kc.GET("/table",kecController.GetDetail)
	//kc.GET("/update/:id",kecController.Update)
	//kc.POST("/update/:id",kecController.DoUpdate)
	//kc.POST("/create",kecController.AddData)
}