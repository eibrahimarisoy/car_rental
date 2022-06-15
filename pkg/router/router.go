package router

import (
	"fmt"

	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitiliazeRoutes(rootRouter *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {

	locationGroup := rootRouter.Group("/locations")
	fmt.Println("locationGroup:", locationGroup)

	locationRepo := location.NewLocationRepository(db)
	locationRepo.Migration()

}
