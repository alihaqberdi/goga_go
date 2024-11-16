package repo

import (
	"github.com/alihaqberdi/goga_go/internal/config/static"
	"github.com/alihaqberdi/goga_go/internal/models"
)

type Repo struct {
	Probs     *Probs
	ApiAccess *ApiAccess
	Clients   *Clients
	Origins   *Origins
}

func New() *Repo {
	return &Repo{
		&Probs{CrudRepo[*models.Prob]{static.CollProbs}},
		&ApiAccess{CrudRepo[*models.ApiAccess]{static.CollApiAccess}},
		&Clients{},
		&Origins{},
	}
}
