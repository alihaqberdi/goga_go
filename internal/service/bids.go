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

type Bids struct {
	Repo  *repo.Repo
	Cache *caching.Cache
}

func (s *Bids) Create(bid *dtos.BidCreate) (*dtos.BidList, error) {
	err := s.validateBid(bid)
	if err != nil {
		return nil, err
	}

	data := models.Bid{
		TenderId:     bid.TenderID,
		ContractorId: bid.ContractorID,
		Price:        bid.Price,
		DeliveryTime: bid.DeliveryTime,
		Comments:     bid.Comments,
		Status:       bid.Status,
	}
	res, err := s.Repo.Bids.Create(&data)
	if err != nil {
		return nil, err
	}

	return &dtos.BidList{
		BidsBase: dtos.BidsBase{
			TenderID:     res.TenderId,
			ContractorID: res.ContractorId,
			Price:        res.Price,
			DeliveryTime: res.DeliveryTime,
			Comments:     res.Comments,
			Status:       res.Status,
		},
		ID: res.ID,
	}, nil
}

func (s *Bids) Delete(id uint, contractorId uint) error {
	bid, err := s.Repo.Bids.GetByID(id)
	if err != nil || bid.ContractorId != contractorId {
		return app_errors.BidNotFoundOrAccessDenied
	}

	return s.Repo.Bids.Delete(id)
}

func (s *Bids) UserBids(userID uint) ([]models.Bid, error) {
	return s.Repo.Bids.UserBids(userID)
}

func (s *Bids) AwardBid(tenderID, id, clientID uint) error {
	tender, err := s.Repo.Tenders.GetByID(tenderID)
	_ = tender
	if err != nil || tender.ClientId != clientID {
		return app_errors.TenderNotFoundOrAccessDenied
	}

	bid, err := s.Repo.Bids.GetByID(id)
	if err != nil {
		return app_errors.BidNotFound
	}

	if bid.Status != types.BidStatusPending {
		return app_errors.BidNotPending
	}
	//if tender.Status != types.TenderStatusClosed {
	//	return app_errors.TenderNotClosed
	//}

	return s.Repo.Bids.AwardBid(id)
}

func (s *Bids) GetList(data *dtos.Bids) ([]dtos.BidList, error) {
	if data.Limit == 0 {
		data.Limit = 10
	}

	if data.TenderID > 0 {
		_, err := s.Repo.Tenders.GetByID(data.TenderID)
		if err != nil {
			return nil, app_errors.TenderNotFoundOrAccessDenied
		}
	}

	list, err := s.Repo.Bids.GetList(data)
	if err != nil {
		return nil, err
	}

	dtoList := make([]dtos.BidList, len(list))
	for i, bid := range list {
		dtoList[i] = *s.mapper(&bid)
	}

	return dtoList, nil
}

func (s *Bids) validateBid(bid *dtos.BidCreate) error {

	if bid.Price <= 0 {
		return app_errors.BidInvalidData
	}

	if bid.Status != types.BidStatusPending {
		return errors.New("invalid status, must be 'pending'")
	}
	tender, err := s.Repo.Tenders.GetByID(bid.TenderID)

	if err != nil {
		return app_errors.TenderNotFound
	}

	if tender.Status != types.TenderStatusOpen {
		return app_errors.BidTenderIsNotOpen
	}
	// You can add more validation rules as needed like rate limiting, etc.
	return nil
}

func (s *Bids) mapper(m *models.Bid) *dtos.BidList {
	return &dtos.BidList{
		BidsBase: dtos.BidsBase{
			TenderID:     m.TenderId,
			ContractorID: m.ContractorId,
			Price:        m.Price,
			DeliveryTime: m.DeliveryTime,
			Comments:     m.Comments,
			Status:       m.Status,
		},
		ID: m.ID,
	}
}
