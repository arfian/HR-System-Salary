package service

import (
	"context"
	"errors"
	"time"

	"hr-system-salary/internal/app/attendance/model"
	"hr-system-salary/internal/app/attendance/payload"
	"hr-system-salary/internal/app/attendance/port"
	userPort "hr-system-salary/internal/app/user/port"
	"hr-system-salary/pkg/helper"
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

func (s *service) AddAttendanceEmployee(ctx context.Context, username string) (res *model.AttendanceModel, err error) {
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

	resData, qerr := s.attendanceRepo.GetAttendanceByUserDate(ctx, string(users[0].ID), t.Format("2006-01-02"))
	if qerr != nil {
		return nil, qerr
	}

	if len(resData) == 0 {
		attendance.CheckIn = time.Now()
		attendance.CreatedBy = users[0].Username
		attendance.Status = "CHECK IN"
	} else if len(resData) > 0 && resData[0].Status == "CHECK IN" {
		attendance = resData[0]
		attendance.CheckOut = time.Now()
		attendance.UpdatedBy = users[0].Username
		attendance.DeletedAt = nil
		attendance.Status = "CHECK OUT"
	} else {
		attendance = resData[0]
		attendance.Status = ""
		return nil, nil
	}

	resAttendence, qerr := s.attendanceRepo.InsertAttendanceEmployee(ctx, attendance)
	if qerr != nil {
		return nil, qerr
	}
	return &resAttendence, nil
}

func (s *service) BulkInserAttendance(ctx context.Context, param payload.ParamBulkAttendance, username string) error {
	listWeekdays := helper.ListWeekdays(param.StartDate, param.EndDate)
	checkAttendances, qerr := s.attendanceRepo.GetDateRangeAttendanceByUser(ctx, param.EmployeeID, param.StartDate, param.EndDate)
	if qerr != nil {
		return qerr
	}

	listAttendances := helper.DifferenceDate(listWeekdays, checkAttendances)
	attendances := []model.AttendanceModel{}
	checkIn := time.Now()
	for _, a := range listAttendances {
		checkIn, _ = time.Parse("2006-01-02", a)
		attendances = append(attendances, model.AttendanceModel{
			CheckIn:   checkIn,
			CreatedBy: username,
			Status:    "CHECK IN",
			Employee:  param.EmployeeID,
		})
	}

	if len(attendances) > 0 {
		qerr = s.attendanceRepo.BulkInsertAttendance(ctx, attendances)
		if qerr != nil {
			return qerr
		}
	}

	return nil
}

func (s *service) AddOvertime(ctx context.Context, param payload.ParamOvertime, username string) (*model.OvertimeModel, error) {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return nil, errors.New("user not found")
	}
	t := time.Now()

	if !validations.IsWeekend(t) {
		checkOut, qerr := s.attendanceRepo.GetAttendanceStatusByUserDate(ctx, users[0].ID.String(), t.Format("2006-01-02"), "CHECK OUT")
		if qerr != nil {
			return nil, qerr
		}
		if len(checkOut) == 0 {
			return nil, errors.New("overtime must be proposed after check out")
		}
	}

	sumHours, qerr := s.attendanceRepo.GetSumOvertimeByUserDate(ctx, users[0].ID.String(), t.Format("2006-01-02"))
	if (sumHours + param.Hours) > 3 {
		return nil, errors.New("overtime max 3 hours")
	}

	overtime := model.OvertimeModel{
		Employee:      users[0].ID.String(),
		OvertimeHours: param.Hours,
		OvertimeDate:  t,
		Status:        "APPROVED",
		CreatedBy:     username,
	}
	overtime, qerr = s.attendanceRepo.InsertOvertime(ctx, overtime)

	return &overtime, nil
}
