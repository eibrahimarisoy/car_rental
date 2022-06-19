package car

import (
	"fmt"
	"testing"

	"github.com/eibrahimarisoy/car_rental/internal/location"
	"github.com/eibrahimarisoy/car_rental/internal/models"
	"github.com/eibrahimarisoy/car_rental/internal/office"
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

func TestCarService_GetCars(t *testing.T) {
	td := setUpService(t)
	defer td()

	officeIds := office.OfficeIds

	mockCarRepository.EXPECT().GetCarsByOfficeIDs(location.FakeDataWithPagination, officeIds).Return(&[]models.Car{}, nil)

	result, err := service.GetCars(&location.FakeDataWithPagination, &FakeCarFilterData)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	fmt.Println(result)

}
