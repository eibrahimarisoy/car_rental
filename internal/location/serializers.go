package location

import (
	"errors"
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"
)

type LocationRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (req *LocationRequest) Validate() error {
	if req.Name == "" {
		return errors.New("Name is required")
	}
	return nil
}

func (req *LocationRequest) ToModel() *models.Location {
	return &models.Location{
		Name:     &req.Name,
		IsActive: req.IsActive,
	}
}

type LocationResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LocationToResponse converts a location to a response
func LocationToResponse(location *models.Location) *LocationResponse {
	return &LocationResponse{
		ID:        location.ID,
		Name:      *location.Name,
		IsActive:  location.IsActive,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}
}
