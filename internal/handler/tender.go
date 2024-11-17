package handler

import (
	"fmt"
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

func (h *Tender) Update(c *gin.Context) {
	fmt.Println(`
c.Param("id")`, c.Param("id"))

	id, err := strconv.Atoi(c.Param("id"))
	if HasErr(c, err) {
		return
	}

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

	data.ID = uint(id)
	data.ClientId = user.Id
	_, err = h.Service.Tenders.UpdateTender(data)

	if HasErr(c, err) {
		return
	}

	Success(c, gin.H{
		"message": "Tender status updated",
	}, 200)
}

func (h *Tender) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HasErr(c, err) {
		return
	}

	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}

	err = h.Service.Tenders.Delete(uint(id), user.Id)
	if HasErr(c, err) {
		return
	}

	Success(c, gin.H{"message": "Tender deleted successfully"}, 200)
}

func (h *Tender) GetListByClient(c *gin.Context) {
	data, err := bind[dtos.Tenders](c)
	if HasErr(c, err) {
		return
	}

	user, ok := h.MW.GetUser(c)
	if !ok {
		FailErr(c, app_errors.InternalServerError)
		return
	}

	data.ClientID = user.Id
	res, err := h.Service.Tenders.GetListTenders(data)
	if HasErr(c, err) {
		return
	}

	Success(c, res, 200)
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
