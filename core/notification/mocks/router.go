// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	log "github.com/goto/siren/core/log"
	mock "github.com/stretchr/testify/mock"

	notification "github.com/goto/siren/core/notification"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

type Router_Expecter struct {
	mock *mock.Mock
}

func (_m *Router) EXPECT() *Router_Expecter {
	return &Router_Expecter{mock: &_m.Mock}
}

// PrepareMetaMessages provides a mock function with given fields: ctx, n
func (_m *Router) PrepareMetaMessages(ctx context.Context, n notification.Notification) ([]notification.MetaMessage, []log.Notification, error) {
	ret := _m.Called(ctx, n)

	if len(ret) == 0 {
		panic("no return value specified for PrepareMetaMessages")
	}

	var r0 []notification.MetaMessage
	var r1 []log.Notification
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, notification.Notification) ([]notification.MetaMessage, []log.Notification, error)); ok {
		return rf(ctx, n)
	}
	if rf, ok := ret.Get(0).(func(context.Context, notification.Notification) []notification.MetaMessage); ok {
		r0 = rf(ctx, n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]notification.MetaMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, notification.Notification) []log.Notification); ok {
		r1 = rf(ctx, n)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]log.Notification)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, notification.Notification) error); ok {
		r2 = rf(ctx, n)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Router_PrepareMetaMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PrepareMetaMessages'
type Router_PrepareMetaMessages_Call struct {
	*mock.Call
}

// PrepareMetaMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - n notification.Notification
func (_e *Router_Expecter) PrepareMetaMessages(ctx interface{}, n interface{}) *Router_PrepareMetaMessages_Call {
	return &Router_PrepareMetaMessages_Call{Call: _e.mock.On("PrepareMetaMessages", ctx, n)}
}

func (_c *Router_PrepareMetaMessages_Call) Run(run func(ctx context.Context, n notification.Notification)) *Router_PrepareMetaMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(notification.Notification))
	})
	return _c
}

func (_c *Router_PrepareMetaMessages_Call) Return(metaMessages []notification.MetaMessage, notificationLogs []log.Notification, err error) *Router_PrepareMetaMessages_Call {
	_c.Call.Return(metaMessages, notificationLogs, err)
	return _c
}

func (_c *Router_PrepareMetaMessages_Call) RunAndReturn(run func(context.Context, notification.Notification) ([]notification.MetaMessage, []log.Notification, error)) *Router_PrepareMetaMessages_Call {
	_c.Call.Return(run)
	return _c
}

// NewRouter creates a new instance of Router. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRouter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Router {
	mock := &Router{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}