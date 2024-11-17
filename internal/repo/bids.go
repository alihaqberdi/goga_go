package repo

import (
	"github.com/alihaqberdi/goga_go/internal/dtos"
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

func (r *Bids) GetList(data *dtos.Bids) ([]models.Bid, error) {
	var bids []models.Bid
	query := r.DB

	if data.ContractorID > 0 && data.TenderID > 0 {
		query = query.Where("contractor_id = ? AND tender_id=?", data.ContractorID, data.TenderID)
	} else if data.ContractorID > 0 {
		query = query.Where("contractor_id=?", data.ContractorID)
	} else if data.TenderID > 0 {
		query = query.Where("tender_id=?", data.TenderID)
	}

	err := query.Limit(data.Limit).Offset(data.Offset).Find(&bids).Error

	if err != nil {
		return nil, err
	}

	return bids, nil
}

func (r *Bids) GetByID(id uint) (*models.Bid, error) {
	var bid models.Bid
	if err := r.DB.First(&bid, id).Error; err != nil {
		return nil, err
	}

	return &bid, nil
}

func (r *Bids) AwardBid(id uint) error {
	if err := r.DB.Model(&models.Bid{}).Where("id = ?", id).Update("status", types.BidStatusAwarded).Error; err != nil {
		return err
	}

	return nil
}

func (r *Bids) Delete(id uint) error {
	if err := r.DB.Delete(&models.Bid{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *Bids) UserBids(userID uint) ([]models.Bid, error) {
	var bids []models.Bid
	if err := r.DB.Where("contractor_id = ?", userID).Find(&bids).Error; err != nil {
		return nil, err
	}

	return bids, nil
}
