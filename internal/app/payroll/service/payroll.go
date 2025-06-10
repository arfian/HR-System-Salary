package service

import (
	"context"
	"errors"
	"time"

	"hr-system-salary/internal/app/payroll/model"
	"hr-system-salary/internal/app/payroll/payload"
	"hr-system-salary/internal/app/payroll/port"
	userPort "hr-system-salary/internal/app/user/port"
)

type service struct {
	payrollRepo port.IPayrollRepository
	userRepo    userPort.IUserRepository
}

func New(payrollRepo port.IPayrollRepository, userRepo userPort.IUserRepository) port.IPayrollService {
	return &service{
		payrollRepo: payrollRepo,
		userRepo:    userRepo,
	}
}

func (s *service) InsertPayroll(ctx context.Context, payroll payload.ParamGeneratePayroll, username string) error {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return errors.New("user not found")
	}

	limit := 50
	countEmployee, qerr := s.userRepo.GetCountEmployee(ctx)
	if qerr != nil {
		return qerr
	}

	t, _ := time.Parse("2006-01", payroll.PayrollDate)
	month := t.Month()
	year := t.Year()

	settingPayroll, qerr := s.payrollRepo.GetSettingPayroll(ctx)
	if len(settingPayroll) == 0 {
		return errors.New("Setting Payroll does not exist")
	}
	pageNo := countEmployee / int64(limit)
	payrollData := []model.PayrollModel{}
	for i := 0; i <= int(pageNo); i++ {
		employes, qerr := s.userRepo.GetAttendanceOvertimeByEmployee(ctx, limit, i, int(month), int(year))
		if qerr != nil {
			return qerr
		}
		payrollData = []model.PayrollModel{}
		for _, e := range employes {
			payrollData = append(payrollData, model.PayrollModel{
				Employee:              e.ID,
				BasicSalary:           e.BasicSalary,
				TotalAttendence:       e.TotalAttendance,
				CountOvertime:         e.SumOvertime,
				OvertimeRateHours:     settingPayroll[0].OvertimeRateHours,
				TotalOvertime:         (float32(e.SumOvertime) * settingPayroll[0].OvertimeRateHours),
				TotalReimbursement:    e.TotalReimbursement,
				Status:                "PAID",
				CreatedBy:             username,
				CountAbsence:          0,
				TotalDeductionAbsence: 0,
				TotalTakeHomePay:      0,
			})
		}
		qerr = s.payrollRepo.GeneratePayroll(ctx, payrollData)
		if qerr != nil {
			return qerr
		}
	}

	return nil
}
