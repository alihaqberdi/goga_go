package models

import (
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"time"
)

type Prob struct {
	DocObjectID `bson:",inline" json:"-"`
	ProbID      string      `bson:"probId" json:"probId"`
	Course      string      `bson:"course" json:"course"`
	Solved      bool        `bson:"solved" json:"solved"`
	Type        types.Prob  `bson:"type" json:"type"`
	UpdatedAt   time.Time   `bson:"updatedAt" json:"updatedAt"`
	Variations  []Variation `bson:"variations" json:"variations"`
}

type Variation struct {
	Solved   bool       `bson:"solved" json:"solved"`
	Question string     `bson:"question" json:"question"`
	MAnswers [][]Answer `bson:"mAnswers" json:"mAnswers,omitempty"`
	SAnswers []Answer   `bson:"sAnswers" json:"sAnswers,omitempty"`
}

type Answer struct {
	ID   string `bson:"id" json:"id"`
	Text string `bson:"text" json:"text"`
}
