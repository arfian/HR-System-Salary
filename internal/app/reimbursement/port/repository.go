package port

import (
	"context"
	"hr-system-salary/internal/app/reimbursement/model"
)

type IReimbursementRepository interface {
	InsertReimbursement(ctx context.Context, reimbursement model.ReimbursementModel) (model.ReimbursementModel, error)
	GetReimbursementByMonth(ctx context.Context, year int, month int, userId string) (res []model.ReimbursementModel, err error)
}
