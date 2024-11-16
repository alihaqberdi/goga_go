package repo

import "gorm.io/gorm"

type Repo struct {
	Tenders *Tenders
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		&Tenders{db},
	}
}
