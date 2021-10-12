// Code generated by mockery 2.9.0. DO NOT EDIT.

package provider

import mock "github.com/stretchr/testify/mock"

// MockProviderRepository is an autogenerated mock type for the ProviderRepository type
type MockProviderRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *MockProviderRepository) Create(_a0 *Provider) (*Provider, error) {
	ret := _m.Called(_a0)

	var r0 *Provider
	if rf, ok := ret.Get(0).(func(*Provider) *Provider); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Provider) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *MockProviderRepository) Delete(_a0 uint64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *MockProviderRepository) Get(_a0 uint64) (*Provider, error) {
	ret := _m.Called(_a0)

	var r0 *Provider
	if rf, ok := ret.Get(0).(func(uint64) *Provider); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Provider)
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

// List provides a mock function with given fields:
func (_m *MockProviderRepository) List() ([]*Provider, error) {
	ret := _m.Called()

	var r0 []*Provider
	if rf, ok := ret.Get(0).(func() []*Provider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Provider)
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

// Migrate provides a mock function with given fields:
func (_m *MockProviderRepository) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *MockProviderRepository) Update(_a0 *Provider) (*Provider, error) {
	ret := _m.Called(_a0)

	var r0 *Provider
	if rf, ok := ret.Get(0).(func(*Provider) *Provider); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Provider) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}