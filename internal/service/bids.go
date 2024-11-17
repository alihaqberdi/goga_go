package service

import (
	"errors"

	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/alihaqberdi/goga_go/internal/models"
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"github.com/alihaqberdi/goga_go/internal/pkg/app_errors"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type bidsService struct {
	Repo  *repo.Repo
	Cache *caching.Cache
}

func (s *bidsService) CreateBid(bid *dtos.BidCreate) (*models.Bid, error) {
	err := s.ValidateBid(bid)
	if err != nil {
		return nil, err
	}
	return s.Repo.Bids.Create(&models.Bid{
		TenderId:     bid.TenderID,
		Price:        bid.Price,
		Status:       bid.Status,
		ContractorId: bid.ContractorID,
		DeliveryTime: bid.DeliveryTime,
		Comments:     bid.Comments,
	})

}
func (s *bidsService) Delete(id uint) error {
	_, err := s.Repo.Bids.GetByID(id)
	if err != nil {
		return app_errors.BidNotFound
	}
	return s.Repo.Bids.Delete(id)
}
func (s *bidsService) UserBids(userID uint) ([]models.Bid, error) {
	return s.Repo.Bids.UserBids(userID)
}
func (s *bidsService) AwardBid(tenderID, id uint) error {
	tender, err := s.Repo.Tenders.GetByID(tenderID)
	bid, err := s.Repo.Bids.GetByID(id)
	if err != nil {
		return app_errors.TenderNotFound
	}
	if bid.Status != types.BidStatusPending {
		return app_errors.BidNotPending
	}
	if tender.Status != types.TenderStatusClosed {
		return app_errors.TenderNotClosed
	}

	return s.Repo.Bids.AwardBid(id)
}
func (s *bidsService) GetList(tenderID uint) ([]models.Bid, error) {
	_, err := s.Repo.Tenders.GetByID(tenderID)
	if err != nil {
		return nil, app_errors.TenderNotFound
	}
	return s.Repo.Bids.GetList(tenderID)
}

func (s *bidsService) ValidateBid(bid *dtos.BidCreate) error {
	if bid.Price <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if bid.Status != types.BidStatusPending {
		return errors.New("invalid status, must be 'pending'")
	}
	tender, err := s.Repo.Tenders.GetByID(bid.TenderID)
	if err != nil {
		return err
	}
	if tender.Status != types.TenderStatusOpen {
		return errors.New("tender is not open")
	}
	// You can add more validation rules as needed like rate limiting, etc.
	return nil
}
