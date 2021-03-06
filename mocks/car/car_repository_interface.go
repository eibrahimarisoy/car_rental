// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eibrahimarisoy/car_rental/internal/car (interfaces: CarRepositoryInterface)

// Package car is a generated GoMock package.
package car

import (
	reflect "reflect"

	models "github.com/eibrahimarisoy/car_rental/internal/models"
	pagination "github.com/eibrahimarisoy/car_rental/pkg/pagination"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockCarRepositoryInterface is a mock of CarRepositoryInterface interface.
type MockCarRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCarRepositoryInterfaceMockRecorder
}

// MockCarRepositoryInterfaceMockRecorder is the mock recorder for MockCarRepositoryInterface.
type MockCarRepositoryInterfaceMockRecorder struct {
	mock *MockCarRepositoryInterface
}

// NewMockCarRepositoryInterface creates a new mock instance.
func NewMockCarRepositoryInterface(ctrl *gomock.Controller) *MockCarRepositoryInterface {
	mock := &MockCarRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockCarRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarRepositoryInterface) EXPECT() *MockCarRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateCar mocks base method.
func (m *MockCarRepositoryInterface) CreateCar(arg0 *models.Car) (*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCar", arg0)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCar indicates an expected call of CreateCar.
func (mr *MockCarRepositoryInterfaceMockRecorder) CreateCar(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCar", reflect.TypeOf((*MockCarRepositoryInterface)(nil).CreateCar), arg0)
}

// GetCarByID mocks base method.
func (m *MockCarRepositoryInterface) GetCarByID(arg0 uuid.UUID) (*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCarByID", arg0)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCarByID indicates an expected call of GetCarByID.
func (mr *MockCarRepositoryInterfaceMockRecorder) GetCarByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCarByID", reflect.TypeOf((*MockCarRepositoryInterface)(nil).GetCarByID), arg0)
}

// GetCarsByOfficeIDs mocks base method.
func (m *MockCarRepositoryInterface) GetCarsByOfficeIDs(arg0 *pagination.Pagination, arg1 []uuid.UUID) (*[]models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCarsByOfficeIDs", arg0, arg1)
	ret0, _ := ret[0].(*[]models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCarsByOfficeIDs indicates an expected call of GetCarsByOfficeIDs.
func (mr *MockCarRepositoryInterfaceMockRecorder) GetCarsByOfficeIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCarsByOfficeIDs", reflect.TypeOf((*MockCarRepositoryInterface)(nil).GetCarsByOfficeIDs), arg0, arg1)
}

// UpdateCarStatus mocks base method.
func (m *MockCarRepositoryInterface) UpdateCarStatus(arg0 *models.Car) (*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCarStatus", arg0)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCarStatus indicates an expected call of UpdateCarStatus.
func (mr *MockCarRepositoryInterfaceMockRecorder) UpdateCarStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCarStatus", reflect.TypeOf((*MockCarRepositoryInterface)(nil).UpdateCarStatus), arg0)
}
