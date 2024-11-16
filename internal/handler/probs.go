package handler

import (
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/gin-gonic/gin"
)

type Probs struct {
	Service *service.Service
}

func (h *Probs) Save(c *gin.Context) {

}
func (h *Probs) LookupProb(c *gin.Context) {}

//func (h *Probs) Save(c *gin.Context) {
//
//	data, err := bind[dtos.SaveProb](c)
//	if HasErr(c, err) {
//		return
//	}
//
//	code, err := h.Service.Probs.Save(data)
//
//	Finish(c, "ok", err, code)
//
//}
//
//func (h *Probs) LookupProb(c *gin.Context) {
//	f := new(dtos.LookupProb)
//	f.ProbId = c.Param("problem_id")
//	f.Course = c.Query("course")
//	f.Question = c.Query("question")
//	f.ExactOnly, _ = strconv.ParseBool(c.Query("exactOnly"))
//
//	res, err := h.Service.Probs.LookupProb(f)
//
//	Finish(c, res.Prob, err, res.Code)
//
//}
