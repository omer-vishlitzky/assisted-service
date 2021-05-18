// Code generated by MockGen. DO NOT EDIT.
// Source: dns.go

// Package dns is a generated GoMock package.
package dns

import (
	context "context"
	dnsproviders "github.com/danielerez/go-dns-client/pkg/dnsproviders"
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	reflect "reflect"
)

// MockDNSApi is a mock of DNSApi interface
type MockDNSApi struct {
	ctrl     *gomock.Controller
	recorder *MockDNSApiMockRecorder
}

// MockDNSApiMockRecorder is the mock recorder for MockDNSApi
type MockDNSApiMockRecorder struct {
	mock *MockDNSApi
}

// NewMockDNSApi creates a new mock instance
func NewMockDNSApi(ctrl *gomock.Controller) *MockDNSApi {
	mock := &MockDNSApi{ctrl: ctrl}
	mock.recorder = &MockDNSApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDNSApi) EXPECT() *MockDNSApiMockRecorder {
	return m.recorder
}

// CreateDNSRecordSets mocks base method
func (m *MockDNSApi) CreateDNSRecordSets(ctx context.Context, cluster *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDNSRecordSets", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDNSRecordSets indicates an expected call of CreateDNSRecordSets
func (mr *MockDNSApiMockRecorder) CreateDNSRecordSets(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDNSRecordSets", reflect.TypeOf((*MockDNSApi)(nil).CreateDNSRecordSets), ctx, cluster)
}

// DeleteDNSRecordSets mocks base method
func (m *MockDNSApi) DeleteDNSRecordSets(ctx context.Context, cluster *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDNSRecordSets", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDNSRecordSets indicates an expected call of DeleteDNSRecordSets
func (mr *MockDNSApiMockRecorder) DeleteDNSRecordSets(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDNSRecordSets", reflect.TypeOf((*MockDNSApi)(nil).DeleteDNSRecordSets), ctx, cluster)
}

// GetDNSDomain mocks base method
func (m *MockDNSApi) GetDNSDomain(clusterName, baseDNSDomainName string) (*DNSDomain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDNSDomain", clusterName, baseDNSDomainName)
	ret0, _ := ret[0].(*DNSDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDNSDomain indicates an expected call of GetDNSDomain
func (mr *MockDNSApiMockRecorder) GetDNSDomain(clusterName, baseDNSDomainName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDNSDomain", reflect.TypeOf((*MockDNSApi)(nil).GetDNSDomain), clusterName, baseDNSDomainName)
}

// ValidateDNSName mocks base method
func (m *MockDNSApi) ValidateDNSName(clusterName, baseDNSDomainName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDNSName", clusterName, baseDNSDomainName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDNSName indicates an expected call of ValidateDNSName
func (mr *MockDNSApiMockRecorder) ValidateDNSName(clusterName, baseDNSDomainName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDNSName", reflect.TypeOf((*MockDNSApi)(nil).ValidateDNSName), clusterName, baseDNSDomainName)
}

// ValidateBaseDNS mocks base method
func (m *MockDNSApi) ValidateBaseDNS(domain *DNSDomain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateBaseDNS", domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateBaseDNS indicates an expected call of ValidateBaseDNS
func (mr *MockDNSApiMockRecorder) ValidateBaseDNS(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateBaseDNS", reflect.TypeOf((*MockDNSApi)(nil).ValidateBaseDNS), domain)
}

// ValidateDNSRecords mocks base method
func (m *MockDNSApi) ValidateDNSRecords(cluster common.Cluster, domain *DNSDomain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDNSRecords", cluster, domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDNSRecords indicates an expected call of ValidateDNSRecords
func (mr *MockDNSApiMockRecorder) ValidateDNSRecords(cluster, domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDNSRecords", reflect.TypeOf((*MockDNSApi)(nil).ValidateDNSRecords), cluster, domain)
}

// MockDNSProviderFactory is a mock of DNSProviderFactory interface
type MockDNSProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDNSProviderFactoryMockRecorder
}

// MockDNSProviderFactoryMockRecorder is the mock recorder for MockDNSProviderFactory
type MockDNSProviderFactoryMockRecorder struct {
	mock *MockDNSProviderFactory
}

// NewMockDNSProviderFactory creates a new mock instance
func NewMockDNSProviderFactory(ctrl *gomock.Controller) *MockDNSProviderFactory {
	mock := &MockDNSProviderFactory{ctrl: ctrl}
	mock.recorder = &MockDNSProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDNSProviderFactory) EXPECT() *MockDNSProviderFactoryMockRecorder {
	return m.recorder
}

// GetProviderByRecordType mocks base method
func (m *MockDNSProviderFactory) GetProviderByRecordType(domain *DNSDomain, recordType string) dnsproviders.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProviderByRecordType", domain, recordType)
	ret0, _ := ret[0].(dnsproviders.Provider)
	return ret0
}

// GetProviderByRecordType indicates an expected call of GetProviderByRecordType
func (mr *MockDNSProviderFactoryMockRecorder) GetProviderByRecordType(domain, recordType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProviderByRecordType", reflect.TypeOf((*MockDNSProviderFactory)(nil).GetProviderByRecordType), domain, recordType)
}

// GetProvider mocks base method
func (m *MockDNSProviderFactory) GetProvider(domain *DNSDomain) dnsproviders.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProvider", domain)
	ret0, _ := ret[0].(dnsproviders.Provider)
	return ret0
}

// GetProvider indicates an expected call of GetProvider
func (mr *MockDNSProviderFactoryMockRecorder) GetProvider(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProvider", reflect.TypeOf((*MockDNSProviderFactory)(nil).GetProvider), domain)
}
