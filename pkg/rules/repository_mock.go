// Code generated by mockery v2.6.0. DO NOT EDIT.

package rules

import (
	"github.com/odpf/siren/domain"
	mock "github.com/stretchr/testify/mock"
)

// RuleRepository is an autogenerated mock type for the RuleRepository type
type RuleRepositoryMock struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *RuleRepositoryMock) Get(_a0 string, _a1 string, _a2 string, _a3 string, _a4 string) ([]Rule, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 []Rule
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) []Rule); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields:
func (_m *RuleRepositoryMock) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upsert provides a mock function with given fields: _a0, _a1
func (_m *RuleRepositoryMock) Upsert(_a0 *Rule, _a1 cortexCaller, service domain.TemplatesService) (*Rule, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Rule
	if rf, ok := ret.Get(0).(func(*Rule, cortexCaller) *Rule); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Rule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Rule, cortexCaller) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}