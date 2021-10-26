package repository

import (
	"fmt"
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type IrigasiDataRepository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.TeknikIrigasi, error)
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
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
func (r *irigasidataRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.TeknikIrigasi, error) {
	var entity []models.TeknikIrigasi
	res := r.DB.Table("teknik_irigasi").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

	for k, v := range keyVal {
		switch operation {
		case "and":
			res = res.Where(k, v)
		case "or":
			res = res.Or(k, v)
		}
	}
	err := res.Find(&entity).Error
	return entity, err

}
func (r irigasidataRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("teknik_irigasi").Count(&count).Error
	return count, err
}

func (r irigasidataRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.TeknikIrigasi{})
	for k, v := range keyVal {
		switch operation {
		case "and":
			q = q.Where(k, v)
		case "or":
			q = q.Or(k, v)
		}
	}

	err := q.Count(&count).Error
	return count, err
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

