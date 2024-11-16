package handler

import (
	"github.com/alihaqberdi/goga_go/internal/handler/mw"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Handlers struct {
	MW    *mw.Middleware
	Probs *Probs
}

func New(serv *service.Service, _cache *caching.Cache) *Handlers {
	return &Handlers{
		mw.New(serv, _cache),
		&Probs{serv},
	}
}
