package handler

import (
	"strconv"

	"github.com/alihaqberdi/goga_go/internal/dtos"
	. "github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/gin-gonic/gin"
)

type Tender struct {
	Service *service.Service
}

func (h *Tender) CreateTender(c *gin.Context) {
	data, err := bind[dtos.Tender](c)
	if HasErr(c, err) {
		return
	}

	res, err := h.Service.Tenders.CreateTender(data)

	if HasErr(c, err) {
		return
	}

	Success(c, res, 201)

}

func (h *Tender) UpdateTender(c *gin.Context) {
	data, err := bind[dtos.Tender](c)
	if HasErr(c, err) {
		return
	}

	res, err := h.Service.Tenders.UpdateTender(data)

	if HasErr(c, err) {
		return
	}

	Success(c, res, 201)
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
