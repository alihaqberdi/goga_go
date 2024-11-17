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

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dtos.Register true "Register"
// @Success 201 {object} dtos.AuthRes
// @Router /auth/register [post]
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

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dtos.Login true "Login"
// @Success 200 {object} dtos.AuthRes
// @Router /auth/login [post]
func (h *Auth) Login(c *gin.Context) {

	data, err := bind[dtos.Login](c)
	if HasErr(c, err) {
		return
	}

	res, err := h.Service.Auth.Login(data)

	if HasErr(c, err, 401) {
		return
	}

	Success(c, res, 200)

}
