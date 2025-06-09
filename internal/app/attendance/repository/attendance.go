package repository

import (
	"context"

	"hr-system-salary/config/db"
	"hr-system-salary/pkg/cache"
	"hr-system-salary/pkg/transaction"

	"hr-system-salary/internal/app/attendance/model"
	"hr-system-salary/internal/app/attendance/port"
)

type repository struct {
	db    *db.GormDB
	cache cache.ICache
}

func NewRepository(db *db.GormDB) port.IAttendanceRepository {
	return repository{db: db}
}

func (r repository) InsertAttendanceEmployee(ctx context.Context, attendance model.AttendanceModel) (model.AttendanceModel, error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&attendance).Error

	return attendance, qres
}
