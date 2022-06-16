package router

import (
	"github.com/eibrahimarisoy/car_rental/internal/location"
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

}
