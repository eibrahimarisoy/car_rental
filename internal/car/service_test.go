package car

import (
	"testing"

	carMock "github.com/eibrahimarisoy/car_rental/mocks/car"
	locationMock "github.com/eibrahimarisoy/car_rental/mocks/location"
	officeMock "github.com/eibrahimarisoy/car_rental/mocks/office"
	"github.com/golang/mock/gomock"
)

var mockCarRepository *carMock.MockCarRepositoryInterface
var mockLocationRepository *locationMock.MockLocationRepositoryInterface
var mockOfficeRepository *officeMock.MockOfficeRepositoryInterface
var service *CarService

func setUpService(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockCarRepository = carMock.NewMockCarRepositoryInterface(ct)
	mockLocationRepository = locationMock.NewMockLocationRepositoryInterface(ct)
	mockOfficeRepository = officeMock.NewMockOfficeRepositoryInterface(ct)

	service = NewCarService(mockCarRepository, mockLocationRepository, mockOfficeRepository)

	return func() {
		service = nil
		defer ct.Finish()
	}
}
