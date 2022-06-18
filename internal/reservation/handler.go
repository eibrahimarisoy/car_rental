package reservation

import (
	"net/http"

	"github.com/eibrahimarisoy/car_rental/pkg/errorHandler"
	mw "github.com/eibrahimarisoy/car_rental/pkg/middlewares"
	paginationHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	reservationService ReservationServiceInterface
}

func NewReservationHandler(r *gin.RouterGroup, reservationService ReservationServiceInterface) {
	handler := &ReservationHandler{
		reservationService: reservationService,
	}

	r.GET("/", mw.PaginationMiddleware(), handler.GetAllReservations)
	r.POST("/", handler.CreateReservation)

}

// GetAllReservations is a handler to get all reservations
// @Summary      List all reservations
// @Description  List all reservations with pagination and search
// @Tags         reservation
// @Accept       json
// @Produce      json
// @Param        q     query    string  false  "Search query"
// @Param        page  query    int     false  "Page number"
// @Param        limit query    int     false  "Page limit"
// @Success      200  {object}  reservation.ReservationListResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /reservations/    [get]
func (h *ReservationHandler) GetAllReservations(c *gin.Context) {
	pagination := c.MustGet("pagination").(*paginationHelper.Pagination)

	reservations, err := h.reservationService.GetReservations(pagination)
	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, ReservationsToReservationListResponse(reservations, pagination))
}

// CreateReservation is a handler to create a reservation
// @Summary      Create a reservation
// @Description  Create a reservation with payload
// @Tags         reservation
// @Accept       json
// @Produce      json
// @Param        body body    reservation.ReservationRequest  true  "Reservation payload"
// @Success      201  {object}  reservation.ReservationResponse
// @Failure	     400  {object}  _type.APIErrorResponse
// @Failure      500  {object}  _type.APIErrorResponse
// @Router       /reservations/    [post]
func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	reqBody := ReservationRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	if err := reqBody.Validate(); err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	reservation, err := h.reservationService.CreateReservation(reqBody.ToReservation())
	if err != nil {
		c.JSON(errorHandler.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, ReservationToResponse(reservation))

	return
}
