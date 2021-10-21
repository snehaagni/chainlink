// Code generated by mockery v2.9.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	bulletprooftxmanager "github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"

	context "context"

	gas "github.com/smartcontractkit/chainlink/core/services/gas"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	uuid "github.com/satori/go.uuid"
)

// TxManager is an autogenerated mock type for the TxManager type
type TxManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TxManager) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateEthTransaction provides a mock function with given fields: db, newTx
func (_m *TxManager) CreateEthTransaction(db *gorm.DB, newTx bulletprooftxmanager.NewTx) (bulletprooftxmanager.EthTx, error) {
	ret := _m.Called(db, newTx)

	var r0 bulletprooftxmanager.EthTx
	if rf, ok := ret.Get(0).(func(*gorm.DB, bulletprooftxmanager.NewTx) bulletprooftxmanager.EthTx); ok {
		r0 = rf(db, newTx)
	} else {
		r0 = ret.Get(0).(bulletprooftxmanager.EthTx)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gorm.DB, bulletprooftxmanager.NewTx) error); ok {
		r1 = rf(db, newTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasEstimator provides a mock function with given fields:
func (_m *TxManager) GetGasEstimator() gas.Estimator {
	ret := _m.Called()

	var r0 gas.Estimator
	if rf, ok := ret.Get(0).(func() gas.Estimator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gas.Estimator)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *TxManager) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *TxManager) OnNewLongestChain(ctx context.Context, head models.Head) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *TxManager) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterResumeCallback provides a mock function with given fields: fn
func (_m *TxManager) RegisterResumeCallback(fn func(uuid.UUID, interface{}) error) {
	_m.Called(fn)
}

// Start provides a mock function with given fields:
func (_m *TxManager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trigger provides a mock function with given fields: addr
func (_m *TxManager) Trigger(addr common.Address) {
	_m.Called(addr)
}
