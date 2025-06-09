package service

import (
	"context"
	"errors"

	"hr-system-salary/internal/app/attendance/model"
	"hr-system-salary/internal/app/attendance/port"
	userPort "hr-system-salary/internal/app/user/port"
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

func (s *service) AddAttendanceEmployee(ctx context.Context, username string) error {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return errors.New("user not found")
	}

	attendance := model.AttendanceModel{}
	_, qerr = s.attendanceRepo.InsertAttendanceEmployee(ctx, attendance)
	if qerr != nil {
		return qerr
	}
	return nil
}
