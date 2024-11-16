package handler

import (
	"github.com/alihaqberdi/goga_go/internal/dtos"
	. "github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-shafaq/timep"
)

type Search struct {
	Service *service.Service
}

func (h *Search) SearchProbs(c *gin.Context) {
	data := new(dtos.SearchProbs)
	data.ProbId = c.Query("probId")
	data.Question = c.Query("question")

	res, err := h.Service.Search.SearchProbs(data)

	Finish(c, res, err)

}

func (h *Search) UpdateProbs(c *gin.Context) {
	param := c.Param("duration")
	dur, err := timep.ParseDuration(param)
	if HasErr(c, err) {
		return
	}

	go h.Service.Search.UpdateProbs(dur)

	Success(c, "going")

}
