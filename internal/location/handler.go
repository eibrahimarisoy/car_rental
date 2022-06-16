package location

import (
	mw "github.com/eibrahimarisoy/car_rental/pkg/middlewares"
	paginationHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	locationService LocationServiceInterface
}

func NewLocationHandler(r *gin.RouterGroup, locationService LocationServiceInterface) {
	handler := &LocationHandler{
		locationService: locationService,
	}

	r.GET("/", mw.PaginationMiddleware(), handler.GetAllLocations)

}

func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)

	pagination, err := h.locationService.GetAllLocations(pagination)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"locations": pagination,
	})
}
