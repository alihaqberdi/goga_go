package mw

import (
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Middleware struct {
	Service *service.Service
	Cache   *caching.Cache
}

func New(service *service.Service, cache *caching.Cache) *Middleware {
	return &Middleware{
		service,
		cache,
	}
}
