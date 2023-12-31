// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/j1cs/api-user/internal/v1/domain"
)

// UserPublisher is an autogenerated mock type for the UserPublisher type
type UserPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, user, header
func (_m *UserPublisher) Publish(ctx context.Context, user domain.User, header domain.Header) (string, error) {
	ret := _m.Called(ctx, user, header)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User, domain.Header) (string, error)); ok {
		return rf(ctx, user, header)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User, domain.Header) string); ok {
		r0 = rf(ctx, user, header)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User, domain.Header) error); ok {
		r1 = rf(ctx, user, header)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserPublisher creates a new instance of UserPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserPublisher {
	mock := &UserPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
