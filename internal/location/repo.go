package location

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"gorm.io/gorm"
)

type LocationRepositoryInterface interface {
	GetAllLocations() ([]*models.Location, error)
}

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

func (r *LocationRepository) Migration() {
	r.db.AutoMigrate(&models.Location{})
}

func (r *LocationRepository) GetAllLocations() ([]*models.Location, error) {
	var locations []*models.Location
	err := r.db.Find(&locations).Error
	return locations, err
}
