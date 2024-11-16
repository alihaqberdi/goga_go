package service

import (
	"errors"

	"github.com/alihaqberdi/goga_go/internal/models"
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type bidsService struct {
	Repo  *repo.Repo
	Cache *caching.Cache
}

func (s *bidsService) CreateBid(bid *models.Bid) error {
	err := s.ValidateBid(bid)
	if err != nil {
		return err
	}
}
func (s *bidsService) ValidateBid(bid *models.Bid) error {
	if bid.Price <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if bid.Status != types.BidStatusPending {
		return errors.New("invalid status, must be 'pending'")
	}
	tender, err := s.Repo.Tenders.GetByID(bid.TenderId)
	if err != nil {
		return err
	}
	if tender.Status != types.TenderStatusOpen {
		return errors.New("tender is not open")
	}
	// You can add more validation rules as needed like rate limiting, etc.
	return nil
}
