package handler

import (
	"github.com/alihaqberdi/goga_go/internal/handler/mw"
	"github.com/alihaqberdi/goga_go/internal/pkg/app_errors"
	"strconv"

	"github.com/alihaqberdi/goga_go/internal/dtos"
	. "github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/gin-gonic/gin"
)

type Tender struct {
	Service *service.Service
	MW      *mw.Middleware
}

func (h *Tender) Create(c *gin.Context) {
	data, err := bind[dtos.Tender](c)
	if err != nil {
		FailErr(c, app_errors.TenderInvalidInput)
		return
	}

	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}

	data.ClientId = user.Id
	res, err := h.Service.Tenders.CreateTender(data)

	if HasErr(c, err) {
		return
	}

	Success(c, res, 201)

}

func (h *Tender) UpdateTender(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("client_id"))
	if HasErr(c, err) {
		return
	}

	data, err := bind[dtos.Tender](c)
	if HasErr(c, err) {
		return
	}

	res, err := h.Service.Tenders.UpdateTender(userID, data)
	if HasErr(c, err) {
		return
	}

	Success(c, res, 201)
}

func (h *Tender) DeleteTender(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("client_id"))
	if HasErr(c, err) {
		return
	}

	tenderID, err := strconv.Atoi(c.Query("tender_id"))
	if HasErr(c, err) {
		return
	}

	err = h.Service.Tenders.DeleteTender(userID, tenderID)
	if HasErr(c, err) {
		return
	}

	Success(c, gin.H{"message": "Tender deleted successfully"}, 200)
}

func (h *Tender) GetListTenders(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if HasErr(c, err) {
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if HasErr(c, err) {
		return
	}

	res, err := h.Service.Tenders.Repo.Tenders.GetList(limit, offset)
	if HasErr(c, err) {
		return
	}

	Success(c, res, 201)
}

func (h *Tender) GetListTendersByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("client_id"))
	if HasErr(c, err) {
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if HasErr(c, err) {
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if HasErr(c, err) {
		return
	}

	res, err := h.Service.Tenders.GetListTendersByUser(userID, limit, offset)
	if HasErr(c, err) {
		return
	}

	Success(c, res, 200)
}
