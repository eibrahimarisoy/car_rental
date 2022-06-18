package mw

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	"github.com/eibrahimarisoy/car_rental/pkg/filters"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PaginationMiddleware is a middleware to paginate the results
func CarFilterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := &filters.CarFilter{}

		date := models.JsonDate{}
		time := models.JsonTime{}

		date = date.FromString(c.Query("pickup_date"))
		filter.PickupDate = date

		time = time.FromString(c.Query("pickup_time"))
		filter.PickupTime = time

		date = date.FromString(c.Query("dropoff_date"))
		filter.DropoffDate = date

		time = time.FromString(c.Query("dropoff_time"))
		filter.DropoffTime = time

		locationId := c.Query("location_id")
		if locationId != "" {
			location, err := uuid.Parse(c.Query("location_id"))
			if err != nil {
				c.JSON(errorHandler.ErrorResponse(err))
				return
			}
			filter.Location = location
		}

		if err := filter.Validate(); err != nil {
			c.JSON(errorHandler.ErrorResponse(err))
			return
		}

		c.Set("carFilter", filter)

		c.Next()
		return
	}
}
