package repo

type Repo struct {
	Probs *Probs
}

func New() *Repo {
	return &Repo{
		&Probs{},
	}
}
