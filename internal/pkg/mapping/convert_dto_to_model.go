package mapping

import (
	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/alihaqberdi/goga_go/internal/models"
)

func ConvertTenderDTOToModel(dto *dtos.Tender) *models.Tender {
	return &models.Tender{
		ClientId:    dto.ClientId,
		Title:       dto.Title,
		Description: dto.Description,
		Deadline:    dto.Deadline,
		Budget:      dto.Budget,
		Status:      dto.Status,
	}
}
