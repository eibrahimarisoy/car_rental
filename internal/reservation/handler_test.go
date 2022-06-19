package reservation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eibrahimarisoy/car_rental/internal/location"
	reservationMock "github.com/eibrahimarisoy/car_rental/mocks/reservation"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	rh          ReservationHandler
	mockService *reservationMock.MockReservationServiceInterface
)

func setUpService(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService = reservationMock.NewMockReservationServiceInterface(ctrl)
	rh = ReservationHandler{reservationService: mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestReservationHandler_GetAllReservations(t *testing.T) {
	trd := setUpService(t)
	defer trd()

	mockService.EXPECT().GetReservations(&location.FakeDataWithPagination).Return(&FakeReservationList, nil)
	router := gin.New()
	router.GET("/api/v1/locations", rh.GetAllReservations)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/api/v1/reservations", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("pagination", &location.FakeDataWithPagination)
	rh.GetAllReservations(c)

	assert.Equal(t, http.StatusOK, w.Code)

}
