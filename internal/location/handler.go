package location

import (
	"fmt"
	"net/http"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
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

// GetAllLocations is a handler to get all locations
// @Summary      List all locations
// @Description  List all locations with pagination and search
// @Tags         location
// @Accept       json
// @Produce      json
// @Param        q     query    string  false  "Search query"
// @Param        page  query    int     false  "Page number"
// @Param        limit query    int     false  "Page limit"
// @Success      200  {object}  pagination.Pagination
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /locations/    [get]
func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)

	pagination, err := h.locationService.GetAllActiveLocations(pagination)
	fmt.Println(err)
	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, pagination)
}
