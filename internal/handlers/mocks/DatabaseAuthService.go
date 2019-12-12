// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/ilyakaznacheev/roster/internal/database/models"
	mock "github.com/stretchr/testify/mock"
)

// DatabaseAuthService is an autogenerated mock type for the DatabaseAuthService type
type DatabaseAuthService struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: c
func (_m *DatabaseAuthService) AddUser(c models.Credentials) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Credentials) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: login
func (_m *DatabaseAuthService) GetUser(login string) (*models.Credentials, error) {
	ret := _m.Called(login)

	var r0 *models.Credentials
	if rf, ok := ret.Get(0).(func(string) *models.Credentials); ok {
		r0 = rf(login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Credentials)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}