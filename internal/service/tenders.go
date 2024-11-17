package service

import (
	"errors"
	"time"

	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"github.com/alihaqberdi/goga_go/internal/pkg/mapping"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type tenderService struct {
	Repo  *repo.Repo
	Cache *caching.Cache
}

func (s *tenderService) CreateTender(tender *dtos.Tender) (dtos.Tender, error) {
	// Validate the DTO
	if err := s.ValidateTender(tender); err != nil {
		return dtos.Tender{}, err
	}

	// Convert DTO to Model
	tenderModel := mapping.ConvertTenderDTOToModel(tender)

	// Call the repository to create the tender
	createdTenderModel, err := s.Repo.Tenders.Create(tenderModel)
	if err != nil {
		return dtos.Tender{}, err
	}

	createdTenderDTO := dtos.Tender{
		ID:          createdTenderModel.ID,
		ClientId:    createdTenderModel.ClientId,
		Title:       createdTenderModel.Title,
		Description: createdTenderModel.Description,
		Deadline:    createdTenderModel.Deadline,
		Budget:      createdTenderModel.Budget,
		Status:      createdTenderModel.Status,
	}

	return createdTenderDTO, nil
}

func (s *tenderService) UpdateTender(tender *dtos.Tender) error {
	if err := s.ValidateTender(tender); err != nil {
		return err
	}

	tenderModel := mapping.ConvertTenderDTOToModel(tender)
	err := s.Repo.Tenders.Update(tenderModel)

	return err
}

func (s *tenderService) GetListTenders(limit, offset int) (*dtos.Tender, error) {
	tenders, err := s.Repo.Tenders.GetList(limit, offset)
	if err != nil {
		return nil, err
	}

	var tenderDTOs = []dtos.Tender
}

func (s *tenderService) ValidateTender(tender *dtos.Tender) error {
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
