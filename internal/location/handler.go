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
	r.POST("/", handler.CreateLocation)

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

// CreateLocation is a handler to create a location
// @Summary      Create a location
// @Description  Create a location with payload
// @Tags         location
// @Accept       json
// @Produce      json
// @Param        body body    location.LocationRequest  true  "Location payload"
// @Success      201  {object}  location.LocationResponse
// @Failure	     400  {object}  _type.APIErrorResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /locations/    [post]
func (h *LocationHandler) CreateLocation(c *gin.Context) {
	reqBody := LocationRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	if err := reqBody.Validate(); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	location := reqBody.ToModel()

	location, err := h.locationService.CreateLocation(location)
	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, LocationToResponse(location))

	return
}
