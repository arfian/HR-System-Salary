package port

import (
	"context"
	"hr-system-salary/internal/app/payroll/payload"
)

type IPayrollService interface {
	InsertPayroll(ctx context.Context, payroll payload.ParamGeneratePayroll, username string) error
}
