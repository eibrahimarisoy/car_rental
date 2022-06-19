package location

import (
	"testing"
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	locationMock "github.com/eibrahimarisoy/car_rental/mocks/location"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var mockLocationRepository *locationMock.MockLocationRepositoryInterface
var service *LocationService

var name1, name2, name3 = "Location 1", "Location 2", "Location 3"

var FakeData = []models.Location{
	{
		Base:     models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     &name1,
		IsActive: true,
	},
	{
		Base:     models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     &name2,
		IsActive: true,
	},
	{
		Base:     models.Base{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     &name3,
		IsActive: false,
	},
}

var FakeDataWithPagination = pgHelper.Pagination{
	Limit: 2,
	Page:  1,
	Q:     "",
}

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

	mockLocationRepository.EXPECT().GetAllActiveLocations(&FakeDataWithPagination).Return(&FakeData, nil)

	result, err := service.GetAllActiveLocations(&FakeDataWithPagination)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotEmpty(t, result)
	assert.Equal(t, len(FakeData), len(*result))

}
