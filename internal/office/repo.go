package office

import (
	"fmt"

	"github.com/eibrahimarisoy/car_rental/internal/models"

	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
)

type OfficeRepositoryInterface interface {
	GetOffices(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateOffice(office *models.Office) (*models.Office, error)
}

type OfficeRepository struct {
	db *gorm.DB
}

func NewOfficeRepository(db *gorm.DB) *OfficeRepository {
	return &OfficeRepository{
		db: db,
	}
}

func (r *OfficeRepository) Migration() {
	r.db.AutoMigrate(&models.Office{})
	r.LoadWorkingDay()
}

// GetOffices returns all offices
func (r *OfficeRepository) GetOffices(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	var offices []models.Office
	var totalRows int64

	query := r.db.Model(&models.Office{}).Preload("Location").Preload("Vendor").Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&offices)

	if query.Error != nil {
		return nil, query.Error
	}

	pg.Rows = &offices
	return pg, nil

}

// CreateOffice creates a new office
func (r *OfficeRepository) CreateOffice(office *models.Office) (*models.Office, error) {
	tx := r.db.Begin()
	workingDays := []models.WorkingDay{}

	for _, item := range office.WorkingDays {
		r.db.First(&item, "value = ?", item.Value)
		workingDays = append(workingDays, item)
	}

	office.WorkingDays = workingDays
	if res := tx.Model(&office).Create(office); res.Error != nil {
		fmt.Println(res.Error)
		tx.Rollback()
		return nil, res.Error
	}

	tx.Commit()
	return office, nil
}

// CreateWorkingDay creates a new working day
func (r *OfficeRepository) LoadWorkingDay() error {
	workingDay := models.WorkingDay{}

	workingDays := []models.WorkingDay{
		{Value: 1, Day: "Monday"},
		{Value: 2, Day: "Tuesday"},
		{Value: 3, Day: "Wednesday"},
		{Value: 4, Day: "Thursday"},
		{Value: 5, Day: "Friday"},
		{Value: 6, Day: "Saturday"},
		{Value: 7, Day: "Sunday"},
	}

	for _, item := range workingDays {
		r.db.Model(&workingDay).FirstOrCreate(&item)
	}

	return nil
}
