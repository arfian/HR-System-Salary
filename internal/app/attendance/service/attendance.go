package service

import (
	"context"
	"errors"
	"time"

	"hr-system-salary/internal/app/attendance/model"
	"hr-system-salary/internal/app/attendance/port"
	userPort "hr-system-salary/internal/app/user/port"
	"hr-system-salary/pkg/validations"
)

type service struct {
	attendanceRepo port.IAttendanceRepository
	userRepo       userPort.IUserRepository
}

func New(attendanceRepo port.IAttendanceRepository, userRepo userPort.IUserRepository) port.IAttendanceService {
	return &service{
		attendanceRepo: attendanceRepo,
		userRepo:       userRepo,
	}
}

func (s *service) AddAttendanceEmployee(ctx context.Context, username string) (*model.AttendanceModel, error) {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return nil, errors.New("user not found")
	}

	t := time.Now()
	if validations.IsWeekend(t) {
		return nil, errors.New("non-working days")
	}

	attendance := model.AttendanceModel{
		Employee: string(users[0].ID),
	}

	checkIn, qerr := s.attendanceRepo.GetCheckInAttendance(ctx, string(users[0].ID), t.Format("2006-01-02"))
	if qerr != nil {
		return nil, qerr
	}

	if len(checkIn) == 0 {
		attendance.CheckIn = time.Now()
		attendance.CreatedBy = users[0].Username
	} else if len(checkIn) > 0 {
		attendance.CheckOut = time.Now()
		attendance.UpdatedBy = users[0].Username
	} else {
		return nil, nil
	}

	res, qerr := s.attendanceRepo.InsertAttendanceEmployee(ctx, attendance)
	if qerr != nil {
		return nil, qerr
	}
	return &res, nil
}
