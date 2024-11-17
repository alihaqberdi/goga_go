package mw

import (
	"github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"github.com/alihaqberdi/goga_go/internal/pkg/app_errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func (mw *Middleware) AuthByRoles(roles ...types.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		if header == "" {
			response.FailErr(c, app_errors.AuthMwMissingToken)
			c.Abort()
			return
		}

		user, err := mw.JwtManager.Parse(header)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if !slices.Contains(roles, user.Role) {
			c.String(http.StatusUnauthorized, "permission denied")
			c.Abort()
			return
		}

		mw.SetUser(c, user)

		c.Next()

	}
}
