package caching

import (
	"github.com/alihaqberdi/goga_go/internal/config"
	"github.com/alihaqberdi/goga_go/internal/models"
	"math/rand/v2"
	"time"
)

type Cache struct {
	ProbAction *ProbAction
	ApiAccess  *ApiAccess
	Clients    *Clients
	Origins    *Origins
}

func New() *Cache {

	rCache := func() *cache.Cache {
		ri := time.Duration(rand.Int64N(42) + 100)
		dur := (ri * config.CACHING_EXPIRATION_DURATION) / 100
		return newCache(dur)
	}

	return &Cache{
		&ProbAction{probId2Mu: rCache()},
		&ApiAccess{AbsCache[models.ApiAccess]{key2val: rCache()}},
		&Clients{AbsCache[models.Client]{key2val: rCache()}},
		&Origins{AbsCache[models.Origin]{key2val: rCache()}},
	}

}

func newCache(ed time.Duration) *cache.Cache {
	cd := (ed * 134) / 100

	return cache.New(ed, cd)
}
