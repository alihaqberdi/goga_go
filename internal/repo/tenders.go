package repo

import (
	"github.com/alihaqberdi/goga_go/internal/dtos"

	"github.com/alihaqberdi/goga_go/internal/models"
	"gorm.io/gorm"
)

type Tenders struct {
	db *gorm.DB
}

func (r *Tenders) Create(tender *models.Tender) (models.Tender, error) {
	if err := r.db.Create(tender).Error; err != nil {
		return models.Tender{}, err
	}
	return *tender, nil
}

func (r *Tenders) GetByID(id uint) (*models.Tender, error) {
	var tender models.Tender
	if err := r.db.First(&tender, id).Error; err != nil {
		return nil, err
	}
	return &tender, nil
}

func (r *Tenders) GetList(data *dtos.Tenders) ([]models.Tender, error) {
	var tenders []models.Tender
	query := r.db.Limit(data.Limit).Offset(data.Offset)

	if data.ClientID > 0 {
		query = query.Where("client_id = ?", data.ClientID)
	}

	err := query.Find(&tenders).Error
	if err != nil {
		return nil, err
	}

	return tenders, nil
}

func (r *Tenders) Update(tender *models.Tender) error {
	return r.db.Model(&tender).Where("id=?", tender.ID).Updates(tender).Error
}

func (r *Tenders) Delete(id uint) error {
	return r.db.
		Where("id = ?", id).
		Delete(&models.Tender{}).Error
}

func (r *Tenders) GetListByUser(userID int, limit, offset int) ([]models.Tender, error) {
	var tenders []models.Tender
	err := r.db.Where("ClientId = ?", userID).Limit(limit).Offset(offset).Find(&tenders).Error
	if err != nil {
		return nil, err
	}
	return tenders, nil
}
