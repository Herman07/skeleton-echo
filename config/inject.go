package config

import (
	"gorm.io/gorm"
	"skeleton-echo/controllers"
	"skeleton-echo/repository"
	"skeleton-echo/services"
)

func InjectAuthController(db *gorm.DB) controllers.FrontAuthController {
	authRepository := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)
	return authController
}
func InjectDashboardController(db *gorm.DB) controllers.P3Controller {
	dashboardRepository := repository.NewP3Repository(db)
	dashboardService := services.NewP3Service(dashboardRepository)
	dashboardController := controllers.NewP3Controller(dashboardService)
	return dashboardController
}

func InjectMasterController(db *gorm.DB) controllers.ProvDataController {
	masterRepository := repository.NewProvDataRepository(db)
	masterService := services.NewProvDataService(masterRepository)
	masterController := controllers.NewProvDataController(masterService)
	return masterController
}

func InjectKabController(db *gorm.DB) controllers.KabDataController {
	kabRepository := repository.NewKabDataRepository(db)
	kabService := services.NewKabDataService(kabRepository)
	kabController := controllers.NewKabDataController(kabService)
	return kabController
}
func InjectKecController(db *gorm.DB) controllers.KecDataController {
	kecRepository := repository.NewKecDataRepository(db)
	kecService := services.NewKecDataService(kecRepository)
	kecController := controllers.NewKecDataController(kecService)
	return kecController
}