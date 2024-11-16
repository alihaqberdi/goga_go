package service

import (
	"errors"
	"fmt"
	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/alihaqberdi/goga_go/internal/models"
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
	"slices"
)

type Probs struct {
	Repo  *repo.Repo
	Cache *caching.Cache
}

func (s *Probs) Save(data *dtos.SaveProb) (status.Code, error) {
	mu := s.Cache.ProbAction.GetMu(data.ProbId)
	mu.Lock()
	defer mu.Unlock()

	return s.save(data)
}

func (s *Probs) LookupProb(f *dtos.LookupProb) (*dtos.ResProb, error) {
	prob, err := s.lookupProb(f)
	res := new(dtos.ResProb)

	if err != nil {
		if isErrNoDocuments(err) {
			res.Code = status.ProbNotFound
		}

		return res, err
	}

	res.Prob = prob
	res.Code = status.ProbIsNotSolved
	if prob.Solved {
		res.Code = status.ProbIsSolved
	}

	return res, err
}

func (s *Probs) save(data *dtos.SaveProb) (status.Code, error) {
	if types.ParseProb(data.Type) < 0 {
		return 0, fmt.Errorf(`invalid problem type:"%s"`, data.Type)
	}

	prob, err := s.Repo.Probs.Find(data.ProbId, data.Course)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, err = s.Repo.Probs.Create(data.GetProb())
			if err != nil {
				return 0, err
			}

			code := status.ProbCreated0
			if data.Solved {
				code = status.ProbCreated1
			}
			return code, nil
		}

		return 0, err
	}

	indexByQuestion := slices.IndexFunc(prob.Variations, func(v models.Variation) bool {
		return v.Question == data.Variation.Question
	})
	if indexByQuestion > -1 {
		variation := prob.Variations[indexByQuestion]
		if variation.Solved {
			if !data.Variation.Solved {
				return status.ProbAlreadyExist, nil
			}

			if ContainsAllAnswers(variation, data.Variation) {
				return status.ProbAlreadySolved, nil
			}

			variation.SAnswers = append(variation.SAnswers, data.SAnswers...)
			variation.MAnswers = append(variation.MAnswers, data.MAnswers...)

			err = s.update(prob)
			if err != nil {
				return 0, err
			}

			return status.ProbSolvedUpdated, nil
		}

		if data.Variation.Solved {
			prob.Variations[indexByQuestion] = data.Variation

			err = s.update(prob)
			if err != nil {
				return 0, err
			}

			return status.ProbSuccessfullySolved, nil
		}

		if len(data.SAnswers) > 0 {
			has := ContainsAll(variation.SAnswers, data.SAnswers, answersEquals)

			if has {
				return status.ProbAlreadyExist, nil
			}
		}

		if len(data.MAnswers) > 0 {
			has := ContainsAll(variation.MAnswers, data.MAnswers, func(arr1, arr2 []models.Answer) bool {
				return slices.EqualFunc(arr1, arr2, answersEquals)
			})

			if has {
				return status.ProbAlreadyExist, nil
			}
		}

		variation.SAnswers = append(variation.SAnswers, data.SAnswers...)
		variation.MAnswers = append(variation.MAnswers, data.MAnswers...)

		err = s.update(prob)
		if err != nil {
			return 0, err
		}

		return status.ProbUpdated, nil
	}

	prob.Variations = append(prob.Variations, data.Variation)

	err = s.update(prob)
	if err != nil {
		return 0, err
	}

	code := status.ProbUpdated
	if data.Solved {
		code = status.ProbSuccessfullySolved
	}

	return code, nil

}

func (s *Probs) lookupProb(f *dtos.LookupProb) (*models.Prob, error) {
	prob, err := s.Repo.Probs.Find(f.ProbId, f.Course)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return s.lookupByQuestion(f.ProbId, f.Question)
	}
	if err != nil {
		return nil, err
	}

	if prob.Solved || f.Question == "" || f.ExactOnly {
		return prob, nil
	}

	prob1, err1 := s.lookupByQuestion(f.ProbId, f.Question)
	if err1 != nil {
		return prob, nil
	}

	if prob1.Solved {
		return prob1, nil
	}

	return prob, nil

}

func (s *Probs) lookupByQuestion(probId, question string) (*models.Prob, error) {
	probs, err := s.Repo.Probs.FindList(
		&dtos.ListOptions{Limit: 100, Filter: bson.M{"probId": probId}},
	)
	if err != nil {
		return nil, err
	}

	fProbs := make([]models.Prob, 0, 10)
p:
	for _, prob := range probs {
		for _, variation := range prob.Variations {
			if question == variation.Question {
				fProbs = append(fProbs, prob)
				continue p
			}
		}
	}

	if len(fProbs) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	for i, prob := range fProbs {
		if prob.Solved {
			return &fProbs[i], nil
		}
	}

	return &fProbs[0], nil
}

func (s *Probs) update(prob *models.Prob) error {

	if !prob.Solved {
		for _, v := range prob.Variations {
			if v.Solved {
				prob.Solved = true
				break
			}
		}
	}

	return s.Repo.Probs.Update(prob)
}
