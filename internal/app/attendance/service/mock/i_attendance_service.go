package mocks

import (
	"context"
	"hr-system-salary/internal/app/attendance/model"

	"github.com/stretchr/testify/mock"
)

type IAttendanceService struct {
	mock.Mock
}

func (_m *IAttendanceService) AddAttendanceEmployee(ctx context.Context, username string) (res *model.AttendanceModel, err error) {
	ret := _m.Called(username)

	var r0 model.AttendanceModel
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) model.AttendanceModel); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(model.AttendanceModel)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return &r0, r1
}

func NewAttendanceServiceItf(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAttendanceService {
	mock := &IAttendanceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
