// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/goquorum/quorum-plugin-hashicorp-vault-account-manager/internal/vault/hashicorp (interfaces: AccountCreator)

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	"crypto/ecdsa"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/golang/mock/gomock"
	"github.com/goquorum/quorum-plugin-hashicorp-vault-account-manager/internal/config"
)

// MockAccountCreator is a mock of AccountCreator interface
type MockAccountCreator struct {
	ctrl     *gomock.Controller
	recorder *MockAccountCreatorMockRecorder
}

// MockAccountCreatorMockRecorder is the mock recorder for MockAccountCreator
type MockAccountCreatorMockRecorder struct {
	mock *MockAccountCreator
}

// NewMockAccountCreator creates a new mock instance
func NewMockAccountCreator(ctrl *gomock.Controller) *MockAccountCreator {
	mock := &MockAccountCreator{ctrl: ctrl}
	mock.recorder = &MockAccountCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountCreator) EXPECT() *MockAccountCreatorMockRecorder {
	return m.recorder
}

// ImportECDSA mocks base method
func (m *MockAccountCreator) ImportECDSA(arg0 *ecdsa.PrivateKey, arg1 config.VaultSecretConfig) (accounts.Account, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportECDSA", arg0, arg1)
	ret0, _ := ret[0].(accounts.Account)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportECDSA indicates an expected call of ImportECDSA
func (mr *MockAccountCreatorMockRecorder) ImportECDSA(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportECDSA", reflect.TypeOf((*MockAccountCreator)(nil).ImportECDSA), arg0, arg1)
}

// NewAccount mocks base method
func (m *MockAccountCreator) NewAccount(arg0 config.VaultSecretConfig) (accounts.Account, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0)
	ret0, _ := ret[0].(accounts.Account)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewAccount indicates an expected call of NewAccount
func (mr *MockAccountCreatorMockRecorder) NewAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountCreator)(nil).NewAccount), arg0)
}
