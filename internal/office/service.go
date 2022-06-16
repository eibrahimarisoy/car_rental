package office

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type OfficeService struct {
	officeRepo OfficeRepositoryInterface
}

type OfficeServiceInterface interface {
	GetOffices(pg *pgHelper.Pagination) (*pgHelper.Pagination, error)
	CreateOffice(office *models.Office) (*models.Office, error)
}

func NewOfficeService(officeRepo OfficeRepositoryInterface) *OfficeService {
	return &OfficeService{
		officeRepo: officeRepo,
	}
}

func (s *OfficeService) GetOffices(pg *pgHelper.Pagination) (*pgHelper.Pagination, error) {
	return s.officeRepo.GetOffices(pg)
}

func (s *OfficeService) CreateOffice(office *models.Office) (*models.Office, error) {
	return s.officeRepo.CreateOffice(office)
}
