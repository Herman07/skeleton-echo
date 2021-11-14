package routes

import (
	"encoding/json"
	"fmt"
	"github.com/foolin/goview"
	echotemplate "github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"skeleton-echo/config"
	"skeleton-echo/middleware"
	"skeleton-echo/utils/session"
)

func Api(e *echo.Echo, db *gorm.DB) {

	mv := echotemplate.NewMiddleware(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"session": func(ctx echo.Context) session.UserInfo {
				ses, _ := session.Manager.Get(ctx, session.SessionId)
				dataSes, _ := json.Marshal(ses)
				var data session.UserInfo
				userInfo := session.UserInfo{
					ID: data.ID,
					Username: data.Username,
				}

				_ = json.Unmarshal(dataSes, &userInfo)
				fmt.Println("session data : ",userInfo)
				return userInfo
			},
		},
	})
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(db)

	adminGroup := e.Group("/admin",mv, middleware.SessionMiddleware(session.Manager))
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
		e.GET("/admin/v1/inventaris/update/:id", dashboardController.Update)
		e.GET("/admin/v1/inventaris/detail/:id", dashboardController.Detail)
		invenGroup.POST("/do-update/:id", dashboardController.DoUpdate)

	}


	m := g.Group("/master-data", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	provController := config.InjectMasterController(db)
	p := m.Group("/provinsi", authorizationMiddleware.AuthorizationMiddleware([]string{"1"}))
	p.GET("", provController.Index)
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

}