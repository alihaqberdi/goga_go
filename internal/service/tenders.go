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
	tender.Status = types.TenderStatusOpen

	if err := s.ValidateTender(tender); err != nil {
		return dtos.Tender{}, err
	}

	tenderModel := mapping.ConvertTenderDTOToModel(tender)

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

func (s *tenderService) UpdateTender(userID int, tender *dtos.Tender) (*dtos.Tender, error) {
	if err := s.ValidateTender(tender); err != nil {
		return nil, err
	}

	tenderModel := mapping.ConvertTenderDTOToModel(tender)
	err := s.Repo.Tenders.Update(userID, tenderModel)

	return nil, err
}

func (s *tenderService) DeleteTender(userID, tenderID int) error {
	// Call the repository to delete the tender
	err := s.Repo.Tenders.Delete(userID, tenderID)
	if err != nil {
		return err
	}

	return nil
}

func (s *tenderService) GetListTenders(limit, offset int) ([]dtos.Tender, error) {
	tenders, err := s.Repo.Tenders.GetList(limit, offset)
	if err != nil {
		return []dtos.Tender{}, err
	}

	tenderDTOs := make([]dtos.Tender, len(tenders))
	for i, model := range tenders {
		tenderDTOs[i] = dtos.Tender{
			ID:          model.ID,
			ClientId:    model.ClientId,
			Title:       model.Title,
			Description: model.Description,
			Deadline:    model.Deadline,
			Budget:      model.Budget,
			Status:      model.Status,
		}
	}

	return tenderDTOs, nil
}

func (s *tenderService) ValidateTender(tender *dtos.Tender) error {
	if tender.Budget <= 0 {
		return errors.New("budget must be greater than zero")
	}

	if tender.Deadline.Before(time.Now()) {
		return errors.New("deadline must be in the future")
	}

	if tender.Status != types.TenderStatusOpen && tender.Status != types.TenderStatusClosed {
		return errors.New("invalid status, must be either 'open' or 'closed'")
	}

	return nil
}

func (s *tenderService) GetListTendersByUser(userID, limit, offset int) ([]dtos.Tender, error) {
	tenders, err := s.Repo.Tenders.GetListByUser(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	tenderDTOs := make([]dtos.Tender, len(tenders))
	for i, model := range tenders {
		tenderDTOs[i] = dtos.Tender{
			ID:          model.ID,
			ClientId:    model.ClientId,
			Title:       model.Title,
			Description: model.Description,
			Deadline:    model.Deadline,
			Budget:      model.Budget,
			Status:      model.Status,
		}
	}

	return tenderDTOs, nil
}
