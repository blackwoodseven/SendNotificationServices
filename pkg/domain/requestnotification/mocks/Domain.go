// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	requestnotification "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
	mock "github.com/stretchr/testify/mock"
)

// Domain is an autogenerated mock type for the Domain type
type Domain struct {
	mock.Mock
}

// RequestNotification provides a mock function with given fields: _a0
func (_m *Domain) RequestNotification(_a0 requestnotification.RequestNotification) (bool, bool, bool) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(requestnotification.RequestNotification) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(requestnotification.RequestNotification) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 bool
	if rf, ok := ret.Get(2).(func(requestnotification.RequestNotification) bool); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}

// Scheduler provides a mock function with given fields:
func (_m *Domain) Scheduler() {
	_m.Called()
}

type mockConstructorTestingTNewDomain interface {
	mock.TestingT
	Cleanup(func())
}

// NewDomain creates a new instance of Domain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDomain(t mockConstructorTestingTNewDomain) *Domain {
	mock := &Domain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
