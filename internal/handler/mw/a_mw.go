package mw

import (
	"github.com/alihaqberdi/goga_go/internal/pkg/jwt_manager"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Middleware struct {
	Service    *service.Service
	Cache      *caching.Cache
	JwtManager *jwt_manager.JwtManager
}

func New(service *service.Service, cache *caching.Cache, jwtManager *jwt_manager.JwtManager) *Middleware {
	return &Middleware{
		service,
		cache,
		jwtManager,
	}
}
