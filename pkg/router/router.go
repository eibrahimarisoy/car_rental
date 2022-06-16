package router

import (
	"github.com/eibrahimarisoy/car_rental/internal/car"
	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/office"
	vendors "github.com/eibrahimarisoy/car_rental/internal/vendors"
	"github.com/eibrahimarisoy/car_rental/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitiliazeRoutes(rootRouter *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	// Initialize location service, repo and handler here
	locationGroup := rootRouter.Group("/locations")

	locationRepo := location.NewLocationRepository(db)
	locationRepo.Migration()

	locationService := location.NewLocationService(locationRepo)
	location.NewLocationHandler(locationGroup, locationService)

	// Initialize vendor service, repo and handler here
	vendorGroup := rootRouter.Group("/vendors")

	vendorRepo := vendors.NewVendorRepository(db)
	vendorRepo.Migration()

	vendorService := vendors.NewVendorService(vendorRepo)
	vendors.NewVendorHandler(vendorGroup, vendorService)

	// Initialize office service, repo and handler here
	officeGroup := rootRouter.Group("/offices")

	officeRepo := office.NewOfficeRepository(db)
	officeRepo.Migration()

	officeService := office.NewOfficeService(officeRepo)
	office.NewOfficeHandler(officeGroup, officeService)

	// Initialize car service, repo and handler here
	carGroup := rootRouter.Group("/cars")

	carRepo := car.NewCarRepository(db)
	carRepo.Migration()

	carService := car.NewCarService(carRepo)
	car.NewCarHandler(carGroup, carService)

}
