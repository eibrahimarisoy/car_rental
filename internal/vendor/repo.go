package vendor

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"gorm.io/gorm"
)

type VendorReporsitoryInterface interface {
	CreateVendor(vendor *models.Vendor) (*models.Vendor, error)
	GetVendors(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
}

type VendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{db: db}
}

func (r *VendorRepository) Migration() {
	r.db.AutoMigrate(&models.Vendor{})
}

func (r *VendorRepository) CreateVendor(vendor *models.Vendor) (*models.Vendor, error) {
	if err := r.db.Create(vendor).Error; err != nil {
		return nil, err
	}
	return vendor, nil
}

func (r *VendorRepository) GetVendors(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	var vendors []models.Vendor
	var totalRows int64

	query := r.db.Model(&models.Vendor{}).Scopes(Search(pg.Q)).Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&vendors)

	if query.Error != nil {
		return nil, query.Error
	}

	pg.Rows = &vendors

	return pg, nil
}
