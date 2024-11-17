package dtos

import (
	"time"

	"github.com/alihaqberdi/goga_go/internal/models/types"
)

type Tender struct {
	ID          uint               `json:"id"`
	ClientId    uint               `json:"username"`
	Title       string             `json:"role"`
	Description string             `json:"email"`
	Deadline    time.Time          `json:"deadline"`
	Budget      float64            `json:"budget"`
	Status      types.TenderStatus `json:"status"`
}
