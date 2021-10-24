package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type StatusDataRepository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.StatusLegal, error)
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	//Create(entity models.MasterDataProvinsi) (*models.MasterDataProvinsi, error)
	//UpdateById(entity models.MasterDataProvinsi)(*models.MasterDataProvinsi, error)
	//Delete(models.MasterDataProvinsi) error
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
func (r *statusdataRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.StatusLegal, error) {
	var entity []models.StatusLegal
	res := r.DB.Table("status_legal").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

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
func (r statusdataRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("status_legal").Count(&count).Error
	return count, err
}

func (r statusdataRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.StatusLegal{})
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
//
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
func (r statusdataRepository) FindById(id string) (*models.StatusLegal, error) {
	var entity models.StatusLegal
	err := r.DB.Table("status_legal").Where("id_status_legal = ?", id).First(&entity).Error
	return &entity, err
}
//
//func (r masterdataRepository) Delete(entity models.MasterDataProvinsi) error {
//	return r.DB.Table("provinsi").Delete(&entity).Error
//}

func (r *statusdataRepository) DbInstance() *gorm.DB {
	return r.DB
}

