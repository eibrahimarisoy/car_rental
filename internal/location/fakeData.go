package location

import (
	"time"

	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	"github.com/google/uuid"
)

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
