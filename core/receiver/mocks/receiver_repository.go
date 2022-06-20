// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	receiver "github.com/odpf/siren/core/receiver"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ReceiverRepository is an autogenerated mock type for the Repository type
type ReceiverRepository struct {
	mock.Mock
}

type ReceiverRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ReceiverRepository) EXPECT() *ReceiverRepository_Expecter {
	return &ReceiverRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *ReceiverRepository) Create(_a0 *receiver.Receiver) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*receiver.Receiver) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceiverRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ReceiverRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - _a0 *receiver.Receiver
func (_e *ReceiverRepository_Expecter) Create(_a0 interface{}) *ReceiverRepository_Create_Call {
	return &ReceiverRepository_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *ReceiverRepository_Create_Call) Run(run func(_a0 *receiver.Receiver)) *ReceiverRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*receiver.Receiver))
	})
	return _c
}

func (_c *ReceiverRepository_Create_Call) Return(_a0 error) *ReceiverRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *ReceiverRepository) Delete(_a0 uint64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceiverRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ReceiverRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - _a0 uint64
func (_e *ReceiverRepository_Expecter) Delete(_a0 interface{}) *ReceiverRepository_Delete_Call {
	return &ReceiverRepository_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *ReceiverRepository_Delete_Call) Run(run func(_a0 uint64)) *ReceiverRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *ReceiverRepository_Delete_Call) Return(_a0 error) *ReceiverRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Get provides a mock function with given fields: _a0
func (_m *ReceiverRepository) Get(_a0 uint64) (*receiver.Receiver, error) {
	ret := _m.Called(_a0)

	var r0 *receiver.Receiver
	if rf, ok := ret.Get(0).(func(uint64) *receiver.Receiver); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*receiver.Receiver)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiverRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ReceiverRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - _a0 uint64
func (_e *ReceiverRepository_Expecter) Get(_a0 interface{}) *ReceiverRepository_Get_Call {
	return &ReceiverRepository_Get_Call{Call: _e.mock.On("Get", _a0)}
}

func (_c *ReceiverRepository_Get_Call) Run(run func(_a0 uint64)) *ReceiverRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *ReceiverRepository_Get_Call) Return(_a0 *receiver.Receiver, _a1 error) *ReceiverRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields:
func (_m *ReceiverRepository) List() ([]*receiver.Receiver, error) {
	ret := _m.Called()

	var r0 []*receiver.Receiver
	if rf, ok := ret.Get(0).(func() []*receiver.Receiver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*receiver.Receiver)
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

// ReceiverRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ReceiverRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
func (_e *ReceiverRepository_Expecter) List() *ReceiverRepository_List_Call {
	return &ReceiverRepository_List_Call{Call: _e.mock.On("List")}
}

func (_c *ReceiverRepository_List_Call) Run(run func()) *ReceiverRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReceiverRepository_List_Call) Return(_a0 []*receiver.Receiver, _a1 error) *ReceiverRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Migrate provides a mock function with given fields:
func (_m *ReceiverRepository) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceiverRepository_Migrate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Migrate'
type ReceiverRepository_Migrate_Call struct {
	*mock.Call
}

// Migrate is a helper method to define mock.On call
func (_e *ReceiverRepository_Expecter) Migrate() *ReceiverRepository_Migrate_Call {
	return &ReceiverRepository_Migrate_Call{Call: _e.mock.On("Migrate")}
}

func (_c *ReceiverRepository_Migrate_Call) Run(run func()) *ReceiverRepository_Migrate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReceiverRepository_Migrate_Call) Return(_a0 error) *ReceiverRepository_Migrate_Call {
	_c.Call.Return(_a0)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *ReceiverRepository) Update(_a0 *receiver.Receiver) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*receiver.Receiver) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceiverRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ReceiverRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - _a0 *receiver.Receiver
func (_e *ReceiverRepository_Expecter) Update(_a0 interface{}) *ReceiverRepository_Update_Call {
	return &ReceiverRepository_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *ReceiverRepository_Update_Call) Run(run func(_a0 *receiver.Receiver)) *ReceiverRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*receiver.Receiver))
	})
	return _c
}

func (_c *ReceiverRepository_Update_Call) Return(_a0 error) *ReceiverRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewReceiverRepository creates a new instance of ReceiverRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewReceiverRepository(t testing.TB) *ReceiverRepository {
	mock := &ReceiverRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}