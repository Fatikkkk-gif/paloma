// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	log "github.com/cometbft/cometbft/libs/log"

	mock "github.com/stretchr/testify/mock"

	schedulertypes "github.com/palomachain/paloma/x/scheduler/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// SchedulerKeeper is an autogenerated mock type for the SchedulerKeeper type
type SchedulerKeeper struct {
	mock.Mock
}

// AddNewJob provides a mock function with given fields: ctx, job
func (_m *SchedulerKeeper) AddNewJob(ctx types.Context, job *schedulertypes.Job) (types.AccAddress, error) {
	ret := _m.Called(ctx, job)

	var r0 types.AccAddress
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *schedulertypes.Job) (types.AccAddress, error)); ok {
		return rf(ctx, job)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *schedulertypes.Job) types.AccAddress); ok {
		r0 = rf(ctx, job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AccAddress)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *schedulertypes.Job) error); ok {
		r1 = rf(ctx, job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: ctx, addr
func (_m *SchedulerKeeper) GetAccount(ctx types.Context, addr types.AccAddress) authtypes.AccountI {
	ret := _m.Called(ctx, addr)

	var r0 authtypes.AccountI
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress) authtypes.AccountI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(authtypes.AccountI)
		}
	}

	return r0
}

// GetJob provides a mock function with given fields: ctx, jobID
func (_m *SchedulerKeeper) GetJob(ctx types.Context, jobID string) (*schedulertypes.Job, error) {
	ret := _m.Called(ctx, jobID)

	var r0 *schedulertypes.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, string) (*schedulertypes.Job, error)); ok {
		return rf(ctx, jobID)
	}
	if rf, ok := ret.Get(0).(func(types.Context, string) *schedulertypes.Job); ok {
		r0 = rf(ctx, jobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulertypes.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, string) error); ok {
		r1 = rf(ctx, jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logger provides a mock function with given fields: ctx
func (_m *SchedulerKeeper) Logger(ctx types.Context) log.Logger {
	ret := _m.Called(ctx)

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(types.Context) log.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// PreJobExecution provides a mock function with given fields: ctx, job
func (_m *SchedulerKeeper) PreJobExecution(ctx types.Context, job *schedulertypes.Job) error {
	ret := _m.Called(ctx, job)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *schedulertypes.Job) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScheduleNow provides a mock function with given fields: ctx, jobID, in, senderAddress, contractAddress
func (_m *SchedulerKeeper) ScheduleNow(ctx types.Context, jobID string, in []byte, senderAddress types.AccAddress, contractAddress types.AccAddress) error {
	ret := _m.Called(ctx, jobID, in, senderAddress, contractAddress)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, string, []byte, types.AccAddress, types.AccAddress) error); ok {
		r0 = rf(ctx, jobID, in, senderAddress, contractAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSchedulerKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewSchedulerKeeper creates a new instance of SchedulerKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSchedulerKeeper(t mockConstructorTestingTNewSchedulerKeeper) *SchedulerKeeper {
	mock := &SchedulerKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
