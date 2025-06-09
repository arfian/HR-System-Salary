package repository

import (
	"context"

	"hr-system-salary/config/db"
	"hr-system-salary/pkg/cache"
	"hr-system-salary/pkg/transaction"

	"hr-system-salary/internal/app/reimbursement/model"
	"hr-system-salary/internal/app/reimbursement/port"
)

type repository struct {
	db    *db.GormDB
	cache cache.ICache
}

func NewRepository(db *db.GormDB) port.IReimbursementRepository {
	return repository{db: db}
}

func (r repository) InsertReimbursement(ctx context.Context, reimbursement model.ReimbursementModel) (model.ReimbursementModel, error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&reimbursement).Error

	return reimbursement, qres
}
