// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/middleware/presentproof (interfaces: Provider,Metadata)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/Universal-Health-Chain/aries-framework-go/pkg/didcomm/common/service"
	presentproof "github.com/Universal-Health-Chain/aries-framework-go/pkg/didcomm/protocol/presentproof"
	vdr "github.com/Universal-Health-Chain/aries-framework-go/pkg/framework/aries/api/vdr"
	verifiable "github.com/Universal-Health-Chain/aries-framework-go/pkg/store/verifiable"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// VDRegistry mocks base method
func (m *MockProvider) VDRegistry() vdr.Registry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VDRegistry")
	ret0, _ := ret[0].(vdr.Registry)
	return ret0
}

// VDRegistry indicates an expected call of VDRegistry
func (mr *MockProviderMockRecorder) VDRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VDRegistry", reflect.TypeOf((*MockProvider)(nil).VDRegistry))
}

// VerifiableStore mocks base method
func (m *MockProvider) VerifiableStore() verifiable.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifiableStore")
	ret0, _ := ret[0].(verifiable.Store)
	return ret0
}

// VerifiableStore indicates an expected call of VerifiableStore
func (mr *MockProviderMockRecorder) VerifiableStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifiableStore", reflect.TypeOf((*MockProvider)(nil).VerifiableStore))
}

// MockMetadata is a mock of Metadata interface
type MockMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataMockRecorder
}

// MockMetadataMockRecorder is the mock recorder for MockMetadata
type MockMetadataMockRecorder struct {
	mock *MockMetadata
}

// NewMockMetadata creates a new mock instance
func NewMockMetadata(ctrl *gomock.Controller) *MockMetadata {
	mock := &MockMetadata{ctrl: ctrl}
	mock.recorder = &MockMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetadata) EXPECT() *MockMetadataMockRecorder {
	return m.recorder
}

// Message mocks base method
func (m *MockMetadata) Message() service.DIDCommMsg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(service.DIDCommMsg)
	return ret0
}

// Message indicates an expected call of Message
func (mr *MockMetadataMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockMetadata)(nil).Message))
}

// Presentation mocks base method
func (m *MockMetadata) Presentation() *presentproof.Presentation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Presentation")
	ret0, _ := ret[0].(*presentproof.Presentation)
	return ret0
}

// Presentation indicates an expected call of Presentation
func (mr *MockMetadataMockRecorder) Presentation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Presentation", reflect.TypeOf((*MockMetadata)(nil).Presentation))
}

// PresentationNames mocks base method
func (m *MockMetadata) PresentationNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PresentationNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// PresentationNames indicates an expected call of PresentationNames
func (mr *MockMetadataMockRecorder) PresentationNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentationNames", reflect.TypeOf((*MockMetadata)(nil).PresentationNames))
}

// Properties mocks base method
func (m *MockMetadata) Properties() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Properties")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Properties indicates an expected call of Properties
func (mr *MockMetadataMockRecorder) Properties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockMetadata)(nil).Properties))
}

// ProposePresentation mocks base method
func (m *MockMetadata) ProposePresentation() *presentproof.ProposePresentation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposePresentation")
	ret0, _ := ret[0].(*presentproof.ProposePresentation)
	return ret0
}

// ProposePresentation indicates an expected call of ProposePresentation
func (mr *MockMetadataMockRecorder) ProposePresentation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposePresentation", reflect.TypeOf((*MockMetadata)(nil).ProposePresentation))
}

// RequestPresentation mocks base method
func (m *MockMetadata) RequestPresentation() *presentproof.RequestPresentation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestPresentation")
	ret0, _ := ret[0].(*presentproof.RequestPresentation)
	return ret0
}

// RequestPresentation indicates an expected call of RequestPresentation
func (mr *MockMetadataMockRecorder) RequestPresentation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestPresentation", reflect.TypeOf((*MockMetadata)(nil).RequestPresentation))
}

// StateName mocks base method
func (m *MockMetadata) StateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// StateName indicates an expected call of StateName
func (mr *MockMetadataMockRecorder) StateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateName", reflect.TypeOf((*MockMetadata)(nil).StateName))
}
