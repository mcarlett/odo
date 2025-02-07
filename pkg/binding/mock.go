// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/binding/interface.go

// Package binding is a generated GoMock package.
package binding

import (
	reflect "reflect"

	parser "github.com/devfile/library/pkg/devfile/parser"
	gomock "github.com/golang/mock/gomock"
	api "github.com/redhat-developer/odo/pkg/api"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddBinding mocks base method.
func (m *MockClient) AddBinding(flags map[string]string, bindingName string, bindAsFiles bool, unstructuredService unstructured.Unstructured, workloadName string, workloadGVK schema.GroupVersionKind) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinding", flags, bindingName, bindAsFiles, unstructuredService, workloadName, workloadGVK)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinding indicates an expected call of AddBinding.
func (mr *MockClientMockRecorder) AddBinding(flags, bindingName, bindAsFiles, unstructuredService, workloadName, workloadGVK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinding", reflect.TypeOf((*MockClient)(nil).AddBinding), flags, bindingName, bindAsFiles, unstructuredService, workloadName, workloadGVK)
}

// AddBindingToDevfile mocks base method.
func (m *MockClient) AddBindingToDevfile(bindingName string, bindAsFiles bool, unstructuredService unstructured.Unstructured, obj parser.DevfileObj) (parser.DevfileObj, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBindingToDevfile", bindingName, bindAsFiles, unstructuredService, obj)
	ret0, _ := ret[0].(parser.DevfileObj)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBindingToDevfile indicates an expected call of AddBindingToDevfile.
func (mr *MockClientMockRecorder) AddBindingToDevfile(bindingName, bindAsFiles, unstructuredService, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBindingToDevfile", reflect.TypeOf((*MockClient)(nil).AddBindingToDevfile), bindingName, bindAsFiles, unstructuredService, obj)
}

// AskBindAsFiles mocks base method.
func (m *MockClient) AskBindAsFiles(flags map[string]string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskBindAsFiles", flags)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskBindAsFiles indicates an expected call of AskBindAsFiles.
func (mr *MockClientMockRecorder) AskBindAsFiles(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskBindAsFiles", reflect.TypeOf((*MockClient)(nil).AskBindAsFiles), flags)
}

// AskBindingName mocks base method.
func (m *MockClient) AskBindingName(serviceName, componentName string, flags map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskBindingName", serviceName, componentName, flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskBindingName indicates an expected call of AskBindingName.
func (mr *MockClientMockRecorder) AskBindingName(serviceName, componentName, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskBindingName", reflect.TypeOf((*MockClient)(nil).AskBindingName), serviceName, componentName, flags)
}

// GetBindingFromCluster mocks base method.
func (m *MockClient) GetBindingFromCluster(name string) (api.ServiceBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBindingFromCluster", name)
	ret0, _ := ret[0].(api.ServiceBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBindingFromCluster indicates an expected call of GetBindingFromCluster.
func (mr *MockClientMockRecorder) GetBindingFromCluster(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBindingFromCluster", reflect.TypeOf((*MockClient)(nil).GetBindingFromCluster), name)
}

// GetBindingsFromDevfile mocks base method.
func (m *MockClient) GetBindingsFromDevfile(devfileObj parser.DevfileObj, context string) ([]api.ServiceBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBindingsFromDevfile", devfileObj, context)
	ret0, _ := ret[0].([]api.ServiceBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBindingsFromDevfile indicates an expected call of GetBindingsFromDevfile.
func (mr *MockClientMockRecorder) GetBindingsFromDevfile(devfileObj, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBindingsFromDevfile", reflect.TypeOf((*MockClient)(nil).GetBindingsFromDevfile), devfileObj, context)
}

// GetFlags mocks base method.
func (m *MockClient) GetFlags(flags map[string]string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlags", flags)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetFlags indicates an expected call of GetFlags.
func (mr *MockClientMockRecorder) GetFlags(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlags", reflect.TypeOf((*MockClient)(nil).GetFlags), flags)
}

// GetServiceInstances mocks base method.
func (m *MockClient) GetServiceInstances() (map[string]unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceInstances")
	ret0, _ := ret[0].(map[string]unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceInstances indicates an expected call of GetServiceInstances.
func (mr *MockClientMockRecorder) GetServiceInstances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceInstances", reflect.TypeOf((*MockClient)(nil).GetServiceInstances))
}

// ListAllBindings mocks base method.
func (m *MockClient) ListAllBindings(devfileObj parser.DevfileObj, context string) ([]api.ServiceBinding, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllBindings", devfileObj, context)
	ret0, _ := ret[0].([]api.ServiceBinding)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAllBindings indicates an expected call of ListAllBindings.
func (mr *MockClientMockRecorder) ListAllBindings(devfileObj, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllBindings", reflect.TypeOf((*MockClient)(nil).ListAllBindings), devfileObj, context)
}

// RemoveBinding mocks base method.
func (m *MockClient) RemoveBinding(bindingName string, obj parser.DevfileObj) (parser.DevfileObj, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBinding", bindingName, obj)
	ret0, _ := ret[0].(parser.DevfileObj)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveBinding indicates an expected call of RemoveBinding.
func (mr *MockClientMockRecorder) RemoveBinding(bindingName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBinding", reflect.TypeOf((*MockClient)(nil).RemoveBinding), bindingName, obj)
}

// SelectServiceInstance mocks base method.
func (m *MockClient) SelectServiceInstance(flags map[string]string, serviceMap map[string]unstructured.Unstructured) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectServiceInstance", flags, serviceMap)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectServiceInstance indicates an expected call of SelectServiceInstance.
func (mr *MockClientMockRecorder) SelectServiceInstance(flags, serviceMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectServiceInstance", reflect.TypeOf((*MockClient)(nil).SelectServiceInstance), flags, serviceMap)
}

// SelectWorkloadInstance mocks base method.
func (m *MockClient) SelectWorkloadInstance(flags map[string]string) (string, schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectWorkloadInstance", flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(schema.GroupVersionKind)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SelectWorkloadInstance indicates an expected call of SelectWorkloadInstance.
func (mr *MockClientMockRecorder) SelectWorkloadInstance(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectWorkloadInstance", reflect.TypeOf((*MockClient)(nil).SelectWorkloadInstance), flags)
}

// ValidateAddBinding mocks base method.
func (m *MockClient) ValidateAddBinding(flags map[string]string, withDevfile bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAddBinding", flags, withDevfile)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateAddBinding indicates an expected call of ValidateAddBinding.
func (mr *MockClientMockRecorder) ValidateAddBinding(flags, withDevfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAddBinding", reflect.TypeOf((*MockClient)(nil).ValidateAddBinding), flags, withDevfile)
}

// ValidateRemoveBinding mocks base method.
func (m *MockClient) ValidateRemoveBinding(flags map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRemoveBinding", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRemoveBinding indicates an expected call of ValidateRemoveBinding.
func (mr *MockClientMockRecorder) ValidateRemoveBinding(flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRemoveBinding", reflect.TypeOf((*MockClient)(nil).ValidateRemoveBinding), flags)
}
