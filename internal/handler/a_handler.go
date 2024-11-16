package handler

import (
	"github.com/alihaqberdi/goga_go/internal/handler/mw"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Handlers struct {
	Dev    *Dev
	MW     *mw.Middleware
	Probs  *Probs
	Search *Search
}

func New(serv *service.Service, _cache *caching.Cache) *Handlers {
	return &Handlers{
		&Dev{serv, _cache},
		mw.New(serv, _cache),
		&Probs{serv},
		&Search{serv},
	}
}
