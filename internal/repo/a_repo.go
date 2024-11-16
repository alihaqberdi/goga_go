package repo

import "gorm.io/gorm"

type Repo struct {
	Tenders *Tenders
	Users   *Users
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		&Tenders{db},
		&Users{db},
	}
}
