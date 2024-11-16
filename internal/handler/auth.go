package handler

import (
	"github.com/alihaqberdi/goga_go/internal/dtos"
	. "github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Service *service.Service
}

func (h *Auth) Register(c *gin.Context) {

	data, err := bind[dtos.Register](c)
	if HasErr(c, err) {
		return
	}

	res, err := h.Service.Auth.Register(data)

	if HasErr(c, err) {
		return
	}

	Success(c, res, 201)

}

func (h *Auth) Login(c *gin.Context) {

}
