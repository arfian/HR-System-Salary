package repository

import (
	"context"

	"hr-system-salary/config/db"
	"hr-system-salary/pkg/cache"
	"hr-system-salary/pkg/transaction"

	"hr-system-salary/internal/app/payroll/model"
	"hr-system-salary/internal/app/payroll/port"
)

type repository struct {
	db    *db.GormDB
	cache cache.ICache
}

func NewRepository(db *db.GormDB) port.IPayrollRepository {
	return repository{db: db}
}

func (r repository) GeneratePayroll(ctx context.Context, payroll []model.PayrollModel) error {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&payroll).Error

	return qres
}

func (r repository) GetSettingPayroll(ctx context.Context) (res []model.SettingPayrollModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Find(&res).Error
	return res, err
}

func (r repository) GetPayrollByMonth(ctx context.Context, year int, month int) (res []model.PayrollModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Where("EXTRACT(MONTH FROM payroll_date) = ?", month).Where("EXTRACT(YEAR FROM payroll_date) = ?", year).Find(&res).Error
	return res, err
}

func (r repository) GetPayrollByMonthUserId(ctx context.Context, year int, month int, userId string) (res []model.PayrollModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Where("employee", userId).Where("EXTRACT(MONTH FROM payroll_date) = ?", month).Where("EXTRACT(YEAR FROM payroll_date) = ?", year).Find(&res).Error
	return res, err
}
