package car

import (
	"net/http"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	mw "github.com/eibrahimarisoy/car_rental/pkg/middlewares"
	paginationHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	carService CarServiceInterface
}

func NewCarHandler(r *gin.RouterGroup, carService CarServiceInterface) {
	handler := &CarHandler{
		carService: carService,
	}

	r.GET("/", mw.PaginationMiddleware(), handler.GetAllCars)
	r.POST("/", handler.CreateCar)

}

// GetAllCars is a handler to get all cars
// @Summary      List all cars
// @Description  List all cars with pagination and search
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        q     query    string  false  "Search query"
// @Param        page  query    int     false  "Page number"
// @Param        limit query    int     false  "Page limit"
// @Success      200  {object}  pagination.Pagination
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /cars/    [get]
func (h *CarHandler) GetAllCars(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)
	filter := &CarFilter{}
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

	pagination, err := h.carService.GetCars(pagination, filter)

	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, pagination)
}

// CreateCar is a handler to create a car
// @Summary      Create a car
// @Description  Create a car with payload
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        body body    car.CarRequest  true  "Car payload"
// @Success      200  {object}  car.CarResponse
// @Failure	     400  {object}  _type.APIErrorResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /cars/    [post]
func (h *CarHandler) CreateCar(c *gin.Context) {
	reqBody := CarRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	if err := reqBody.Validate(); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	car := reqBody.ToCar()

	car, err := h.carService.CreateCar(car)
	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	carResponse := CarResponse{}
	carResponse.FromCar(car)
	c.JSON(http.StatusOK, carResponse)

	return

}
