package repo

import "gorm.io/gorm"

type Repo struct {
	Probs *Probs
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		&Probs{db},
	}
}
