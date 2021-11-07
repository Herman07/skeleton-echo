package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type PengurusDataRepository interface {
	Create(entity models.Pengurus) (*models.Pengurus, error)
	//UpdateById(entity models.MasterDataProvinsi)(*models.MasterDataProvinsi, error)
	//Delete(models.MasterDataProvinsi) error
	//FindById(id string) (*models.MasterDataProvinsi, error)
	DbInstance() *gorm.DB
}


type pengurusdataRepository struct {
	*gorm.DB
}

func NewPengurusDataRepository(db *gorm.DB) PengurusDataRepository {
	return &pengurusdataRepository{
		DB: db,
	}
}

//
func (r pengurusdataRepository) Create(entity models.Pengurus) (*models.Pengurus, error) {
	err := r.DB.Table("provinsi").Create(&entity).Error
	return &entity, err
}
//
//func (r masterdataRepository) UpdateById(entity models.MasterDataProvinsi)(*models.MasterDataProvinsi, error){
//	err := r.DB.Model(&models.MasterDataProvinsi{ID: entity.ID}).Updates(&entity).Error
//	return &entity, err
//}
//
//func (r masterdataRepository) FindById(id string) (*models.MasterDataProvinsi, error) {
//	var entity models.MasterDataProvinsi
//	err := r.DB.Table("provinsi").Where("id_prov = ?", id).First(&entity).Error
//	return &entity, err
//}
//
//func (r masterdataRepository) Delete(entity models.MasterDataProvinsi) error {
//	return r.DB.Table("provinsi").Delete(&entity).Error
//}

func (r *pengurusdataRepository) DbInstance() *gorm.DB {
	return r.DB
}

