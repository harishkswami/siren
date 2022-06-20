// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	template "github.com/odpf/siren/core/template"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// TemplateRepository is an autogenerated mock type for the Repository type
type TemplateRepository struct {
	mock.Mock
}

type TemplateRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TemplateRepository) EXPECT() *TemplateRepository_Expecter {
	return &TemplateRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: _a0
func (_m *TemplateRepository) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TemplateRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type TemplateRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - _a0 string
func (_e *TemplateRepository_Expecter) Delete(_a0 interface{}) *TemplateRepository_Delete_Call {
	return &TemplateRepository_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *TemplateRepository_Delete_Call) Run(run func(_a0 string)) *TemplateRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TemplateRepository_Delete_Call) Return(_a0 error) *TemplateRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetByName provides a mock function with given fields: _a0
func (_m *TemplateRepository) GetByName(_a0 string) (*template.Template, error) {
	ret := _m.Called(_a0)

	var r0 *template.Template
	if rf, ok := ret.Get(0).(func(string) *template.Template); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*template.Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateRepository_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type TemplateRepository_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//  - _a0 string
func (_e *TemplateRepository_Expecter) GetByName(_a0 interface{}) *TemplateRepository_GetByName_Call {
	return &TemplateRepository_GetByName_Call{Call: _e.mock.On("GetByName", _a0)}
}

func (_c *TemplateRepository_GetByName_Call) Run(run func(_a0 string)) *TemplateRepository_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TemplateRepository_GetByName_Call) Return(_a0 *template.Template, _a1 error) *TemplateRepository_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Index provides a mock function with given fields: _a0
func (_m *TemplateRepository) Index(_a0 string) ([]template.Template, error) {
	ret := _m.Called(_a0)

	var r0 []template.Template
	if rf, ok := ret.Get(0).(func(string) []template.Template); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]template.Template)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateRepository_Index_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Index'
type TemplateRepository_Index_Call struct {
	*mock.Call
}

// Index is a helper method to define mock.On call
//  - _a0 string
func (_e *TemplateRepository_Expecter) Index(_a0 interface{}) *TemplateRepository_Index_Call {
	return &TemplateRepository_Index_Call{Call: _e.mock.On("Index", _a0)}
}

func (_c *TemplateRepository_Index_Call) Run(run func(_a0 string)) *TemplateRepository_Index_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TemplateRepository_Index_Call) Return(_a0 []template.Template, _a1 error) *TemplateRepository_Index_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Migrate provides a mock function with given fields:
func (_m *TemplateRepository) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TemplateRepository_Migrate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Migrate'
type TemplateRepository_Migrate_Call struct {
	*mock.Call
}

// Migrate is a helper method to define mock.On call
func (_e *TemplateRepository_Expecter) Migrate() *TemplateRepository_Migrate_Call {
	return &TemplateRepository_Migrate_Call{Call: _e.mock.On("Migrate")}
}

func (_c *TemplateRepository_Migrate_Call) Run(run func()) *TemplateRepository_Migrate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TemplateRepository_Migrate_Call) Return(_a0 error) *TemplateRepository_Migrate_Call {
	_c.Call.Return(_a0)
	return _c
}

// Render provides a mock function with given fields: _a0, _a1
func (_m *TemplateRepository) Render(_a0 string, _a1 map[string]string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, map[string]string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateRepository_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type TemplateRepository_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//  - _a0 string
//  - _a1 map[string]string
func (_e *TemplateRepository_Expecter) Render(_a0 interface{}, _a1 interface{}) *TemplateRepository_Render_Call {
	return &TemplateRepository_Render_Call{Call: _e.mock.On("Render", _a0, _a1)}
}

func (_c *TemplateRepository_Render_Call) Run(run func(_a0 string, _a1 map[string]string)) *TemplateRepository_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(map[string]string))
	})
	return _c
}

func (_c *TemplateRepository_Render_Call) Return(_a0 string, _a1 error) *TemplateRepository_Render_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Upsert provides a mock function with given fields: _a0
func (_m *TemplateRepository) Upsert(_a0 *template.Template) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*template.Template) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TemplateRepository_Upsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upsert'
type TemplateRepository_Upsert_Call struct {
	*mock.Call
}

// Upsert is a helper method to define mock.On call
//  - _a0 *template.Template
func (_e *TemplateRepository_Expecter) Upsert(_a0 interface{}) *TemplateRepository_Upsert_Call {
	return &TemplateRepository_Upsert_Call{Call: _e.mock.On("Upsert", _a0)}
}

func (_c *TemplateRepository_Upsert_Call) Run(run func(_a0 *template.Template)) *TemplateRepository_Upsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*template.Template))
	})
	return _c
}

func (_c *TemplateRepository_Upsert_Call) Return(_a0 error) *TemplateRepository_Upsert_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewTemplateRepository creates a new instance of TemplateRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTemplateRepository(t testing.TB) *TemplateRepository {
	mock := &TemplateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}