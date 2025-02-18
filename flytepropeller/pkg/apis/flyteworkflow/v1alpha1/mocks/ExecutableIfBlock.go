// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"
)

// ExecutableIfBlock is an autogenerated mock type for the ExecutableIfBlock type
type ExecutableIfBlock struct {
	mock.Mock
}

type ExecutableIfBlock_GetCondition struct {
	*mock.Call
}

func (_m ExecutableIfBlock_GetCondition) Return(_a0 *core.BooleanExpression) *ExecutableIfBlock_GetCondition {
	return &ExecutableIfBlock_GetCondition{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableIfBlock) OnGetCondition() *ExecutableIfBlock_GetCondition {
	c_call := _m.On("GetCondition")
	return &ExecutableIfBlock_GetCondition{Call: c_call}
}

func (_m *ExecutableIfBlock) OnGetConditionMatch(matchers ...interface{}) *ExecutableIfBlock_GetCondition {
	c_call := _m.On("GetCondition", matchers...)
	return &ExecutableIfBlock_GetCondition{Call: c_call}
}

// GetCondition provides a mock function with given fields:
func (_m *ExecutableIfBlock) GetCondition() *core.BooleanExpression {
	ret := _m.Called()

	var r0 *core.BooleanExpression
	if rf, ok := ret.Get(0).(func() *core.BooleanExpression); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BooleanExpression)
		}
	}

	return r0
}

type ExecutableIfBlock_GetThenNode struct {
	*mock.Call
}

func (_m ExecutableIfBlock_GetThenNode) Return(_a0 *string) *ExecutableIfBlock_GetThenNode {
	return &ExecutableIfBlock_GetThenNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableIfBlock) OnGetThenNode() *ExecutableIfBlock_GetThenNode {
	c_call := _m.On("GetThenNode")
	return &ExecutableIfBlock_GetThenNode{Call: c_call}
}

func (_m *ExecutableIfBlock) OnGetThenNodeMatch(matchers ...interface{}) *ExecutableIfBlock_GetThenNode {
	c_call := _m.On("GetThenNode", matchers...)
	return &ExecutableIfBlock_GetThenNode{Call: c_call}
}

// GetThenNode provides a mock function with given fields:
func (_m *ExecutableIfBlock) GetThenNode() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}
