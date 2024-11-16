package repo

import (
	"github.com/alihaqberdi/goga_go/internal/models"
	"gorm.io/gorm"
)

type TenderRepository interface {
	Create(tender *models.Tender) (uint, error)
	GetByID(id uint) (*models.Tender, error)
	GetAll() ([]models.Tender, error)
	Update(tender *models.Tender) error
	Delete(id uint) error
}

type tenderRepository struct {
	db *gorm.DB
}

func NewTenderRepository(db *gorm.DB) TenderRepository {
	return &tenderRepository{db: db}
}

func (r *tenderRepository) Create(tender *models.Tender) (uint, error) {
	if err := r.db.Create(tender).Error; err != nil {
		return 0, err
	}
	return tender.ID, nil
}

func (r *tenderRepository) GetByID(id uint) (*models.Tender, error) {
	var tender models.Tender
	if err := r.db.First(&tender, id).Error; err != nil {
		return nil, err
	}
	return &tender, nil
}

func (r *tenderRepository) GetAll() ([]models.Tender, error) {
	var tenders []models.Tender
	if err := r.db.Find(&tenders).Error; err != nil {
		return nil, err
	}
	return tenders, nil
}

func (r *tenderRepository) Update(tender *models.Tender) error {
	return r.db.Save(tender).Error
}

func (r *tenderRepository) Delete(id uint) error {
	return r.db.Delete(&models.Tender{}, id).Error
}
