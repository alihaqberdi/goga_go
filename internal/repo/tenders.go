package repo

import (
	"errors"

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

func (r *Tenders) GetList(limit, offset int) ([]models.Tender, error) {
	var tenders []models.Tender
	query := r.db.Limit(limit).Offset(offset)

	if err := query.Find(&tenders).Error; err != nil {
		return nil, err
	}
	return tenders, nil
}

func (r *Tenders) Update(userID int, tender *models.Tender) error {
	var existingTender models.Tender
	if err := r.db.Where("id = ? AND client_id = ?", tender.ID, userID).First(&existingTender).Error; err != nil {
		return errors.New("tender not found or you do not have permission to update it")
	}

	return r.db.Model(&existingTender).Updates(tender).Error
}

func (r *Tenders) Delete(userID, tenderID int) error {
	var tender models.Tender
	if err := r.db.Where("id = ? AND client_id = ?", tenderID, userID).First(&tender).Error; err != nil {
		return errors.New("tender not found or you do not have permission to delete it")
	}

	return r.db.Delete(&tender).Error
}

func (r *Tenders) GetListByUser(userID int, limit, offset int) ([]models.Tender, error) {
	var tenders []models.Tender
	err := r.db.Where("ClientId = ?", userID).Limit(limit).Offset(offset).Find(&tenders).Error
	if err != nil {
		return nil, err
	}
	return tenders, nil
}
