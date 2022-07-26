// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	static "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/static"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// WorkflowStaticObjectStore is an autogenerated mock type for the WorkflowStaticObjectStore type
type WorkflowStaticObjectStore struct {
	mock.Mock
}

type WorkflowStaticObjectStore_Get struct {
	*mock.Call
}

func (_m WorkflowStaticObjectStore_Get) Return(_a0 *static.WorkflowStaticExecutionObj, _a1 error) *WorkflowStaticObjectStore_Get {
	return &WorkflowStaticObjectStore_Get{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *WorkflowStaticObjectStore) OnGet(ctx context.Context, wf *v1alpha1.FlyteWorkflow) *WorkflowStaticObjectStore_Get {
	c_call := _m.On("Get", ctx, wf)
	return &WorkflowStaticObjectStore_Get{Call: c_call}
}

func (_m *WorkflowStaticObjectStore) OnGetMatch(matchers ...interface{}) *WorkflowStaticObjectStore_Get {
	c_call := _m.On("Get", matchers...)
	return &WorkflowStaticObjectStore_Get{Call: c_call}
}

// Get provides a mock function with given fields: ctx, wf
func (_m *WorkflowStaticObjectStore) Get(ctx context.Context, wf *v1alpha1.FlyteWorkflow) (*static.WorkflowStaticExecutionObj, error) {
	ret := _m.Called(ctx, wf)

	var r0 *static.WorkflowStaticExecutionObj
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.FlyteWorkflow) *static.WorkflowStaticExecutionObj); ok {
		r0 = rf(ctx, wf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*static.WorkflowStaticExecutionObj)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.FlyteWorkflow) error); ok {
		r1 = rf(ctx, wf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, wf
func (_m *WorkflowStaticObjectStore) Remove(ctx context.Context, wf *v1alpha1.FlyteWorkflow) {
	_m.Called(ctx, wf)
}
