package repo

import (
	"github.com/alihaqberdi/goga_go/internal/models"
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"gorm.io/gorm"
)

type Bids struct {
	DB *gorm.DB
}

func (r *Bids) Create(bid *models.Bid) (*models.Bid, error) {
	if err := r.DB.Create(bid).Error; err != nil {
		return nil, err
	}

	return bid, nil
}

func (r *Bids) GetList(tenderId uint) ([]models.Bid, error) {
	var bids []models.Bid
	if err := r.DB.Where("tender_id = ?", tenderId).Find(&bids).Error; err != nil {
		return nil, err
	}

	return bids, nil
}

func (r *Bids) AwardBid(id uint) error {
	if err := r.DB.Model(&models.Bid{}).Where("id = ?", id).Update("status", types.BidStatusAwarded).Error; err != nil {
		return err
	}

	return nil
}
