package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type TaniDataRepository interface {
	Create(entity models.TeknikPertanian) (*models.TeknikPertanian, error)
	UpdateById(entity models.TeknikPertanian)(*models.TeknikPertanian, error)
	Delete(pertanian models.TeknikPertanian) error
	FindById(id string) (*models.TeknikPertanian, error)
	DbInstance() *gorm.DB
}


type tanidataRepository struct {
	*gorm.DB
}

func NewTaniDataRepository(db *gorm.DB) TaniDataRepository {
	return &tanidataRepository{
		DB: db,
	}
}

func (r tanidataRepository) Create(entity models.TeknikPertanian) (*models.TeknikPertanian, error) {
	err := r.DB.Table("teknik_pertanian").Create(&entity).Error
	return &entity, err
}

func (r tanidataRepository) UpdateById(entity models.TeknikPertanian)(*models.TeknikPertanian, error){
	err := r.DB.Model(&models.TeknikPertanian{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r tanidataRepository) FindById(id string) (*models.TeknikPertanian, error) {
	var entity models.TeknikPertanian
	err := r.DB.Table("teknik_pertanian").Where("id_t_pertanian = ?", id).First(&entity).Error
	return &entity, err
}

func (r tanidataRepository) Delete(entity models.TeknikPertanian) error {
	return r.DB.Table("teknik_pertanian").Delete(&entity).Error
}

func (r *tanidataRepository) DbInstance() *gorm.DB {
	return r.DB
}

