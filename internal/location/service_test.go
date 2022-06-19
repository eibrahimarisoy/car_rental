package location

import (
	"testing"

	locationMock "github.com/eibrahimarisoy/car_rental/mocks/location"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockLocationRepository *locationMock.MockLocationRepositoryInterface
var service *LocationService

func setUp(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockLocationRepository = locationMock.NewMockLocationRepositoryInterface(ct)
	service = NewLocationService(mockLocationRepository)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestLocationService_GetAllActiveLocations(t *testing.T) {
	td := setUp(t)
	defer td()

	mockLocationRepository.EXPECT().GetAllActiveLocations(&FakeDataWithPagination).Return(&FakeLocationsData, nil)

	result, err := service.GetAllActiveLocations(&FakeDataWithPagination)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotEmpty(t, result)
	assert.Equal(t, len(FakeLocationsData), len(*result))
}
