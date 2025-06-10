package port

import (
	"context"
	"hr-system-salary/internal/app/payroll/payload"
)

type IPayrollService interface {
	InsertPayroll(ctx context.Context, payroll payload.ParamGeneratePayroll, username string) error
	GetPayrollByMonth(ctx context.Context, payroll payload.ParamGeneratePayroll, username string) (*payload.ResPayslip, error)
	GetAllPayrollByMonth(ctx context.Context, payroll payload.ParamGeneratePayroll, username string) (*payload.ResAllPayslip, error)
}
