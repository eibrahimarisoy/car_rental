package location

import (
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	locationService LocationServiceInterface
}

func NewLocationHandler(r *gin.RouterGroup, locationService LocationServiceInterface) {

	handler := &LocationHandler{
		locationService: locationService,
	}

	r.GET("/", handler.GetAllLocations)

}

func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	locations, err := h.locationService.GetAllLocations()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"locations": locations,
	})
}
