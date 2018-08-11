// Code generated by MockGen. DO NOT EDIT.
// Source: aws/iface/aws_IF.go

// Package mock_iface is a generated GoMock package.
package mock_iface

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEC2IF is a mock of EC2IF interface
type MockEC2IF struct {
	ctrl     *gomock.Controller
	recorder *MockEC2IFMockRecorder
}

// MockEC2IFMockRecorder is the mock recorder for MockEC2IF
type MockEC2IFMockRecorder struct {
	mock *MockEC2IF
}

// NewMockEC2IF creates a new mock instance
func NewMockEC2IF(ctrl *gomock.Controller) *MockEC2IF {
	mock := &MockEC2IF{ctrl: ctrl}
	mock.recorder = &MockEC2IFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2IF) EXPECT() *MockEC2IFMockRecorder {
	return m.recorder
}

// DescribeVpcs mocks base method
func (m *MockEC2IF) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	ret := m.ctrl.Call(m, "DescribeVpcs", input)
	ret0, _ := ret[0].(*ec2.DescribeVpcsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcs indicates an expected call of DescribeVpcs
func (mr *MockEC2IFMockRecorder) DescribeVpcs(input interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcs", reflect.TypeOf((*MockEC2IF)(nil).DescribeVpcs), input)
}
