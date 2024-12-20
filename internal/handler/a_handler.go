package handler

import (
	"github.com/alihaqberdi/goga_go/internal/handler/mw"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Handlers struct {
	Tenders *Tenders
	Auth    *Auth
	Bids    *Bids
}

func New(serv *service.Service, _cache *caching.Cache, mw *mw.Middleware) *Handlers {
	return &Handlers{
		&Tenders{serv, mw},
		&Auth{serv},
		&Bids{serv, mw},
	}
}
