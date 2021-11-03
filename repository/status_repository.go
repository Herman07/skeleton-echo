package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type StatusDataRepository interface {
	Create(entity models.StatusLegal) (*models.StatusLegal, error)
	UpdateById(entity models.StatusLegal)(*models.StatusLegal, error)
	Delete(models.StatusLegal) error
	FindById(id string) (*models.StatusLegal, error)
	DbInstance() *gorm.DB
}


type statusdataRepository struct {
	*gorm.DB
}

func NewStatusDataRepository(db *gorm.DB) StatusDataRepository {
	return &statusdataRepository{
		DB: db,
	}
}

func (r statusdataRepository) Create(entity models.StatusLegal) (*models.StatusLegal, error) {
	err := r.DB.Table("status_legal").Create(&entity).Error
	return &entity, err
}

func (r statusdataRepository) UpdateById(entity models.StatusLegal)(*models.StatusLegal, error){
	err := r.DB.Model(&models.StatusLegal{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r statusdataRepository) FindById(id string) (*models.StatusLegal, error) {
	var entity models.StatusLegal
	err := r.DB.Table("status_legal").Where("id_status_legal = ?", id).First(&entity).Error
	return &entity, err
}

func (r statusdataRepository) Delete(entity models.StatusLegal) error {
	return r.DB.Table("status_legal").Delete(&entity).Error
}

func (r *statusdataRepository) DbInstance() *gorm.DB {
	return r.DB
}

