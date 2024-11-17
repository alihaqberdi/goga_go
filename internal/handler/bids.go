package handler

import (
	"fmt"
	"strconv"

	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/alihaqberdi/goga_go/internal/handler/mw"
	. "github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/pkg/app_errors"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/gin-gonic/gin"
)

type Bids struct {
	Service *service.Service
	MW      *mw.Middleware
}

// CreateBid godoc
// @Summary Create a new bid
// @Description Create a new bid
// @Tags Bids
// @Accept json
// @Produce json
// @Param tender_id path int true "Tenders ID"
// @Param bid body dtos.BidCreate true "Bid object"
// @Success 200 {object} dtos.BidList
// @Router /api/contractor/tenders/{tender_id}/bid [post]
func (h *Bids) Create(c *gin.Context) {
	tenderId, err := strconv.ParseUint(c.Param("tender_id"), 10, 32)

	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}

	data, err := bind[dtos.BidCreate](c)
	if HasErr(c, err) {
		return
	}

	data.ContractorID = user.Id
	data.Status = "pending"
	data.TenderID = uint(tenderId)

	fmt.Println(data)
	res, err := h.Service.Bids.Create(data)
	if HasErr(c, err) {
		return
	}

	Success(c, res, 201)

}

// GetList godoc
// @Summary Get list of bids
// @Description Get list of bids
// @Tags Bids
// @Accept json
// @Produce json
// @Param tender_id path int true "Tenders ID"
// @Success 200 {object} dtos.BidList
// @Router /api/client/tenders/{tender_id}/bids [get]
func (h *Bids) GetList(c *gin.Context) {
	tenderIdStr := c.Param("id")
	tenderId, err := strconv.ParseUint(tenderIdStr, 10, 32)
	if HasErr(c, err) {
		return
	}

	_ = tenderId
	res, err := h.Service.Bids.GetList(&dtos.Bids{
		TenderID: uint(tenderId),
	})
	if HasErr(c, err) {
		return
	}

	Success(c, res)

}

func (h *Bids) GetListByContractor(c *gin.Context) {
	data, err := bind[dtos.Bids](c)
	if HasErr(c, err) {
		return
	}

	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}

	data.ContractorID = user.Id
	res, err := h.Service.Bids.GetList(data)
	if HasErr(c, err) {
		return
	}

	Success(c, res, 200)
}

// AwardBid godoc
// @Summary Award a bid
// @Description Award a bid
// @Tags Bids
// @Accept json
// @Produce json
// @Param tender_id path int true "Tenders ID"
// @Param id path int true "Bid ID"
// @Success 200
// @Router /api/client/tenders/{tender_id}/award/{id} [post]
func (h *Bids) AwardBid(c *gin.Context) {
	tenderIdStr := c.Param("id")
	tenderId, err := strconv.ParseUint(tenderIdStr, 10, 32)
	if HasErr(c, err) {
		return
	}
	idStr := c.Param("bid_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if HasErr(c, err) {
		return
	}

	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}

	err = h.Service.Bids.AwardBid(uint(tenderId), uint(id), user.Id)
	if HasErr(c, err) {
		return
	}

	Success(c, gin.H{
		"message": "Bid awarded successfully",
	})

}

// Delete godoc
// @Summary Delete a bid
// @Description Delete a bid
// @Tags Bids
// @Accept json
// @Produce json
// @Param id path int true "Bid ID"
// @Success 200
// @Router /api/contractor/bids/{tender_id}/bid/{id} [delete]
func (h *Bids) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if HasErr(c, err) {
		return
	}

	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}

	err = h.Service.Bids.Delete(uint(id), user.Id)
	if HasErr(c, err) {
		return
	}

	Success(c, gin.H{"message": "Bid deleted successfully"})

}

// UserBids godoc
// @Summary Get user bids
// @Description Get user bids
// @Tags Bids
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dtos.BidList
// @Router /users/{id}/bids [get]
func (h *Bids) UserBids(c *gin.Context) {
	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}
	userId := user.Id
	res, err := h.Service.Bids.UserBids(uint(userId))
	if HasErr(c, err) {
		return
	}
	Success(c, res)
}
