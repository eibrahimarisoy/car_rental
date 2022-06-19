package car

import (
	"fmt"
	"net/http"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	filterHelper "github.com/eibrahimarisoy/car_rental/pkg/filters"
	mw "github.com/eibrahimarisoy/car_rental/pkg/middlewares"
	paginationHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	carService CarServiceInterface
}

func NewCarHandler(r *gin.RouterGroup, carService CarServiceInterface) {
	handler := &CarHandler{
		carService: carService,
	}

	r.GET("/", mw.PaginationMiddleware(), mw.CarFilterMiddleware(), handler.GetAllCars)
	r.POST("/", handler.CreateCar)

}

// GetAllCars is a handler to get all cars
// @Summary      List all cars
// @Description  List all cars with pagination and search
// @Tags         car
// @Accept       json
// @Produce      json
// @Param        q     			query   string  false  	"Search query"
// @Param        page  			query   int     false  	"Page number"
// @Param        limit 			query   int     false  	"Page limit"
// @Param		 pickup_date 	query 	string 	true 	"Pickup date"
// @Param		 pickup_time 	query 	string 	true 	"Pickup time"
// @Param		 dropoff_date 	query 	string 	true 	"Dropoff date"
// @Param		 dropoff_time 	query 	string 	true 	"Dropoff time"
// @Param		 location_id 	query	string	true 	"UUID formatted ID"
// @Success      200  {object}  car.CarListResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /cars/    [get]
func (h *CarHandler) GetAllCars(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)
	filter := c.MustGet("carFilter").(*filterHelper.CarFilter)

	fmt.Println(*filter)
	cars, err := h.carService.GetCars(pagination, filter)

	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, CarsToCarListResponse(cars, pagination))
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

	c.JSON(http.StatusCreated, CarToCarResponse(car))

	return
}
