package office

import (
	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/vendors"
	"github.com/google/uuid"
)

type OfficeRequest struct {
	OpeningHours models.JsonTime `json:"opening_hours"`
	ClosingHours models.JsonTime `json:"closing_hours"`
	VendorID     uuid.UUID       `json:"vendor_id"`
	LocationID   uuid.UUID       `json:"location_id"`
	WorkingDays  []uint          `json:"working_days"`
}

// Validate validates the OfficeRequest
func (r *OfficeRequest) Validate() error {
	_, err := uuid.Parse(r.VendorID.String())
	if err != nil {
		return err
	}
	return nil
}

// ToOffice converts the OfficeRequest to Office
func (r *OfficeRequest) ToOffice() *models.Office {
	workingDays := []models.WorkingDay{}
	for _, day := range r.WorkingDays {
		workingDays = append(workingDays, models.WorkingDay{
			Value: day,
		})
	}

	return &models.Office{
		OpeningHours: r.OpeningHours,
		ClosingHours: r.ClosingHours,
		VendorID:     r.VendorID,
		LocationID:   r.LocationID,
		WorkingDays:  workingDays,
	}
}

type OfficeResponse struct {
	ID           uuid.UUID                 `json:"id"`
	OpeningHours models.JsonTime           `json:"opening_hours"`
	ClosingHours models.JsonTime           `json:"closing_hours"`
	VendorID     uuid.UUID                 `json:"vendor_id"`
	Vendor       vendors.VendorResponse    `json:"vendor"`
	LocationID   uuid.UUID                 `json:"location_id"`
	Location     location.LocationResponse `json:"location"`
	WorkingDays  []string                  `json:"working_days"`
}

// FromOffice converts the Office to OfficeResponse
func (r *OfficeResponse) FromOffice(office *models.Office) {

	r.ID = office.ID
	r.OpeningHours = office.OpeningHours
	r.ClosingHours = office.ClosingHours
	r.VendorID = office.VendorID
	r.LocationID = office.LocationID

	r.Vendor = vendors.VendorResponse(r.Vendor)
	r.Location = location.LocationResponse(r.Location)

	r.WorkingDays = []string{}
	for _, day := range office.WorkingDays {
		r.WorkingDays = append(r.WorkingDays, day.Day)
	}

}
