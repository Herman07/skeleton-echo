package repository

import (
	"gorm.io/gorm"
)

type UsersDataRepository interface {
	//Create(entity models.MasterDataProvinsi) (*models.MasterDataProvinsi, error)
	//UpdateById(entity models.MasterDataProvinsi)(*models.MasterDataProvinsi, error)
	//Delete(models.MasterDataProvinsi) error
	//FindById(id string) (*models.MasterDataProvinsi, error)
	DbInstance() *gorm.DB
}


type usersdataRepository struct {
	*gorm.DB
}

func NewUsersDataRepository(db *gorm.DB) UsersDataRepository {
	return &usersdataRepository{
		DB: db,
	}
}
//func (r masterdataRepository) Create(entity models.MasterDataProvinsi) (*models.MasterDataProvinsi, error) {
//	err := r.DB.Table("provinsi").Create(&entity).Error
//	return &entity, err
//}
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

func (r *usersdataRepository) DbInstance() *gorm.DB {
	return r.DB
}

