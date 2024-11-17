package repo

import (
	"github.com/alihaqberdi/goga_go/internal/models"
	"gorm.io/gorm"
)

type Tenders struct {
	db *gorm.DB
}

func (r *Tenders) Create(tender *models.Tender) (models.Tender, error) {
	if err := r.db.Create(tender).Error; err != nil {
		return err
	}
	return nil
}

func (r *Tenders) GetByID(id uint) (*models.Tender, error) {
	var tender models.Tender
	if err := r.db.First(&tender, id).Error; err != nil {
		return nil, err
	}
	return &tender, nil
}

func (r *Tenders) GetList(limit, offset int) ([]models.Tender, error) {
	var tenders []models.Tender
	query := r.db.Limit(limit).Offset(offset)

	if err := query.Find(&tenders).Error; err != nil {
		return nil, err
	}
	return tenders, nil
}

func (r *Tenders) Update(tender *models.Tender) error {
	return r.db.Save(tender).Error
}

func (r *Tenders) Delete(id uint) error {
	return r.db.Delete(&models.Tender{}, id).Error
}
