package service

import (
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Service struct {
	Tenders *tenderService
	Auth    *Auth
	Bids    *bidsService
	//ApiLogging *ApiLogging
	//ApiAccess  *ApiAccess
	//Search     *Search
}

func New(repo_ *repo.Repo, cache *caching.Cache) *Service {
	return &Service{
		&tenderService{repo_, cache},
		&Auth{repo_, cache},
		&bidsService{repo_, cache},
		//&ApiLogging{repo_, cache},
		//&ApiAccess{repo_, cache},
		//&Search{repo_, cache},
	}
}
