package models

import (
	"github.com/alihaqberdi/goga_go/internal/models/types"
	"time"
)

type Tender struct {
	Id          int
	ClientId    int
	Title       string
	Description string
	Deadline    time.Time
	Budget      float64
	Status      types.TenderStatus
}
