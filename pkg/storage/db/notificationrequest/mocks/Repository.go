// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	notificationrequest "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/storage/db/notificationrequest"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ComponentExists provides a mock function with given fields: id, name
func (_m *Repository) ComponentExists(id string, name string) bool {
	ret := _m.Called(id, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(id, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Create provides a mock function with given fields: n
func (_m *Repository) Create(n *notificationrequest.NotificationRequest) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(*notificationrequest.NotificationRequest) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() ([]*notificationrequest.NotificationRequest, error) {
	ret := _m.Called()

	var r0 []*notificationrequest.NotificationRequest
	if rf, ok := ret.Get(0).(func() []*notificationrequest.NotificationRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*notificationrequest.NotificationRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHashkey provides a mock function with given fields: n
func (_m *Repository) UpdateHashkey(n *notificationrequest.NotificationRequest) error {
	ret := _m.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(*notificationrequest.NotificationRequest) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
