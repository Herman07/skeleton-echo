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
func InjectStatusController(db *gorm.DB) controllers.StatusDataController {
	statusRepository := repository.NewStatusDataRepository(db)
	statusService := services.NewStatusDataService(statusRepository)
	statusController := controllers.NewStatusDataController(statusService)
	return statusController
}
func InjectPengurusController(db *gorm.DB) controllers.PengurusDataController {
	pengurusRepository := repository.NewPengurusDataRepository(db)
	pengurusService := services.NewPengurusDataService(pengurusRepository)
	pengurusController := controllers.NewPengurusDataController(pengurusService)
	return pengurusController
}
func InjectIrigasiController(db *gorm.DB) controllers.IrigasiDataController {
	irigasiRepository := repository.NewIrigasiDataRepository(db)
	irigasiService := services.NewIrigasiDataService(irigasiRepository)
	irigasiController := controllers.NewIrigasiDataController(irigasiService)
	return irigasiController
}
func InjectPertanianController(db *gorm.DB) controllers.TaniDataController {
	taniRepository := repository.NewTaniDataRepository(db)
	taniService := services.NewTaniDataService(taniRepository)
	taniController := controllers.NewTaniDataController(taniService)
	return taniController
}