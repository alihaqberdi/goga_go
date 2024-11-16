package service

import (
	"errors"
	"time"

	"github.com/alihaqberdi/goga_go/internal/models"
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type tenderService struct {
	Repo  *repo.Repo
	Cache *caching.Cache
}

func (s *tenderService) CreateTender(tender *models.Tender) (uint, error) {
	if err := s.ValidateTender(tender); err != nil {
		return 0, err
	}

	return s.Repo.Tenders.Create(tender)
}

func (s *tenderService) UpdateTender(tender *models.Tender) error {
	if err := s.ValidateTender(tender); err != nil {
		return err
	}

	return s.Repo.Tenders.Update(tender)
}

func (s *tenderService) GetListTenders(limit, offset int) ([]models.Tender, error) {
	return s.Repo.Tenders.GetList(limit, offset)
}

func (s *tenderService) ValidateTender(tender *models.Tender) error {
	// Ensure the budget is greater than 0
	if tender.Budget <= 0 {
		return errors.New("budget must be greater than zero")
	}

	// Ensure the deadline is in the future
	if tender.Deadline.Before(time.Now()) {
		return errors.New("deadline must be in the future")
	}

	// Ensure that the status is valid (you can expand this based on your business rules)
	if tender.Status != types.TenderStatusOpen && tender.Status != types.TenderStatusClosed {
		return errors.New("invalid status, must be either 'open' or 'closed'")
	}

	// You can add more validation rules as needed

	return nil
}
