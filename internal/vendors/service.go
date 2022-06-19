package vendors

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type VendorServiceInterface interface {
	CreateVendor(vendor *models.Vendor) (*models.Vendor, error)
	GetVendors(pg *pgHelper.Pagination) (*[]models.Vendor, error)
}

type VendorService struct {
	vendorRepo VendorReporsitoryInterface
}

// NewVendorService creates a new vendor service
func NewVendorService(vendorRepo VendorReporsitoryInterface) *VendorService {
	return &VendorService{vendorRepo: vendorRepo}
}

// CreateVendor creates a new vendor in database
func (s *VendorService) CreateVendor(vendor *models.Vendor) (*models.Vendor, error) {
	return s.vendorRepo.CreateVendor(vendor)
}

// GetVendors gets vendors from database with pagination
func (s *VendorService) GetVendors(pg *pgHelper.Pagination) (*[]models.Vendor, error) {
	return s.vendorRepo.GetVendors(pg)
}
