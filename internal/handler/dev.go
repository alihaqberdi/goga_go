package handler

import (
	. "github.com/alihaqberdi/goga_go/internal/handler/response"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Dev struct {
	Service *service.Service
	Cache   *caching.Cache
}

func (h *Dev) CacheFlush(c *gin.Context) {

	h.Cache.ProbAction.Flush()
	h.Cache.ApiAccess.Flush()
	h.Cache.Clients.Flush()
	h.Cache.Origins.Flush()

	Success(c, "done")
}

func (h *Dev) GetEnv(c *gin.Context) {
	envs := os.Environ()
	Success(c, envs)
}

func (h *Dev) GetIP(c *gin.Context) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(
		resp.StatusCode,
		resp.ContentLength,
		resp.Header.Get("Content-Type"),
		resp.Body, map[string]string{},
	)

}
