package port

import (
	"context"
	"hr-system-salary/internal/app/reimbursement/model"
	"hr-system-salary/internal/app/reimbursement/payload"
)

type IReimbursementService interface {
	AddReimbursement(ctx context.Context, username string, param payload.ParamReimbursement) (res *model.ReimbursementModel, err error)
}
