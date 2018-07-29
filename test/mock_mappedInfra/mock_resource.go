// Code generated by MockGen. DO NOT EDIT.
// Source: mappedInfra/resource.go

// Package mock_mappedInfra is a generated GoMock package.
package mock_mappedInfra

import (
	gomock "github.com/golang/mock/gomock"
	aws "github.com/thomas.obenaus/inframapper/aws"
	mappedInfra "github.com/thomas.obenaus/inframapper/mappedInfra"
	terraform "github.com/thomas.obenaus/inframapper/terraform"
	reflect "reflect"
)

// MockMappedResource is a mock of MappedResource interface
type MockMappedResource struct {
	ctrl     *gomock.Controller
	recorder *MockMappedResourceMockRecorder
}

// MockMappedResourceMockRecorder is the mock recorder for MockMappedResource
type MockMappedResourceMockRecorder struct {
	mock *MockMappedResource
}

// NewMockMappedResource creates a new mock instance
func NewMockMappedResource(ctrl *gomock.Controller) *MockMappedResource {
	mock := &MockMappedResource{ctrl: ctrl}
	mock.recorder = &MockMappedResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMappedResource) EXPECT() *MockMappedResourceMockRecorder {
	return m.recorder
}

// Aws mocks base method
func (m *MockMappedResource) Aws() aws.Resource {
	ret := m.ctrl.Call(m, "Aws")
	ret0, _ := ret[0].(aws.Resource)
	return ret0
}

// Aws indicates an expected call of Aws
func (mr *MockMappedResourceMockRecorder) Aws() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aws", reflect.TypeOf((*MockMappedResource)(nil).Aws))
}

// HasAws mocks base method
func (m *MockMappedResource) HasAws() bool {
	ret := m.ctrl.Call(m, "HasAws")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasAws indicates an expected call of HasAws
func (mr *MockMappedResourceMockRecorder) HasAws() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAws", reflect.TypeOf((*MockMappedResource)(nil).HasAws))
}

// HasTerraform mocks base method
func (m *MockMappedResource) HasTerraform() bool {
	ret := m.ctrl.Call(m, "HasTerraform")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasTerraform indicates an expected call of HasTerraform
func (mr *MockMappedResourceMockRecorder) HasTerraform() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTerraform", reflect.TypeOf((*MockMappedResource)(nil).HasTerraform))
}

// Terraform mocks base method
func (m *MockMappedResource) Terraform() terraform.Resource {
	ret := m.ctrl.Call(m, "Terraform")
	ret0, _ := ret[0].(terraform.Resource)
	return ret0
}

// Terraform indicates an expected call of Terraform
func (mr *MockMappedResourceMockRecorder) Terraform() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terraform", reflect.TypeOf((*MockMappedResource)(nil).Terraform))
}

// ResourceType mocks base method
func (m *MockMappedResource) ResourceType() mappedInfra.Type {
	ret := m.ctrl.Call(m, "ResourceType")
	ret0, _ := ret[0].(mappedInfra.Type)
	return ret0
}

// ResourceType indicates an expected call of ResourceType
func (mr *MockMappedResourceMockRecorder) ResourceType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceType", reflect.TypeOf((*MockMappedResource)(nil).ResourceType))
}

// String mocks base method
func (m *MockMappedResource) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockMappedResourceMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockMappedResource)(nil).String))
}