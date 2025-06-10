package service

import (
	"context"
	"errors"
	"time"

	"hr-system-salary/internal/app/payroll/model"
	"hr-system-salary/internal/app/payroll/payload"
	"hr-system-salary/internal/app/payroll/port"
	reimbursementPort "hr-system-salary/internal/app/reimbursement/port"
	userPort "hr-system-salary/internal/app/user/port"
	"hr-system-salary/pkg/helper"
)

type service struct {
	payrollRepo       port.IPayrollRepository
	userRepo          userPort.IUserRepository
	reimbursementRepo reimbursementPort.IReimbursementRepository
}

func New(payrollRepo port.IPayrollRepository, userRepo userPort.IUserRepository, reimbursementRepo reimbursementPort.IReimbursementRepository) port.IPayrollService {
	return &service{
		payrollRepo:       payrollRepo,
		userRepo:          userRepo,
		reimbursementRepo: reimbursementRepo,
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
	checkPayrollDate, qerr := s.payrollRepo.GetPayrollByMonth(ctx, int(year), int(month))
	if len(checkPayrollDate) > 0 || qerr != nil {
		return errors.New("Payroll already exist")
	}

	settingPayroll, qerr := s.payrollRepo.GetSettingPayroll(ctx)
	if len(settingPayroll) == 0 {
		return errors.New("Setting Payroll does not exist")
	}

	pageNo := countEmployee / int64(limit)
	payrollData := []model.PayrollModel{}
	var (
		TotalOvertime         float32 = 0
		CountAbsence          int     = 0
		TotalDeductionAbsence float32 = 0
		PercentageDeduction   float32 = 0
		TotalWeekdays         int     = 0
	)
	for i := 0; i <= int(pageNo); i++ {
		employes, qerr := s.userRepo.GetAttendanceOvertimeByEmployee(ctx, limit, i, int(year), int(month))
		if qerr != nil {
			return qerr
		}

		payrollData = []model.PayrollModel{}
		for _, e := range employes {
			TotalOvertime = (float32(e.SumOvertime) * settingPayroll[0].OvertimeRateHours)
			TotalWeekdays = helper.CountWeekdays(int(month), int(year))
			CountAbsence = TotalWeekdays - e.TotalAttendance
			PercentageDeduction = 100 - ((float32(e.TotalAttendance) / float32(TotalWeekdays)) * 100)
			TotalDeductionAbsence = (e.BasicSalary * (PercentageDeduction / 100))

			payrollData = append(payrollData, model.PayrollModel{
				Employee:              e.ID,
				PayrollDate:           t,
				BasicSalary:           e.BasicSalary,
				TotalAttendence:       e.TotalAttendance,
				CountOvertime:         e.SumOvertime,
				OvertimeRateHours:     settingPayroll[0].OvertimeRateHours,
				TotalOvertime:         TotalOvertime,
				TotalReimbursement:    e.TotalReimbursement,
				Status:                "PAID",
				CreatedBy:             username,
				CountAbsence:          CountAbsence,
				TotalDeductionAbsence: TotalDeductionAbsence,
				TotalTakeHomePay:      (e.BasicSalary - TotalDeductionAbsence + e.TotalReimbursement + TotalOvertime),
			})
		}

		if len(payrollData) > 0 {
			qerr = s.payrollRepo.GeneratePayroll(ctx, payrollData)
			if qerr != nil {
				return qerr
			}
		}
	}

	return nil
}

func (s *service) GetPayrollByMonth(ctx context.Context, payroll payload.ParamGeneratePayroll, username string) (*payload.ResPayslip, error) {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return nil, errors.New("user not found")
	}

	t, _ := time.Parse("2006-01", payroll.PayrollDate)
	month := t.Month()
	year := t.Year()
	payrollData, qerr := s.payrollRepo.GetPayrollByMonthUserId(ctx, int(year), int(month), users[0].ID.String())
	if qerr != nil {
		return nil, qerr
	}

	if len(payrollData) == 0 {
		return nil, errors.New("Payroll not found")
	}

	reimbursementsData, qerr := s.reimbursementRepo.GetReimbursementByMonth(ctx, int(year), int(month), users[0].ID.String())
	if qerr != nil {
		return nil, qerr
	}

	res := &payload.ResPayslip{
		Payslip:        payrollData[0],
		Reimbursements: reimbursementsData,
	}
	return res, nil
}

func (s *service) GetAllPayrollByMonth(ctx context.Context, payroll payload.ParamGeneratePayroll, username string) (*payload.ResAllPayslip, error) {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return nil, errors.New("user not found")
	}

	t, _ := time.Parse("2006-01", payroll.PayrollDate)
	month := t.Month()
	year := t.Year()
	payrollData, qerr := s.payrollRepo.GetPayrollByMonth(ctx, int(year), int(month))
	if qerr != nil {
		return nil, qerr
	}

	if len(payrollData) == 0 {
		return nil, errors.New("Payroll not found")
	}

	var TotalTakeHomePay float32 = 0
	for _, e := range payrollData {
		TotalTakeHomePay += e.TotalTakeHomePay
	}

	res := &payload.ResAllPayslip{
		Payslip:          payrollData,
		TotalTakeHomePay: TotalTakeHomePay,
	}

	return res, nil
}
