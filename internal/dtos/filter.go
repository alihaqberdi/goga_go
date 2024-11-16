package dtos

type LookupProb struct {
	ProbId    string
	Course    string
	Question  string
	ExactOnly bool
}

type SearchProbs struct {
	ProbId   string
	Question string
}
