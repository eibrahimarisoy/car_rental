package location

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"

	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
)

type LocationRepositoryInterface interface {
	GetAllActiveLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateLocation(location *models.Location) (*models.Location, error)
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

func (r *LocationRepository) GetAllActiveLocations(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	var locations []*models.Location
	var totalRows int64

	query := r.db.Model(&models.Location{}).Where("is_active = ?", true).Scopes(Search(pg.Q)).Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&locations)

	if query.Error != nil {
		return nil, query.Error
	}

	pg.Rows = &locations
	return pg, nil
}

func (r *LocationRepository) CreateLocation(location *models.Location) (*models.Location, error) {

	if err := r.db.Create(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}
