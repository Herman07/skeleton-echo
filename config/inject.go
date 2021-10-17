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
func InjectDashboardController(db *gorm.DB) controllers.DashboardController {
	dashboardRepository := repository.NewDashboardRepository(db)
	dashboardService := services.NewDashboardService(dashboardRepository)
	dashboardController := controllers.NewDashboardController(dashboardService)
	return dashboardController
}

func InjectMasterController(db *gorm.DB) controllers.MasterDataController {
	masterRepository := repository.NewMasterDataRepository(db)
	masterService := services.NewMasterDataService(masterRepository)
	masterController := controllers.NewMasterDataController(masterService)
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