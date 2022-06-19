package vendors

import (
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/google/uuid"
)

var vendorName1, vendorName2 = "vendorName1", "vendorName2"
var (
	FakeVendor_1 = models.Vendor{
		Base: models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name: &vendorName1,
	}
	FakeVendor_2 = models.Vendor{
		Base: models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name: &vendorName2,
	}
)

var FakeVendorList = []models.Vendor{FakeVendor_1, FakeVendor_2}
