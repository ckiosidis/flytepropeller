// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	executors "github.com/lyft/flytepropeller/pkg/controller/executors"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

type Node_AbortHandler struct {
	*mock.Call
}

func (_m Node_AbortHandler) Return(_a0 error) *Node_AbortHandler {
	return &Node_AbortHandler{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnAbortHandler(ctx context.Context, w v1alpha1.ExecutableWorkflow, currentNode v1alpha1.ExecutableNode, reason string) *Node_AbortHandler {
	c := _m.On("AbortHandler", ctx, w, currentNode, reason)
	return &Node_AbortHandler{Call: c}
}

func (_m *Node) OnAbortHandlerMatch(matchers ...interface{}) *Node_AbortHandler {
	c := _m.On("AbortHandler", matchers...)
	return &Node_AbortHandler{Call: c}
}

// AbortHandler provides a mock function with given fields: ctx, w, currentNode, reason
func (_m *Node) AbortHandler(ctx context.Context, w v1alpha1.ExecutableWorkflow, currentNode v1alpha1.ExecutableNode, reason string) error {
	ret := _m.Called(ctx, w, currentNode, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1alpha1.ExecutableWorkflow, v1alpha1.ExecutableNode, string) error); ok {
		r0 = rf(ctx, w, currentNode, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Node_FinalizeHandler struct {
	*mock.Call
}

func (_m Node_FinalizeHandler) Return(_a0 error) *Node_FinalizeHandler {
	return &Node_FinalizeHandler{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnFinalizeHandler(ctx context.Context, w v1alpha1.ExecutableWorkflow, currentNode v1alpha1.ExecutableNode) *Node_FinalizeHandler {
	c := _m.On("FinalizeHandler", ctx, w, currentNode)
	return &Node_FinalizeHandler{Call: c}
}

func (_m *Node) OnFinalizeHandlerMatch(matchers ...interface{}) *Node_FinalizeHandler {
	c := _m.On("FinalizeHandler", matchers...)
	return &Node_FinalizeHandler{Call: c}
}

// FinalizeHandler provides a mock function with given fields: ctx, w, currentNode
func (_m *Node) FinalizeHandler(ctx context.Context, w v1alpha1.ExecutableWorkflow, currentNode v1alpha1.ExecutableNode) error {
	ret := _m.Called(ctx, w, currentNode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1alpha1.ExecutableWorkflow, v1alpha1.ExecutableNode) error); ok {
		r0 = rf(ctx, w, currentNode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Node_Initialize struct {
	*mock.Call
}

func (_m Node_Initialize) Return(_a0 error) *Node_Initialize {
	return &Node_Initialize{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnInitialize(ctx context.Context) *Node_Initialize {
	c := _m.On("Initialize", ctx)
	return &Node_Initialize{Call: c}
}

func (_m *Node) OnInitializeMatch(matchers ...interface{}) *Node_Initialize {
	c := _m.On("Initialize", matchers...)
	return &Node_Initialize{Call: c}
}

// Initialize provides a mock function with given fields: ctx
func (_m *Node) Initialize(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Node_RecursiveNodeHandler struct {
	*mock.Call
}

func (_m Node_RecursiveNodeHandler) Return(_a0 executors.NodeStatus, _a1 error) *Node_RecursiveNodeHandler {
	return &Node_RecursiveNodeHandler{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Node) OnRecursiveNodeHandler(ctx context.Context, w v1alpha1.ExecutableWorkflow, currentNode v1alpha1.ExecutableNode) *Node_RecursiveNodeHandler {
	c := _m.On("RecursiveNodeHandler", ctx, w, currentNode)
	return &Node_RecursiveNodeHandler{Call: c}
}

func (_m *Node) OnRecursiveNodeHandlerMatch(matchers ...interface{}) *Node_RecursiveNodeHandler {
	c := _m.On("RecursiveNodeHandler", matchers...)
	return &Node_RecursiveNodeHandler{Call: c}
}

// RecursiveNodeHandler provides a mock function with given fields: ctx, w, currentNode
func (_m *Node) RecursiveNodeHandler(ctx context.Context, w v1alpha1.ExecutableWorkflow, currentNode v1alpha1.ExecutableNode) (executors.NodeStatus, error) {
	ret := _m.Called(ctx, w, currentNode)

	var r0 executors.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, v1alpha1.ExecutableWorkflow, v1alpha1.ExecutableNode) executors.NodeStatus); ok {
		r0 = rf(ctx, w, currentNode)
	} else {
		r0 = ret.Get(0).(executors.NodeStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1alpha1.ExecutableWorkflow, v1alpha1.ExecutableNode) error); ok {
		r1 = rf(ctx, w, currentNode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Node_SetInputsForStartNode struct {
	*mock.Call
}

func (_m Node_SetInputsForStartNode) Return(_a0 executors.NodeStatus, _a1 error) *Node_SetInputsForStartNode {
	return &Node_SetInputsForStartNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Node) OnSetInputsForStartNode(ctx context.Context, w v1alpha1.BaseWorkflowWithStatus, inputs *core.LiteralMap) *Node_SetInputsForStartNode {
	c := _m.On("SetInputsForStartNode", ctx, w, inputs)
	return &Node_SetInputsForStartNode{Call: c}
}

func (_m *Node) OnSetInputsForStartNodeMatch(matchers ...interface{}) *Node_SetInputsForStartNode {
	c := _m.On("SetInputsForStartNode", matchers...)
	return &Node_SetInputsForStartNode{Call: c}
}

// SetInputsForStartNode provides a mock function with given fields: ctx, w, inputs
func (_m *Node) SetInputsForStartNode(ctx context.Context, w v1alpha1.BaseWorkflowWithStatus, inputs *core.LiteralMap) (executors.NodeStatus, error) {
	ret := _m.Called(ctx, w, inputs)

	var r0 executors.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, v1alpha1.BaseWorkflowWithStatus, *core.LiteralMap) executors.NodeStatus); ok {
		r0 = rf(ctx, w, inputs)
	} else {
		r0 = ret.Get(0).(executors.NodeStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1alpha1.BaseWorkflowWithStatus, *core.LiteralMap) error); ok {
		r1 = rf(ctx, w, inputs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
