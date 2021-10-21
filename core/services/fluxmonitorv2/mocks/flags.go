// Code generated by mockery v2.9.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"

	generated "github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Flags is an autogenerated mock type for the Flags type
type Flags struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *Flags) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// ContractExists provides a mock function with given fields:
func (_m *Flags) ContractExists() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsLowered provides a mock function with given fields: contractAddr
func (_m *Flags) IsLowered(contractAddr common.Address) (bool, error) {
	ret := _m.Called(contractAddr)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(contractAddr)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(contractAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseLog provides a mock function with given fields: log
func (_m *Flags) ParseLog(log types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(log)

	var r0 generated.AbigenLog
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
