package service

import (
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
)

type Service struct {
	Probs *Probs
	Auth  *Auth
	//ApiLogging *ApiLogging
	//ApiAccess  *ApiAccess
	//Search     *Search
}

func New(repo_ *repo.Repo, cache *caching.Cache) *Service {
	return &Service{
		&Probs{repo_, cache},
		&Auth{repo_, cache},
		//&ApiLogging{repo_, cache},
		//&ApiAccess{repo_, cache},
		//&Search{repo_, cache},
	}
}
