package port

import (
	"context"
	"hr-system-salary/internal/app/payroll/model"
)

type IPayrollRepository interface {
	GeneratePayroll(ctx context.Context, payroll []model.PayrollModel) error
	GetSettingPayroll(ctx context.Context) (res []model.SettingPayrollModel, err error)
	GetPayrollByMonth(ctx context.Context, year int, month int) (res []model.PayrollModel, err error)
}
