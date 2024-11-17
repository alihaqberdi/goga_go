package handler

import (
	"net/http"

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

func (h *Tender) CreateTenderHandler(c *gin.Context) {
	var tenderDTO dtos.Tender

	if err := c.ShouldBindJSON(&tenderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input: " + err.Error()})
		return
	}

	createdTender, err := h.Service.Tenders.CreateTender(&tenderDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create tender: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTender)
}

func (h *Tender) UpdateTenderHandler(c *gin.Context) {
	var tenderDTO dtos.Tender

	if err := c.ShouldBindJSON(&tenderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input: " + err.Error()})
		return
	}

	if err := h.Service.Tenders.UpdateTender(&tenderDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update tender: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tender updated successfully"})
}

// func (h *Tender) GetListTendersHandler(c *gin.Context) {
// 	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
// 	if err != nil || limit <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
// 		return
// 	}

// 	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
// 	if err != nil || offset < 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
// 		return
// 	}

// 	tenders, err := h.Service.Tenders.Repo.Tenders.GetList(limit, offset)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve tenders: " + err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, tenders)
// }
