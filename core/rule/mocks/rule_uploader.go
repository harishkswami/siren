// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	rule "github.com/odpf/siren/core/rule"
	mock "github.com/stretchr/testify/mock"

	template "github.com/odpf/siren/core/template"
)

// RuleUploader is an autogenerated mock type for the RuleUploader type
type RuleUploader struct {
	mock.Mock
}

type RuleUploader_Expecter struct {
	mock *mock.Mock
}

func (_m *RuleUploader) EXPECT() *RuleUploader_Expecter {
	return &RuleUploader_Expecter{mock: &_m.Mock}
}

// UpsertRule provides a mock function with given fields: ctx, rl, templateToUpdate, namespaceURN
func (_m *RuleUploader) UpsertRule(ctx context.Context, rl *rule.Rule, templateToUpdate *template.Template, namespaceURN string) error {
	ret := _m.Called(ctx, rl, templateToUpdate, namespaceURN)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *rule.Rule, *template.Template, string) error); ok {
		r0 = rf(ctx, rl, templateToUpdate, namespaceURN)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RuleUploader_UpsertRule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertRule'
type RuleUploader_UpsertRule_Call struct {
	*mock.Call
}

// UpsertRule is a helper method to define mock.On call
//  - ctx context.Context
//  - rl *rule.Rule
//  - templateToUpdate *template.Template
//  - namespaceURN string
func (_e *RuleUploader_Expecter) UpsertRule(ctx interface{}, rl interface{}, templateToUpdate interface{}, namespaceURN interface{}) *RuleUploader_UpsertRule_Call {
	return &RuleUploader_UpsertRule_Call{Call: _e.mock.On("UpsertRule", ctx, rl, templateToUpdate, namespaceURN)}
}

func (_c *RuleUploader_UpsertRule_Call) Run(run func(ctx context.Context, rl *rule.Rule, templateToUpdate *template.Template, namespaceURN string)) *RuleUploader_UpsertRule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*rule.Rule), args[2].(*template.Template), args[3].(string))
	})
	return _c
}

func (_c *RuleUploader_UpsertRule_Call) Return(_a0 error) *RuleUploader_UpsertRule_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewRuleUploader interface {
	mock.TestingT
	Cleanup(func())
}

// NewRuleUploader creates a new instance of RuleUploader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRuleUploader(t mockConstructorTestingTNewRuleUploader) *RuleUploader {
	mock := &RuleUploader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
