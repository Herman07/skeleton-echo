package repository

import (
	"fmt"
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type IrigasiDataRepository interface {
	Create(entity models.TeknikIrigasi) (*models.TeknikIrigasi, error)
	UpdateById(models.TeknikIrigasi)(*models.TeknikIrigasi, error)
	Delete(irigasi models.TeknikIrigasi) error
	FindById(id string) (*models.TeknikIrigasi, error)
	DbInstance() *gorm.DB
}


type irigasidataRepository struct {
	*gorm.DB
}

func NewIrigasiDataRepository(db *gorm.DB) IrigasiDataRepository {
	return &irigasidataRepository{
		DB: db,
	}
}
func (r irigasidataRepository) Create(entity models.TeknikIrigasi) (*models.TeknikIrigasi, error) {
	err := r.DB.Table("teknik_irigasi").Create(&entity).Error
	return &entity, err
}

func (r irigasidataRepository) UpdateById(entity models.TeknikIrigasi)(*models.TeknikIrigasi, error){
	fmt.Println("Data Repo", entity)
	err := r.DB.Model(&models.TeknikIrigasi{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r irigasidataRepository) FindById(id string) (*models.TeknikIrigasi, error) {
	var entity models.TeknikIrigasi
	err := r.DB.Table("teknik_irigasi").Where("id_t_irigasi = ?", id).First(&entity).Error
	return &entity, err
}

func (r irigasidataRepository) Delete(entity models.TeknikIrigasi) error {
	return r.DB.Delete(&entity).Error
}

func (r *irigasidataRepository) DbInstance() *gorm.DB {
	return r.DB
}

