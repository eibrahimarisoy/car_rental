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

func (r *LocationResponse) FromModel(m *models.Location) *LocationResponse {
	r.ID = m.ID
	r.Name = *m.Name
	r.IsActive = m.IsActive
	r.CreatedAt = m.CreatedAt
	r.UpdatedAt = m.UpdatedAt

	return r
}
