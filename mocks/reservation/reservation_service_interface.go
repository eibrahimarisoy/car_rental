// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eibrahimarisoy/car_rental/internal/reservation (interfaces: ReservationServiceInterface)

// Package reservation is a generated GoMock package.
package reservation

import (
	reflect "reflect"

	models "github.com/eibrahimarisoy/car_rental/internal/models"
	pagination "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	gomock "github.com/golang/mock/gomock"
)

// MockReservationServiceInterface is a mock of ReservationServiceInterface interface.
type MockReservationServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockReservationServiceInterfaceMockRecorder
}

// MockReservationServiceInterfaceMockRecorder is the mock recorder for MockReservationServiceInterface.
type MockReservationServiceInterfaceMockRecorder struct {
	mock *MockReservationServiceInterface
}

// NewMockReservationServiceInterface creates a new mock instance.
func NewMockReservationServiceInterface(ctrl *gomock.Controller) *MockReservationServiceInterface {
	mock := &MockReservationServiceInterface{ctrl: ctrl}
	mock.recorder = &MockReservationServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReservationServiceInterface) EXPECT() *MockReservationServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateReservation mocks base method.
func (m *MockReservationServiceInterface) CreateReservation(arg0 *models.Reservation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReservation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReservation indicates an expected call of CreateReservation.
func (mr *MockReservationServiceInterfaceMockRecorder) CreateReservation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReservation", reflect.TypeOf((*MockReservationServiceInterface)(nil).CreateReservation), arg0)
}

// GetReservations mocks base method.
func (m *MockReservationServiceInterface) GetReservations(arg0 *pagination.Pagination) (*[]models.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservations", arg0)
	ret0, _ := ret[0].(*[]models.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservations indicates an expected call of GetReservations.
func (mr *MockReservationServiceInterfaceMockRecorder) GetReservations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservations", reflect.TypeOf((*MockReservationServiceInterface)(nil).GetReservations), arg0)
}
