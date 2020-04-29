// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pubsub "cloud.google.com/go/pubsub"
)

// PublishResultWaiter is an autogenerated mock type for the PublishResultWaiter type
type PublishResultWaiter struct {
	mock.Mock
}

// WaitForResult provides a mock function with given fields: _a0, _a1
func (_m *PublishResultWaiter) WaitForResult(_a0 context.Context, _a1 *pubsub.PublishResult) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *pubsub.PublishResult) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pubsub.PublishResult) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}