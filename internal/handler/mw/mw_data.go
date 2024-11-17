package mw

import (
	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/gin-gonic/gin"
)

const contextUserKey = "ContextUserKey"

func (mw *Middleware) SetUser(c *gin.Context, user *dtos.JwtUser) {
	c.Set(contextUserKey, user)
}

func (mw *Middleware) GetUser(c *gin.Context) (*dtos.JwtUser, bool) {
	val, ok := c.Get(contextUserKey)
	if !ok {
		return nil, false
	}

	user, ok := val.(*dtos.JwtUser)

	return user, ok
}
