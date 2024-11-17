package handler

import (
	"strconv"

	"github.com/alihaqberdi/goga_go/internal/dtos"
	. "github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/gin-gonic/gin"
)

type Bids struct {
	Service *service.Service
}

// CreateBid godoc
// @Summary Create a new bid
// @Description Create a new bid
// @Tags Bids
// @Accept json
// @Produce json
// @Param tender_id path int true "Tender ID"
// @Param bid body dtos.BidsCreate true "Bid object"
// @Success 200 {object} dtos.BidList
// @Router /api/v1/bids/{tender_id} [post]
func (h *Bids) CreateBid(c *gin.Context) {
	data, err := bind[dtos.BidsCreate](c)
	if HasErr(c, err) {
		return
	}
	res, err := h.Service.Bids.CreateBid(data)
	if HasErr(c, err) {
		return
	}
	Success(c, res)
}

// GetList godoc
// @Summary Get list of bids
// @Description Get list of bids
// @Tags Bids
// @Accept json
// @Produce json
// @Param tender_id path int true "Tender ID"
// @Success 200 {object} dtos.BidList
// @Router /api/v1/bids/{tender_id} [get]
func (h *Bids) GetList(c *gin.Context) {
	tenderIdStr := c.Param("tender_id")
	tenderId, err := strconv.ParseUint(tenderIdStr, 10, 32)
	if HasErr(c, err) {
		return
	}
	res, err := h.Service.Bids.GetList(uint(tenderId))
	if HasErr(c, err) {
		return
	}
	Success(c, res)
}
