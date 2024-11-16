package dtos

import (
	"github.com/alihaqberdi/goga_go/internal/models"
	"github.com/alihaqberdi/goga_go/internal/pkg/status"
)

type SaveProb struct {
	Variation
	Course string `json:"course"`
	ProbId string `json:"probId"`
	Type   string `json:"type"`

	LocationHref string      `json:"locationHref"`
	ResCode      status.Code `json:"-"`
}

type ResProb struct {
	*models.Prob
	Code status.Code
}

type ResSearch struct {
	ListByQuestion []models.Prob `json:"listByQuestion,omitempty"`
	ListByProbId   []models.Prob `json:"listByProbId,omitempty"`
}
