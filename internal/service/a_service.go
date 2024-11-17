package service

import (
	"github.com/alihaqberdi/goga_go/internal/pkg/jwt_manager"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Service struct {
	Tenders *Tenders
	Auth    *Auth
	Bids    *Bids
}

func New(repo_ *repo.Repo, cache *caching.Cache, jwtManager *jwt_manager.JwtManager) *Service {
	return &Service{
		&Tenders{repo_, cache},
		&Auth{repo_, cache, jwtManager},
		&Bids{repo_, cache},
	}
}
