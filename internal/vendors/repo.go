package vendors

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VendorReporsitoryInterface interface {
	CreateVendor(vendor *models.Vendor) (*models.Vendor, error)
	GetVendors(pg *pgHelper.Pagination) (*[]models.Vendor, error)
}

type VendorRepository struct {
	db *gorm.DB
}

// NewVendorRepository creates a new vendor repository
func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{db: db}
}

// Migration migrates database
func (r *VendorRepository) Migration() {
	r.db.AutoMigrate(&models.Vendor{})
}

// CreateVendor creates a new vendor in database
func (r *VendorRepository) CreateVendor(vendor *models.Vendor) (*models.Vendor, error) {
	zap.L().Debug("vendor.repo.CreateVendor", zap.Reflect("vendor", *vendor))

	if err := r.db.Create(vendor).Error; err != nil {
		return nil, err
	}
	return vendor, nil
}

// GetVendors gets vendors from database with pagination
func (r *VendorRepository) GetVendors(pg *pgHelper.Pagination) (*[]models.Vendor, error) {
	zap.L().Debug("vendor.repo.GetVendors", zap.Reflect("pg", *pg))

	var vendors *[]models.Vendor
	var totalRows int64

	query := r.db.Model(&models.Vendor{}).Scopes(Search(pg.Q)).Count(&totalRows)
	query.Scopes(pgHelper.Paginate(totalRows, pg, r.db)).Find(&vendors)

	if query.Error != nil {
		return nil, query.Error
	}

	return vendors, nil
}
