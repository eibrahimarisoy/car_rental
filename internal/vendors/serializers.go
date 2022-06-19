package vendors

import (
	"errors"
	"fmt"
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/google/uuid"
)

type VendorRequest struct {
	Name *string `json:"name"`
}

func (req *VendorRequest) Validate() error {
	if req.Name == nil || *req.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (req *VendorRequest) ToVendor() *models.Vendor {
	return &models.Vendor{
		Name: req.Name,
	}
}

type VendorListResponse struct {
	pgHelper.Pagination
	Data []VendorResponse `json:"data"`
}

type VendorResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// VendorToResponse converts a vendor to a response
func VendorToResponse(vendor *models.Vendor) *VendorResponse {
	fmt.Println(vendor)
	vendorResponse := &VendorResponse{
		ID:        vendor.ID,
		Name:      *vendor.Name,
		CreatedAt: vendor.CreatedAt,
		UpdatedAt: vendor.UpdatedAt,
	}
	return vendorResponse
}

// VendorsToVendorListResponse converts a list of reservations to a reservation list response
func VendorsToVendorListResponse(reservations *[]models.Vendor, pagination *pgHelper.Pagination) *VendorListResponse {
	response := &VendorListResponse{}
	response.Pagination = *pagination
	for _, reservation := range *reservations {
		response.Data = append(response.Data, *VendorToResponse(&reservation))
	}
	return response
}
