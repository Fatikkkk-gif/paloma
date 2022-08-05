// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/palomachain/paloma/x/consensus/types"
)

// QueuedSignedMessageI is an autogenerated mock type for the QueuedSignedMessageI type
type QueuedSignedMessageI struct {
	mock.Mock
}

// AddEvidence provides a mock function with given fields: _a0
func (_m *QueuedSignedMessageI) AddEvidence(_a0 types.Evidence) {
	_m.Called(_a0)
}

// AddSignData provides a mock function with given fields: _a0
func (_m *QueuedSignedMessageI) AddSignData(_a0 *types.SignData) {
	_m.Called(_a0)
}

// ConsensusMsg provides a mock function with given fields: _a0
func (_m *QueuedSignedMessageI) ConsensusMsg(_a0 codectypes.AnyUnpacker) (types.ConsensusMsg, error) {
	ret := _m.Called(_a0)

	var r0 types.ConsensusMsg
	if rf, ok := ret.Get(0).(func(codectypes.AnyUnpacker) types.ConsensusMsg); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ConsensusMsg)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(codectypes.AnyUnpacker) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddedAt provides a mock function with given fields:
func (_m *QueuedSignedMessageI) GetAddedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// GetAddedAtBlockHeight provides a mock function with given fields:
func (_m *QueuedSignedMessageI) GetAddedAtBlockHeight() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetBytesToSign provides a mock function with given fields:
func (_m *QueuedSignedMessageI) GetBytesToSign() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// GetEvidence provides a mock function with given fields:
func (_m *QueuedSignedMessageI) GetEvidence() []*types.Evidence {
	ret := _m.Called()

	var r0 []*types.Evidence
	if rf, ok := ret.Get(0).(func() []*types.Evidence); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Evidence)
		}
	}

	return r0
}

// GetId provides a mock function with given fields:
func (_m *QueuedSignedMessageI) GetId() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetPublicAccessData provides a mock function with given fields:
func (_m *QueuedSignedMessageI) GetPublicAccessData() *types.PublicAccessData {
	ret := _m.Called()

	var r0 *types.PublicAccessData
	if rf, ok := ret.Get(0).(func() *types.PublicAccessData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.PublicAccessData)
		}
	}

	return r0
}

// GetSignData provides a mock function with given fields:
func (_m *QueuedSignedMessageI) GetSignData() []*types.SignData {
	ret := _m.Called()

	var r0 []*types.SignData
	if rf, ok := ret.Get(0).(func() []*types.SignData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.SignData)
		}
	}

	return r0
}

// Nonce provides a mock function with given fields:
func (_m *QueuedSignedMessageI) Nonce() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// ProtoMessage provides a mock function with given fields:
func (_m *QueuedSignedMessageI) ProtoMessage() {
	_m.Called()
}

// Reset provides a mock function with given fields:
func (_m *QueuedSignedMessageI) Reset() {
	_m.Called()
}

// SetPublicAccessData provides a mock function with given fields: _a0
func (_m *QueuedSignedMessageI) SetPublicAccessData(_a0 *types.PublicAccessData) {
	_m.Called(_a0)
}

// String provides a mock function with given fields:
func (_m *QueuedSignedMessageI) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewQueuedSignedMessageI interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueuedSignedMessageI creates a new instance of QueuedSignedMessageI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueuedSignedMessageI(t mockConstructorTestingTNewQueuedSignedMessageI) *QueuedSignedMessageI {
	mock := &QueuedSignedMessageI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
