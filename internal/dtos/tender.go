package dtos

import (
	"time"

	"github.com/alihaqberdi/goga_go/internal/models/types"
)

type Tender struct {
	ID          uint               `json:"id"`
	ClientId    uint               `json:"client_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Deadline    time.Time          `json:"deadline"`
	Budget      float64            `json:"budget"`
	Status      types.TenderStatus `json:"status"`
}
