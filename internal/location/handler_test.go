package location

import (
	"net/http"
	"net/http/httptest"
	"testing"

	locationMock "github.com/eibrahimarisoy/car_rental/mocks/location"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var lh LocationHandler
var mockService *locationMock.MockLocationServiceInterface

func setUpService(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService = locationMock.NewMockLocationServiceInterface(ctrl)
	lh = LocationHandler{locationService: mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestLocationHandler_GetAllLocations(t *testing.T) {
	trd := setUpService(t)
	defer trd()

	mockService.EXPECT().GetAllActiveLocations(&FakeDataWithPagination).Return(&FakeLocationsData, nil)

	router := gin.New()
	router.GET("/api/v1/locations", lh.GetAllLocations)

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/api/v1/locations", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("pagination", &FakeDataWithPagination)
	lh.GetAllLocations(c)

	assert.Equal(t, http.StatusOK, w.Code)
}
