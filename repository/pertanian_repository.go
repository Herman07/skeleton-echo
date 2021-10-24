package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type TaniDataRepository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.TeknikPertanian, error)
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
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
func (r *tanidataRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.TeknikPertanian, error) {
	var entity []models.TeknikPertanian
	res := r.DB.Table("teknik_pertanian").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

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
func (r tanidataRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("teknik_pertanian").Count(&count).Error
	return count, err
}

func (r tanidataRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.TeknikPertanian{})
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

