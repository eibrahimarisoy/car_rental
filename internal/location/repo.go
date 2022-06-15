package location

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"gorm.io/gorm"
)

type LocationRepositoryInterface interface {
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
