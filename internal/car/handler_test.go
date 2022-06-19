package car

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eibrahimarisoy/car_rental/internal/location"
	carMock "github.com/eibrahimarisoy/car_rental/mocks/car"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var ch CarHandler
var mockService *carMock.MockCarServiceInterface

func setUpHandler(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService = carMock.NewMockCarServiceInterface(ctrl)
	ch = CarHandler{carService: mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestCarHandler_GetAllCars(t *testing.T) {
	trd := setUpHandler(t)
	defer trd()

	mockService.EXPECT().GetCars(&location.FakeDataWithPagination, &FakeCarFilterData).Return(&FakeCarList, nil)

	router := gin.New()
	router.GET("/api/v1/cars", ch.GetAllCars)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/api/v1/cars", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("pagination", &location.FakeDataWithPagination)
	c.Set("carFilter", &FakeCarFilterData)
	ch.GetAllCars(c)

	assert.Equal(t, http.StatusOK, w.Code)
}
