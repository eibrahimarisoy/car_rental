package location

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

//go:generate mockgen -destination=../../mocks/location/location_repository_interface.go -package=location github.com/eibrahimarisoy/car_rental/internal/location LocationRepositoryInterface
type LocationRepositoryInterface interface {
	GetAllActiveLocations(pg *pgHelper.Pagination) (*[]models.Location, error)
	CreateLocation(location *models.Location) (*models.Location, error)
	GetLocationByID(id uuid.UUID) (*models.Location, error)
}

// NewLocationRepository creates a new location repository
func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

// Migration for location table
func (r *LocationRepository) Migration() {
	r.db.AutoMigrate(&models.Location{})
}

// GetAllActiveLocations gets all active locations from database
func (r *LocationRepository) GetAllActiveLocations(pg *pgHelper.Pagination) (*[]models.Location, error) {
	zap.L().Debug("location.repo.GetAllActiveLocations", zap.Reflect("pg", *pg))

	var locations *[]models.Location
	var totalRows int64

	query := r.db.Model(&models.Location{}).Where("is_active = ?", true).Scopes(Search(pg.Q)).Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&locations)

	if query.Error != nil {
		return nil, query.Error
	}

	return locations, nil
}

// CreateLocation creates a location and returns it
func (r *LocationRepository) CreateLocation(location *models.Location) (*models.Location, error) {
	zap.L().Debug("location.repo.CreateLocation", zap.Reflect("location", *location))

	if err := r.db.Create(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}

// GetLocationByID gets location from database by id
func (r *LocationRepository) GetLocationByID(id uuid.UUID) (*models.Location, error) {
	zap.L().Debug("location.repo.GetLocationByID", zap.Reflect("id", id))

	var location models.Location

	query := r.db.Model(&models.Location{}).Where("is_active = ? AND id = ?", true, id).First(&location)
	if query.Error == gorm.ErrRecordNotFound {
		return nil, errorHandler.LocationNotFoundError
	} else if query.Error != nil {
		return nil, query.Error
	}

	return &location, nil
}
