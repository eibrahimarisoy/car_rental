package office

import (
	"github.com/eibrahimarisoy/car_rental/internal/models"
	pgHelper "github.com/eibrahimarisoy/car_rental/pkg/pagination"
)

type OfficeService struct {
	officeRepo OfficeRepositoryInterface
}

type OfficeServiceInterface interface {
	GetOffices(pg *pgHelper.Pagination) (*[]models.Office, error)
	CreateOffice(office *models.Office) (*models.Office, error)
}

// NewOfficeService creates a new office service
func NewOfficeService(officeRepo OfficeRepositoryInterface) *OfficeService {
	return &OfficeService{
		officeRepo: officeRepo,
	}
}

// GetOffices returns all offices
func (s *OfficeService) GetOffices(pg *pgHelper.Pagination) (*[]models.Office, error) {
	return s.officeRepo.GetOffices(pg)
}

// CreateOffice creates a office and returns it
func (s *OfficeService) CreateOffice(office *models.Office) (*models.Office, error) {
	workingDays, err := s.officeRepo.GetWorkingDaysByValues(&office.WorkingDays)
	if err != nil {
		return nil, err
	}
	office.WorkingDays = *workingDays
	return s.officeRepo.CreateOffice(office)

}
