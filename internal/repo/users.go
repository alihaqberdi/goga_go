package repo

import (
	"github.com/alihaqberdi/goga_go/internal/models"
	"gorm.io/gorm"
)

type Users struct {
	DB *gorm.DB
}

func (r *Users) Create(m *models.User) error {
	err := r.DB.Create(m).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Users) GetByID(id uint) (*models.Tender, error) {
	var m models.Tender
	if err := r.DB.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *Users) GetList(limit, offset int) ([]models.Tender, error) {
	var list []models.Tender
	query := r.DB.Limit(limit).Offset(offset)

	err := query.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Users) Update(tender *models.Tender) error {
	return r.DB.Save(tender).Error
}

func (r *Users) Delete(id uint) error {
	return r.DB.Delete(&models.Tender{}, id).Error
}
